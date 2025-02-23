// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// package jsonrpc implements JSON-RPC endpoints according to
// https://eth.wiki/json-rpc/API
package jsonrpc

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"

	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/metrics"
	vmerrors "github.com/iotaledger/wasp/packages/vm/core/errors"
)

type EthService struct {
	evmChain *EVMChain
	accounts *AccountManager
	metrics  *metrics.ChainWebAPIMetrics
}

func NewEthService(evmChain *EVMChain, accounts *AccountManager, metrics *metrics.ChainWebAPIMetrics) *EthService {
	return &EthService{
		evmChain: evmChain,
		accounts: accounts,
		metrics:  metrics,
	}
}

func (e *EthService) ProtocolVersion() hexutil.Uint {
	return hexutil.Uint(eth.ETH66)
}

func (e *EthService) resolveError(err error) error {
	if err == nil {
		return nil
	}
	if vmError, ok := err.(*isc.UnresolvedVMError); ok {
		resolvedErr, resolveErr := vmerrors.Resolve(vmError, e.evmChain.ViewCaller(e.evmChain.backend.ISCLatestState()))
		if resolveErr != nil {
			return fmt.Errorf("could not resolve VMError %w: %v", vmError, resolveErr)
		}
		return resolvedErr.AsGoError()
	}
	return err
}

func (e *EthService) getTransactionCount(address common.Address, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	n, err := e.evmChain.TransactionCount(address, blockNumberOrHash)
	if err != nil {
		return 0, e.resolveError(err)
	}
	return hexutil.Uint64(n), nil
}

func (e *EthService) GetTransactionCount(address common.Address, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	return withMetrics(
		e.metrics, "eth_getTransactionCount",
		func() (hexutil.Uint64, error) {
			return e.getTransactionCount(address, blockNumberOrHash)
		},
	)
}

func (e *EthService) blockNumber() *hexutil.Big {
	n := e.evmChain.BlockNumber()
	return (*hexutil.Big)(n)
}

func (e *EthService) BlockNumber() (*hexutil.Big, error) {
	return withMetrics(
		e.metrics, "eth_blockNumber",
		func() (*hexutil.Big, error) {
			return e.blockNumber(), nil
		},
	)
}

func (e *EthService) getBlockByNumber(blockNumber rpc.BlockNumber, full bool) (map[string]interface{}, error) {
	block, err := e.evmChain.BlockByNumber(parseBlockNumber(blockNumber))
	if err != nil {
		return nil, e.resolveError(err)
	}
	if block == nil {
		return nil, nil
	}
	return RPCMarshalBlock(block, true, full)
}

func (e *EthService) GetBlockByNumber(blockNumber rpc.BlockNumber, full bool) (map[string]interface{}, error) {
	return withMetrics(
		e.metrics, "eth_getBlockByNumber",
		func() (map[string]interface{}, error) {
			return e.getBlockByNumber(blockNumber, full)
		},
	)
}

func (e *EthService) getBlockByHash(hash common.Hash, full bool) (map[string]interface{}, error) {
	block := e.evmChain.BlockByHash(hash)
	if block == nil {
		return nil, nil
	}
	return RPCMarshalBlock(block, true, full)
}

func (e *EthService) GetBlockByHash(hash common.Hash, full bool) (map[string]interface{}, error) {
	return withMetrics(
		e.metrics, "eth_getBlockByHash",
		func() (map[string]interface{}, error) {
			return e.getBlockByHash(hash, full)
		},
	)
}

func (e *EthService) getTransactionByHash(hash common.Hash) (*RPCTransaction, error) {
	tx, blockHash, blockNumber, index, err := e.evmChain.TransactionByHash(hash)
	if err != nil {
		return nil, e.resolveError(err)
	}
	if tx == nil {
		return nil, nil
	}
	return newRPCTransaction(tx, blockHash, blockNumber, index), err
}

