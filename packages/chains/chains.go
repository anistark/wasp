// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package chains

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/iotaledger/hive.go/ds/shrinkingmap"
	"github.com/iotaledger/hive.go/lo"
	"github.com/iotaledger/hive.go/logger"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/chain/cmt_log"
	"github.com/iotaledger/wasp/packages/chain/statemanager/sm_gpa"
	"github.com/iotaledger/wasp/packages/chain/statemanager/sm_gpa/sm_gpa_utils"
	"github.com/iotaledger/wasp/packages/chains/access_mgr"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/database"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/shutdown"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/state/indexedstore"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/processors"
	"github.com/iotaledger/wasp/packages/webapi/interfaces"
)

type Provider func() *Chains // TODO: Use DI instead of that.

type ChainProvider func(chainID isc.ChainID) chain.Chain

type Chains struct {
	ctx                              context.Context
	log                              *logger.Logger
	nodeConnection                   chain.NodeConnection
	processorConfig                  *processors.Config
	offledgerBroadcastUpToNPeers     int
	offledgerBroadcastInterval       time.Duration
	pullMissingRequestsFromCommittee bool
	deriveAliasOutputByQuorum        bool
	pipeliningLimit                  int
	consensusDelay                   time.Duration

	networkProvider              peering.NetworkProvider
	trustedNetworkManager        peering.TrustedNetworkManager
	trustedNetworkListenerCancel context.CancelFunc
	chainStateStoreProvider      database.ChainStateKVStoreProvider

	walEnabled                          bool
	walFolderPath                       string
	smBlockCacheMaxSize                 int
	smBlockCacheBlocksInCacheDuration   time.Duration
	smBlockCacheBlockCleaningPeriod     time.Duration
	smStateManagerGetBlockRetry         time.Duration
	smStateManagerRequestCleaningPeriod time.Duration
	smStateManagerTimerTickPeriod       time.Duration
	smPruningMinStatesToKeep            int
	smPruningMaxStatesToDelete          int

	chainRecordRegistryProvider registry.ChainRecordRegistryProvider
	dkShareRegistryProvider     registry.DKShareRegistryProvider
	nodeIdentityProvider        registry.NodeIdentityProvider
	consensusStateRegistry      cmt_log.ConsensusStateRegistry
	chainListener               chain.ChainListener

	mutex     *sync.RWMutex
	allChains *shrinkingmap.ShrinkingMap[isc.ChainID, *activeChain]
	accessMgr access_mgr.AccessMgr

	cleanupFunc         context.CancelFunc
	shutdownCoordinator *shutdown.Coordinator

	chainMetricsProvider *metrics.ChainMetricsProvider

	validatorFeeAddr iotago.Address
}

type activeChain struct {
	chain      chain.Chain
	cancelFunc context.CancelFunc
}

