package services

import (
	"errors"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/labstack/echo/v4"

	hivedb "github.com/iotaledger/hive.go/kvstore/database"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/evm/jsonrpc"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/publisher"
	"github.com/iotaledger/wasp/packages/webapi/interfaces"
)

type chainServer struct {
	backend *jsonrpc.WaspEVMBackend
	rpc     *rpc.Server
}

type EVMService struct {
	evmBackendMutex sync.Mutex
	evmChainServers map[isc.ChainID]*chainServer

	chainService    interfaces.ChainService
	networkProvider peering.NetworkProvider
	publisher       *publisher.Publisher
	isArchiveNode   bool
	indexDbPath     string
	metrics         *metrics.ChainMetricsProvider
	log             *logger.Logger
}

func NewEVMService(
	chainService interfaces.ChainService,
	networkProvider peering.NetworkProvider,
	pub *publisher.Publisher,
	isArchiveNode bool,
	indexDbPath string,
	metrics *metrics.ChainMetricsProvider,
	log *logger.Logger,
) interfaces.EVMService {
	return &EVMService{
		chainService:    chainService,
		evmChainServers: map[isc.ChainID]*chainServer{},
		evmBackendMutex: sync.Mutex{},
		networkProvider: networkProvider,
		publisher:       pub,
		isArchiveNode:   isArchiveNode,
		indexDbPath:     indexDbPath,
		metrics:         metrics,
		log:             log,
	}
}

func (e *EVMService) getEVMBackend(chainID isc.ChainID) (*chainServer, error) {
	e.evmBackendMutex.Lock()
	defer e.evmBackendMutex.Unlock()

	if e.evmChainServers[chainID] != nil {
		return e.evmChainServers[chainID], nil
	}

	chain, err := e.chainService.GetChainByID(chainID)
	if err != nil {
		return nil, err
	}

	nodePubKey := e.networkProvider.Self().PubKey()
	backend := jsonrpc.NewWaspEVMBackend(chain, nodePubKey, parameters.L1().BaseToken)

	srv, err := jsonrpc.NewServer(
		jsonrpc.NewEVMChain(backend, e.publisher, e.isArchiveNode, hivedb.EngineRocksDB, e.indexDbPath, e.log.Named("EVMChain")),
		jsonrpc.NewAccountManager(nil),
		e.metrics.GetChainMetrics(chainID).WebAPI,
	)
	if err != nil {
		return nil, err
	}

	e.evmChainServers[chainID] = &chainServer{
		backend: backend,
		rpc:     srv,
	}

	return e.evmChainServers[chainID], nil
}

func (e *EVMService) HandleJSONRPC(chainID isc.ChainID, request *http.Request, response *echo.Response) error {
	evmServer, err := e.getEVMBackend(chainID)
	if err != nil {
		return err
	}

	evmServer.rpc.ServeHTTP(response, request)

	return nil
}

func (e *EVMService) HandleWebsocket(chainID isc.ChainID, request *http.Request, response *echo.Response) error {
	evmServer, err := e.getEVMBackend(chainID)
	if err != nil {
		return err
	}

	allowedOrigins := []string{"*"}
	evmServer.rpc.WebsocketHandler(allowedOrigins).ServeHTTP(response, request)

	return nil
}

func (e *EVMService) GetRequestID(chainID isc.ChainID, hash string) (isc.RequestID, error) {
	evmServer, err := e.getEVMBackend(chainID)
	if err != nil {
		return isc.RequestID{}, err
	}

	txHash := common.HexToHash(hash)
	reqID, ok := evmServer.backend.RequestIDByTransactionHash(txHash)
	if !ok {
		return isc.RequestID{}, errors.New("request id not found")
	}

	return reqID, nil
}