func (e *EthService) GetTransactionByHash(hash common.Hash) (*RPCTransaction, error) {
	return withMetrics(
		e.metrics, "eth_getTransactionByHash",
		func() (*RPCTransaction, error) {
			return e.getTransactionByHash(hash)
		},
	)
}

func (e *EthService) getTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (*RPCTransaction, error) {
	tx, blockNumber, err := e.evmChain.TransactionByBlockHashAndIndex(blockHash, uint64(index))
	if err != nil {
		return nil, e.resolveError(err)
	}
	if tx == nil {
		return nil, nil
	}
	return newRPCTransaction(tx, blockHash, blockNumber, uint64(index)), err
}

func (e *EthService) GetTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (*RPCTransaction, error) {
	return withMetrics(
		e.metrics, "eth_getTransactionByBlockHashAndIndex",
		func() (*RPCTransaction, error) {
			return e.getTransactionByBlockHashAndIndex(blockHash, index)
		},
	)
}

func (e *EthService) getTransactionByBlockNumberAndIndex(blockNumberOrTag rpc.BlockNumber, index hexutil.Uint) (*RPCTransaction, error) {
	tx, blockHash, blockNumber, err := e.evmChain.TransactionByBlockNumberAndIndex(parseBlockNumber(blockNumberOrTag), uint64(index))
	if err != nil {
		return nil, e.resolveError(err)
	}
	if tx == nil {
		return nil, nil
	}
	return newRPCTransaction(tx, blockHash, blockNumber, uint64(index)), err
}

func (e *EthService) GetTransactionByBlockNumberAndIndex(blockNumberOrTag rpc.BlockNumber, index hexutil.Uint) (*RPCTransaction, error) {
	return withMetrics(
		e.metrics, "eth_getTransactionByBlockNumberAndIndex",
		func() (*RPCTransaction, error) {
			return e.getTransactionByBlockNumberAndIndex(blockNumberOrTag, index)
		},
	)
}

func (e *EthService) getBalance(address common.Address, blockNumberOrHash *rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	bal, err := e.evmChain.Balance(address, blockNumberOrHash)
	if err != nil {
		return nil, e.resolveError(err)
	}
	return (*hexutil.Big)(bal), nil
}

func (e *EthService) GetBalance(address common.Address, blockNumberOrHash *rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	return withMetrics(
		e.metrics, "eth_getBalance",
		func() (*hexutil.Big, error) {
			return e.getBalance(address, blockNumberOrHash)
		},
	)
}

func (e *EthService) getCode(address common.Address, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	code, err := e.evmChain.Code(address, blockNumberOrHash)
	if err != nil {
		return nil, e.resolveError(err)
	}
	return code, nil
}

func (e *EthService) GetCode(address common.Address, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	return withMetrics(
		e.metrics, "eth_getCode",
		func() (hexutil.Bytes, error) {
			return e.getCode(address, blockNumberOrHash)
		},
	)
}

func (e *EthService) getTransactionReceipt(txHash common.Hash) (map[string]interface{}, error) {
	r := e.evmChain.TransactionReceipt(txHash)
	if r == nil {
		return nil, nil
	}
	tx, _, _, _, err := e.evmChain.TransactionByHash(txHash)
	if err != nil {
		return nil, e.resolveError(err)
	}
	return RPCMarshalReceipt(r, tx), nil
}

func (e *EthService) GetTransactionReceipt(txHash common.Hash) (map[string]interface{}, error) {
	return withMetrics(
		e.metrics, "eth_getTransactionReceipt",
		func() (map[string]interface{}, error) {
			return e.getTransactionReceipt(txHash)
		},
	)
}

func (e *EthService) sendRawTransaction(txBytes hexutil.Bytes) (common.Hash, error) {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(txBytes, tx); err != nil {
		return common.Hash{}, err
	}
	if err := e.evmChain.SendTransaction(tx); err != nil {
		return common.Hash{}, e.resolveError(err)
	}
	return tx.Hash(), nil
}

