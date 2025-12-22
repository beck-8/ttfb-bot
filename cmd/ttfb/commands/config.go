package commands

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filoz/ttfb-bot/pkg/contracts/registry"
	"github.com/filoz/ttfb-bot/pkg/contracts/warmstorage"
	"github.com/filoz/ttfb-bot/pkg/pdp"
	"github.com/urfave/cli/v2"
)

// Network Config
const (
	CalibRPC         = "https://api.calibration.node.glif.io/rpc/v1"
	CalibWarmStorage = "0x02925630df557F957f70E112bA06e50965417CA0"

	MainnetRPC         = "https://api.node.glif.io/rpc/v1"
	MainnetWarmStorage = "0x8408502033C418E1bbC97cE9ac48E5528F371A9f"
)

type Services struct {
	Client           *ethclient.Client
	Discovery        *pdp.DiscoveryService
	Dataset          *pdp.DatasetService
	WarmStorageAddr  common.Address
	RegistryContract *registry.ServiceProviderRegistry
}

func SetupServices(c *cli.Context) (*Services, error) {
	network := c.String("network")
	rpcURL := c.String("rpc")

	var warmStorageHex string

	if network == "mainnet" {
		if rpcURL == "" {
			rpcURL = MainnetRPC
		}
		warmStorageHex = MainnetWarmStorage
	} else {
		// Default to calibration
		if rpcURL == "" {
			rpcURL = CalibRPC
		}
		warmStorageHex = CalibWarmStorage
	}

	// Allow override via env or flag, but simple "network" flag is easier for user
	if c.String("warm-storage") != "" {
		warmStorageHex = c.String("warm-storage")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("connect rpc: %w", err)
	}

	warmStorageAddr := common.HexToAddress(warmStorageHex)

	// Resolve Registry
	ws, err := warmstorage.NewWarmStorageService(warmStorageAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind warm storage: %w", err)
	}

	registryAddr, err := ws.ServiceProviderRegistry(nil)
	if err != nil {
		return nil, fmt.Errorf("resolve registry: %w", err)
	}

	disco, err := pdp.NewDiscoveryService(client, registryAddr)
	if err != nil {
		return nil, fmt.Errorf("init discovery: %w", err)
	}

	ds, err := pdp.NewDatasetService(client, warmStorageAddr)
	if err != nil {
		return nil, fmt.Errorf("init dataset: %w", err)
	}

	// Create raw registry binding if needed explicitly, but disco has it
	reg, err := registry.NewServiceProviderRegistry(registryAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind registry: %w", err)
	}

	return &Services{
		Client:           client,
		Discovery:        disco,
		Dataset:          ds,
		WarmStorageAddr:  warmStorageAddr,
		RegistryContract: reg,
	}, nil
}