func New(
	log *logger.Logger,
	nodeConnection chain.NodeConnection,
	processorConfig *processors.Config,
	validatorAddrStr string,
	offledgerBroadcastUpToNPeers int, // TODO: Unused for now.
	offledgerBroadcastInterval time.Duration, // TODO: Unused for now.
	pullMissingRequestsFromCommittee bool, // TODO: Unused for now.
	deriveAliasOutputByQuorum bool,
	pipeliningLimit int,
	consensusDelay time.Duration,
	networkProvider peering.NetworkProvider,
	trustedNetworkManager peering.TrustedNetworkManager,
	chainStateStoreProvider database.ChainStateKVStoreProvider,
	walEnabled bool,
	walFolderPath string,
	smBlockCacheMaxSize int,
	smBlockCacheBlocksInCacheDuration time.Duration,
	smBlockCacheBlockCleaningPeriod time.Duration,
	smStateManagerGetBlockRetry time.Duration,
	smStateManagerRequestCleaningPeriod time.Duration,
	smStateManagerTimerTickPeriod time.Duration,
	smPruningMinStatesToKeep int,
	smPruningMaxStatesToDelete int,
	chainRecordRegistryProvider registry.ChainRecordRegistryProvider,
	dkShareRegistryProvider registry.DKShareRegistryProvider,
	nodeIdentityProvider registry.NodeIdentityProvider,
	consensusStateRegistry cmt_log.ConsensusStateRegistry,
	chainListener chain.ChainListener,
	shutdownCoordinator *shutdown.Coordinator,
	chainMetricsProvider *metrics.ChainMetricsProvider,
) *Chains {
	var validatorFeeAddr iotago.Address
	if validatorAddrStr != "" {
		bechPrefix, addr, err := iotago.ParseBech32(validatorAddrStr)
		if err != nil {
			panic(fmt.Errorf("error parsing validator.address: %s", err.Error()))
		}
		if bechPrefix != nodeConnection.GetL1Params().Protocol.Bech32HRP {
			panic(fmt.Errorf("validator.address Bech32 HRP does not match network HRP, expected: %s, got: %s", nodeConnection.GetL1Params().Protocol.Bech32HRP, bechPrefix))
		}
		validatorFeeAddr = addr
	}
	ret := &Chains{
		log:                                 log,
		mutex:                               &sync.RWMutex{},
		allChains:                           shrinkingmap.New[isc.ChainID, *activeChain](),
		nodeConnection:                      nodeConnection,
		processorConfig:                     processorConfig,
		offledgerBroadcastUpToNPeers:        offledgerBroadcastUpToNPeers,
		offledgerBroadcastInterval:          offledgerBroadcastInterval,
		pullMissingRequestsFromCommittee:    pullMissingRequestsFromCommittee,
		deriveAliasOutputByQuorum:           deriveAliasOutputByQuorum,
		pipeliningLimit:                     pipeliningLimit,
		consensusDelay:                      consensusDelay,
		networkProvider:                     networkProvider,
		trustedNetworkManager:               trustedNetworkManager,
		chainStateStoreProvider:             chainStateStoreProvider,
		walEnabled:                          walEnabled,
		walFolderPath:                       walFolderPath,
		smBlockCacheMaxSize:                 smBlockCacheMaxSize,
		smBlockCacheBlocksInCacheDuration:   smBlockCacheBlocksInCacheDuration,
		smBlockCacheBlockCleaningPeriod:     smBlockCacheBlockCleaningPeriod,
		smStateManagerGetBlockRetry:         smStateManagerGetBlockRetry,
		smStateManagerRequestCleaningPeriod: smStateManagerRequestCleaningPeriod,
		smStateManagerTimerTickPeriod:       smStateManagerTimerTickPeriod,
		smPruningMinStatesToKeep:            smPruningMinStatesToKeep,
		smPruningMaxStatesToDelete:          smPruningMaxStatesToDelete,
		chainRecordRegistryProvider:         chainRecordRegistryProvider,
		dkShareRegistryProvider:             dkShareRegistryProvider,
		nodeIdentityProvider:                nodeIdentityProvider,
		chainListener:                       nil, // See bellow.
		consensusStateRegistry:              consensusStateRegistry,
		shutdownCoordinator:                 shutdownCoordinator,
		chainMetricsProvider:                chainMetricsProvider,
		validatorFeeAddr:                    validatorFeeAddr,
	}
	ret.chainListener = NewChainsListener(chainListener, ret.chainAccessUpdatedCB)
	return ret
}