func (e *EthService) SendRawTransaction(txBytes hexutil.Bytes) (common.Hash, error) {
	return withMetrics(
		e.metrics, "eth_sendRawTransaction",
		func() (common.Hash, error) {
			return e.sendRawTransaction(txBytes)
		},
	)
}

func (e *EthService) call(args *RPCCallArgs, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	ret, err := e.evmChain.CallContract(args.parse(), blockNumberOrHash)
	return ret, e.resolveError(err)
}

func (e *EthService) Call(args *RPCCallArgs, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	return withMetrics(
		e.metrics, "eth_call",
		func() (hexutil.Bytes, error) {
			return e.call(args, blockNumberOrHash)
		},
	)
}

func (e *EthService) estimateGas(args *RPCCallArgs, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	gas, err := e.evmChain.EstimateGas(args.parse(), blockNumberOrHash)
	return hexutil.Uint64(gas), e.resolveError(err)
}

func (e *EthService) EstimateGas(args *RPCCallArgs, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	return withMetrics(
		e.metrics, "eth_estimateGas",
		func() (hexutil.Uint64, error) {
			return e.estimateGas(args, blockNumberOrHash)
		},
	)
}

func (e *EthService) getStorageAt(address common.Address, key string, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	ret, err := e.evmChain.StorageAt(address, common.HexToHash(key), blockNumberOrHash)
	return ret[:], e.resolveError(err)
}

func (e *EthService) GetStorageAt(address common.Address, key string, blockNumberOrHash *rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	return withMetrics(
		e.metrics, "eth_getStorageAt",
		func() (hexutil.Bytes, error) {
			return e.getStorageAt(address, key, blockNumberOrHash)
		},
	)
}

func (e *EthService) getBlockTransactionCountByHash(blockHash common.Hash) hexutil.Uint {
	ret := e.evmChain.BlockTransactionCountByHash(blockHash)
	return hexutil.Uint(ret)
}

func (e *EthService) GetBlockTransactionCountByHash(blockHash common.Hash) (hexutil.Uint, error) {
	return withMetrics(
		e.metrics, "eth_getBlockTransactionCountByHash",
		func() (hexutil.Uint, error) {
			return e.getBlockTransactionCountByHash(blockHash), nil
		},
	)
}

func (e *EthService) getBlockTransactionCountByNumber(blockNumber rpc.BlockNumber) (hexutil.Uint, error) {
	ret, err := e.evmChain.BlockTransactionCountByNumber(parseBlockNumber(blockNumber))
	return hexutil.Uint(ret), e.resolveError(err)
}

func (e *EthService) GetBlockTransactionCountByNumber(blockNumber rpc.BlockNumber) (hexutil.Uint, error) {
	return withMetrics(
		e.metrics, "eth_getBlockTransactionCountByNumber",
		func() (hexutil.Uint, error) {
			return e.getBlockTransactionCountByNumber(blockNumber)
		},
	)
}

func (e *EthService) GetUncleCountByBlockHash(blockHash common.Hash) hexutil.Uint {
	return hexutil.Uint(0) // no uncles are ever generated
}

func (e *EthService) GetUncleCountByBlockNumber(blockNumber rpc.BlockNumber) hexutil.Uint {
	return hexutil.Uint(0) // no uncles are ever generated
}

func (e *EthService) GetUncleByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) map[string]interface{} {
	return nil // no uncles are ever generated
}

func (e *EthService) GetUncleByBlockNumberAndIndex(blockNumberOrTag rpc.BlockNumber, index hexutil.Uint) map[string]interface{} {
	return nil // no uncles are ever generated
}

func (e *EthService) Accounts() []common.Address {
	return e.accounts.Addresses()
}

// expressed in wei
// 1 Ether =
// 1_000_000_000 Gwei
// 1_000_000_000_000_000_000 wei
func (e *EthService) gasPrice() *hexutil.Big {
	return (*hexutil.Big)(e.evmChain.GasPrice())
}

