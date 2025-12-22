package pdp

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filoz/ttfb-bot/pkg/contracts/registry"
)

type DiscoveryService struct {
	client           *ethclient.Client
	registryContract *registry.ServiceProviderRegistry
}

func NewDiscoveryService(client *ethclient.Client, registryAddr common.Address) (*DiscoveryService, error) {
	reg, err := registry.NewServiceProviderRegistry(registryAddr, client)
	if err != nil {
		return nil, err
	}
	return &DiscoveryService{
		client:           client,
		registryContract: reg,
	}, nil
}

type ProviderInfo struct {
	ID          uint64
	Address     common.Address
	ServiceURL  string
	IsActive    bool
	Region      string
	Description string
	Name        string
}

func (s *DiscoveryService) GetActiveProviders(ctx context.Context) ([]ProviderInfo, error) {
	count, err := s.registryContract.GetProviderCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get provider count: %w", err)
	}

	var providers []ProviderInfo

	// Loop through providers
	// Note: optimization possible with concurrency
	for i := uint64(1); i <= count.Uint64(); i++ {
		// 0 is PDP Product Type
		info, err := s.registryContract.GetProviderWithProduct(&bind.CallOpts{Context: ctx}, big.NewInt(int64(i)), 0)
		if err != nil {
			// Provider might not have the product or other error, skip
			continue
		}

		if !info.ProviderInfo.IsActive || !info.Product.IsActive {
			continue
		}

		serviceURL := ""
		region := ""

		// Decode capabilities
		// Capabilities are keys (string[]) and values (bytes[])
		// We match them by index
		for idx, key := range info.Product.CapabilityKeys {
			if idx >= len(info.ProductCapabilityValues) {
				break
			}
			valBytes := info.ProductCapabilityValues[idx]

			switch key {
			case "serviceURL":
				serviceURL = string(valBytes)
			case "location":
				region = string(valBytes)
			}
		}

		// Clean up Service URL (remove trailing slash, etc)
		serviceURL = strings.TrimSpace(serviceURL)
		serviceURL = strings.TrimRight(serviceURL, "/")

		if serviceURL != "" {
			providers = append(providers, ProviderInfo{
				ID:          i,
				Address:     info.ProviderInfo.ServiceProvider,
				ServiceURL:  serviceURL,
				IsActive:    true,
				Region:      region,
				Name:        info.ProviderInfo.Name,
				Description: info.ProviderInfo.Description,
			})
		}
	}

	return providers, nil
}