func (c *Chains) Run(ctx context.Context) error {
	if err := c.nodeConnection.WaitUntilInitiallySynced(ctx); err != nil {
		return fmt.Errorf("waiting for L1 node to become sync failed, error: %w", err)
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.ctx != nil {
		return errors.New("chains already running")
	}
	c.ctx = ctx

	c.accessMgr = access_mgr.New(ctx, c.chainServersUpdatedCB, c.nodeIdentityProvider.NodeIdentity(), c.networkProvider, c.log.Named("AM"))
	c.trustedNetworkListenerCancel = c.trustedNetworkManager.TrustedPeersListener(c.trustedPeersUpdatedCB)

	unhook := c.chainRecordRegistryProvider.Events().ChainRecordModified.Hook(func(event *registry.ChainRecordModifiedEvent) {
		c.mutex.RLock()
		defer c.mutex.RUnlock()
		if chain, exists := c.allChains.Get(event.ChainRecord.ChainID()); exists {
			chain.chain.ConfigUpdated(event.ChainRecord.AccessNodes)
		}
	}).Unhook
	c.cleanupFunc = unhook

	return c.activateAllFromRegistry() //nolint:contextcheck
}

func (c *Chains) Close() {
	util.ExecuteIfNotNil(c.cleanupFunc)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.allChains.ForEach(func(_ isc.ChainID, ac *activeChain) bool {
		ac.cancelFunc()
		return true
	})
	c.shutdownCoordinator.WaitNestedWithLogging(1 * time.Second)
	c.shutdownCoordinator.Done()
	util.ExecuteIfNotNil(c.trustedNetworkListenerCancel)
	c.trustedNetworkListenerCancel = nil
}

func (c *Chains) trustedPeersUpdatedCB(trustedPeers []*peering.TrustedPeer) {
	trustedPubKeys := lo.Map(trustedPeers, func(tp *peering.TrustedPeer) *cryptolib.PublicKey { return tp.PubKey() })
	c.accessMgr.TrustedNodes(trustedPubKeys)
}

func (c *Chains) chainServersUpdatedCB(chainID isc.ChainID, servers []*cryptolib.PublicKey) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	ch, exists := c.allChains.Get(chainID)
	if !exists {
		return
	}
	ch.chain.ServersUpdated(servers)
}

func (c *Chains) chainAccessUpdatedCB(chainID isc.ChainID, accessNodes []*cryptolib.PublicKey) {
	c.accessMgr.ChainAccessNodes(chainID, accessNodes)
}

func (c *Chains) activateAllFromRegistry() error {
	var innerErr error
	if err := c.chainRecordRegistryProvider.ForEachActiveChainRecord(func(chainRecord *registry.ChainRecord) bool {
		chainID := chainRecord.ChainID()
		if err := c.activateWithoutLocking(chainID); err != nil {
			innerErr = fmt.Errorf("cannot activate chain %s: %w", chainRecord.ChainID(), err)
			return false
		}

		return true
	}); err != nil {
		return err
	}

	return innerErr
}