func (e *EthService) GasPrice() (*hexutil.Big, error) {
	return withMetrics(
		e.metrics, "eth_gasPrice",
		func() (*hexutil.Big, error) {
			return e.gasPrice(), nil
		},
	)
}

func (e *EthService) Mining() bool {
	return false
}

func (e *EthService) Hashrate() float64 {
	return 0
}

func (e *EthService) Coinbase() common.Address {
	return common.Address{}
}

func (e *EthService) Syncing() bool {
	return false
}

func (e *EthService) GetCompilers() []string {
	return []string{}
}

func (e *EthService) sign(addr common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	account := e.accounts.Get(addr)
	if account == nil {
		return nil, errors.New("account is not unlocked")
	}

	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(msg))
	hash := hasher.Sum(nil)

	signed, err := crypto.Sign(hash, account)
	if err == nil {
		signed[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}
	return signed, err
}

func (e *EthService) Sign(addr common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	return withMetrics(
		e.metrics, "eth_sign",
		func() (hexutil.Bytes, error) {
			return e.sign(addr, data)
		},
	)
}

func (e *EthService) signTransaction(args *SendTxArgs) (hexutil.Bytes, error) {
	tx, err := e.parseTxArgs(args)
	if err != nil {
		return nil, err
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *EthService) SignTransaction(args *SendTxArgs) (hexutil.Bytes, error) {
	return withMetrics(
		e.metrics, "eth_signTransaction",
		func() (hexutil.Bytes, error) {
			return e.signTransaction(args)
		},
	)
}

func (e *EthService) sendTransaction(args *SendTxArgs) (common.Hash, error) {
	tx, err := e.parseTxArgs(args)
	if err != nil {
		return common.Hash{}, err
	}
	if err := e.evmChain.SendTransaction(tx); err != nil {
		return common.Hash{}, e.resolveError(err)
	}
	return tx.Hash(), nil
}

func (e *EthService) SendTransaction(args *SendTxArgs) (common.Hash, error) {
	return withMetrics(
		e.metrics, "eth_sendTransaction",
		func() (common.Hash, error) {
			return e.sendTransaction(args)
		},
	)
}

func (e *EthService) parseTxArgs(args *SendTxArgs) (*types.Transaction, error) {
	account := e.accounts.Get(args.From)
	if account == nil {
		return nil, errors.New("account is not unlocked")
	}
	if err := args.setDefaults(e); err != nil {
		return nil, err
	}
	signer, err := e.evmChain.Signer()
	if err != nil {
		return nil, err
	}
	return types.SignTx(args.toTransaction(), signer, account)
}

func (e *EthService) getLogs(q *RPCFilterQuery) ([]*types.Log, error) {
	logs, err := e.evmChain.Logs((*ethereum.FilterQuery)(q))
	if err != nil {
		return nil, e.resolveError(err)
	}
	return logs, nil
}

func (e *EthService) GetLogs(q *RPCFilterQuery) ([]*types.Log, error) {
	return withMetrics(
		e.metrics, "eth_getLogs",
		func() ([]*types.Log, error) {
			return e.getLogs(q)
		},
	)
}

// ChainID implements the eth_chainId method according to https://eips.ethereum.org/EIPS/eip-695
func (e *EthService) chainID() hexutil.Uint {
	chainID := e.evmChain.ChainID()
	return hexutil.Uint(chainID)
}

//nolint:revive // needs to be ChainId to match the interface
func (e *EthService) ChainId() (hexutil.Uint, error) {
	return withMetrics(
		e.metrics, "eth_chainId",
		func() (hexutil.Uint, error) {
			return e.chainID(), nil
		},
	)
}

func (e *EthService) NewHeads(ctx context.Context) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	rpcSub := notifier.CreateSubscription()

	go func() {
		headers := make(chan *types.Header)
		unsubscribe := e.evmChain.SubscribeNewHeads(headers)
		defer unsubscribe()

		for {
			select {
			case h := <-headers:
				_ = notifier.Notify(rpcSub.ID, h)
			case <-rpcSub.Err():
				return
			case <-notifier.Closed():
				return
			}
		}
	}()

	return rpcSub, nil
}