// activateWithoutLocking activates a chain in the node.
func (c *Chains) activateWithoutLocking(chainID isc.ChainID) error {
	if c.ctx == nil {
		return errors.New("run chains first")
	}
	if c.ctx.Err() != nil {
		return errors.New("node is shutting down")
	}

	//
	// Check, maybe it is already running.
	if c.allChains.Has(chainID) {
		c.log.Debugf("Chain %v = %v is already activated", chainID.ShortString(), chainID.String())
		return nil
	}
	//
	// Activate the chain in the persistent store, if it is not activated yet.
	chainRecord, err := c.chainRecordRegistryProvider.ChainRecord(chainID)
	if err != nil {
		return fmt.Errorf("cannot get chain record for %v: %w", chainID, err)
	}
	if !chainRecord.Active {
		if _, err2 := c.chainRecordRegistryProvider.ActivateChainRecord(chainID); err2 != nil {
			return fmt.Errorf("cannot activate chain: %w", err2)
		}
	}

	chainKVStore, err := c.chainStateStoreProvider(chainID)
	if err != nil {
		return fmt.Errorf("error when creating chain KV store: %w", err)
	}

	chainMetrics := c.chainMetricsProvider.GetChainMetrics(chainID)

	// Initialize WAL
	chainLog := c.log.Named(chainID.ShortString())
	var chainWAL sm_gpa_utils.BlockWAL
	if c.walEnabled {
		chainWAL, err = sm_gpa_utils.NewBlockWAL(chainLog.Named("WAL"), c.walFolderPath, chainID, chainMetrics.BlockWAL)
		if err != nil {
			panic(fmt.Errorf("cannot create WAL: %w", err))
		}
	} else {
		chainWAL = sm_gpa_utils.NewEmptyBlockWAL()
	}

	stateManagerParameters := sm_gpa.NewStateManagerParameters()
	stateManagerParameters.BlockCacheMaxSize = c.smBlockCacheMaxSize
	stateManagerParameters.BlockCacheBlocksInCacheDuration = c.smBlockCacheBlocksInCacheDuration
	stateManagerParameters.BlockCacheBlockCleaningPeriod = c.smBlockCacheBlockCleaningPeriod
	stateManagerParameters.StateManagerGetBlockRetry = c.smStateManagerGetBlockRetry
	stateManagerParameters.StateManagerRequestCleaningPeriod = c.smStateManagerRequestCleaningPeriod
	stateManagerParameters.StateManagerTimerTickPeriod = c.smStateManagerTimerTickPeriod
	stateManagerParameters.PruningMinStatesToKeep = c.smPruningMinStatesToKeep
	stateManagerParameters.PruningMaxStatesToDelete = c.smPruningMaxStatesToDelete

	chainCtx, chainCancel := context.WithCancel(c.ctx)
	validatorAgentID := accounts.CommonAccount()
	if c.validatorFeeAddr != nil {
		validatorAgentID = isc.NewAgentID(c.validatorFeeAddr)
	}
	newChain, err := chain.New(
		chainCtx,
		chainLog,
		chainID,
		indexedstore.New(state.NewStoreWithMetrics(chainKVStore, chainMetrics.State)),
		c.nodeConnection,
		c.nodeIdentityProvider.NodeIdentity(),
		c.processorConfig,
		c.dkShareRegistryProvider,
		c.consensusStateRegistry,
		chainWAL,
		c.chainListener,
		chainRecord.AccessNodes,
		c.networkProvider,
		chainMetrics,
		c.shutdownCoordinator.Nested(fmt.Sprintf("Chain-%s", chainID.AsAddress().String())),
		func() { c.chainMetricsProvider.RegisterChain(chainID) },
		func() { c.chainMetricsProvider.UnregisterChain(chainID) },
		c.deriveAliasOutputByQuorum,
		c.pipeliningLimit,
		c.consensusDelay,
		validatorAgentID,
		stateManagerParameters,
	)
	if err != nil {
		chainCancel()
		return fmt.Errorf("Chains.Activate: failed to create chain object: %w", err)
	}
	c.allChains.Set(chainID, &activeChain{
		chain:      newChain,
		cancelFunc: chainCancel,
	})

	c.log.Infof("activated chain: %v = %s", chainID.ShortString(), chainID.String())
	return nil
}

// Activate activates a chain in the node.
func (c *Chains) Activate(chainID isc.ChainID) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.activateWithoutLocking(chainID)
}

// Deactivate a chain in the node.
func (c *Chains) Deactivate(chainID isc.ChainID) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, err := c.chainRecordRegistryProvider.DeactivateChainRecord(chainID); err != nil {
		return fmt.Errorf("cannot deactivate chain %v: %w", chainID, err)
	}

	ch, exists := c.allChains.Get(chainID)
	if !exists {
		c.log.Debugf("chain is not active: %v = %s", chainID.ShortString(), chainID.String())
		return nil
	}
	ch.cancelFunc()
	c.accessMgr.ChainDismissed(chainID)
	c.allChains.Delete(chainID)
	c.log.Debugf("chain has been deactivated: %v = %s", chainID.ShortString(), chainID.String())
	return nil
}

// Get returns active chain object or nil if it doesn't exist
// lazy unsubscribing
func (c *Chains) Get(chainID isc.ChainID) (chain.Chain, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	ret, exists := c.allChains.Get(chainID)
	if !exists {
		return nil, interfaces.ErrChainNotFound
	}
	return ret.chain, nil
}

func (c *Chains) ValidatorAddress() iotago.Address {
	return c.validatorFeeAddr
}

func (c *Chains) IsArchiveNode() bool {
	return c.smPruningMinStatesToKeep < 1
}