func (e *EthService) Logs(ctx context.Context, q *RPCFilterQuery) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	rpcSub := notifier.CreateSubscription()

	go func() {
		matchedLogs := make(chan []*types.Log)
		unsubscribe := e.evmChain.SubscribeLogs((*ethereum.FilterQuery)(q), matchedLogs)
		defer unsubscribe()

		for {
			select {
			case logs := <-matchedLogs:
				for _, log := range logs {
					_ = notifier.Notify(rpcSub.ID, log)
				}
			case <-rpcSub.Err():
				return
			case <-notifier.Closed():
				return
			}
		}
	}()

	return rpcSub, nil
}

/*
Not implemented:
func (e *EthService) NewFilter()
func (e *EthService) NewBlockFilter()
func (e *EthService) NewPendingTransactionFilter()
func (e *EthService) UninstallFilter()
func (e *EthService) GetFilterChanges()
func (e *EthService) GetFilterLogs()
func (e *EthService) SubmitWork()
func (e *EthService) GetWork()
func (e *EthService) SubmitHashrate()
func (e *EthService) CompileLLL()
func (e *EthService) CompileSolidity()
func (e *EthService) CompileSerpent()
*/

type NetService struct {
	chainID int
}

func NewNetService(chainID int) *NetService {
	return &NetService{
		chainID: chainID,
	}
}

func (s *NetService) Version() string {
	return strconv.Itoa(s.chainID)
}

func (s *NetService) Listening() bool         { return true }
func (s *NetService) PeerCount() hexutil.Uint { return 0 }

type Web3Service struct{}

func NewWeb3Service() *Web3Service {
	return &Web3Service{}
}

func (s *Web3Service) ClientVersion() string {
	return "wasp/evmproxy"
}

func (s *Web3Service) Sha3(input hexutil.Bytes) hexutil.Bytes {
	return crypto.Keccak256(input)
}

type TxPoolService struct{}

func NewTxPoolService() *TxPoolService {
	return &TxPoolService{}
}

func (s *TxPoolService) Content() map[string]map[string]map[string]*RPCTransaction {
	return map[string]map[string]map[string]*RPCTransaction{
		"pending": make(map[string]map[string]*RPCTransaction),
		"queued":  make(map[string]map[string]*RPCTransaction),
	}
}

func (s *TxPoolService) Inspect() map[string]map[string]map[string]string {
	return map[string]map[string]map[string]string{
		"pending": make(map[string]map[string]string),
		"queued":  make(map[string]map[string]string),
	}
}

func (s *TxPoolService) Status() map[string]hexutil.Uint {
	return map[string]hexutil.Uint{
		"pending": hexutil.Uint(0),
		"queued":  hexutil.Uint(0),
	}
}

type DebugService struct {
	evmChain *EVMChain
	metrics  *metrics.ChainWebAPIMetrics
}

func NewDebugService(evmChain *EVMChain, metrics *metrics.ChainWebAPIMetrics) *DebugService {
	return &DebugService{
		evmChain: evmChain,
		metrics:  metrics,
	}
}

func (d *DebugService) traceTransaction(txHash common.Hash, config *tracers.TraceConfig) (interface{}, error) {
	return d.evmChain.TraceTransaction(txHash, config)
}

func (d *DebugService) TraceTransaction(txHash common.Hash, config *tracers.TraceConfig) (interface{}, error) {
	return withMetrics(
		d.metrics, "debug_traceTransaction",
		func() (interface{}, error) {
			return d.traceTransaction(txHash, config)
		},
	)
}
