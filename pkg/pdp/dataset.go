package pdp

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filoz/ttfb-bot/pkg/contracts/verifier"
	"github.com/filoz/ttfb-bot/pkg/contracts/view"
	"github.com/filoz/ttfb-bot/pkg/contracts/warmstorage"
	"github.com/ipfs/go-cid"
)

type DatasetService struct {
	client          *ethclient.Client
	warmStorage     *warmstorage.WarmStorageService
	warmStorageView *view.WarmStorageView
	pdpVerifier     *verifier.PDPVerifier
}

type DatasetInfo struct {
	ID         uint64
	ProviderID uint64
	PieceCID   string // First piece CID for testing
	PieceSize  uint64 // Not available in View directly easily without looking up PieceMetaData, assumes default for now or we skip size
}

func NewDatasetService(client *ethclient.Client, warmStorageAddr common.Address) (*DatasetService, error) {
	ws, err := warmstorage.NewWarmStorageService(warmStorageAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind warm storage: %w", err)
	}

	viewAddr, err := ws.ViewContractAddress(nil)
	if err != nil {
		return nil, fmt.Errorf("get view addr: %w", err)
	}

	verifierAddr, err := ws.PdpVerifierAddress(nil)
	if err != nil {
		return nil, fmt.Errorf("get verifier addr: %w", err)
	}

	vContract, err := view.NewWarmStorageView(viewAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind view: %w", err)
	}

	verContract, err := verifier.NewPDPVerifier(verifierAddr, client)
	if err != nil {
		return nil, fmt.Errorf("bind verifier: %w", err)
	}

	return &DatasetService{
		client:          client,
		warmStorage:     ws,
		warmStorageView: vContract,
		pdpVerifier:     verContract,
	}, nil
}

func (s *DatasetService) GetDatasetsForProvider(ctx context.Context, providerID uint64, scanDepth uint64) ([]DatasetInfo, error) {
	// 1. Get total dataset count
	nextID, err := s.pdpVerifier.GetNextDataSetId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("get next dataset id: %w", err)
	}

	maxID := nextID

	var datasets []DatasetInfo
	var mu sync.Mutex

	startID := uint64(1)
	if maxID > scanDepth {
		startID = maxID - scanDepth
	}

	// Use a semaphore for concurrency
	sem := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for id := maxID - 1; id >= startID; id-- {
		// Skip invalid ID 0
		if id == 0 {
			continue
		}

		wg.Add(1)
		sem <- struct{}{}

		go func(dataID uint64) {
			defer wg.Done()
			defer func() { <-sem }()

			info, err := s.warmStorageView.GetDataSet(&bind.CallOpts{Context: ctx}, big.NewInt(int64(dataID)))
			if err != nil {
				// Failed to fetch
				// fmt.Printf("Debug: Failed to get dataset %d: %v\n", dataID, err)
				return
			}

			if info.ProviderId.Uint64() == providerID && info.PdpRailId.Uint64() != 0 {
				// Match!
				pieceCID := ""

				nextPieceID, err := s.pdpVerifier.GetNextPieceId(&bind.CallOpts{Context: ctx}, big.NewInt(int64(dataID)))
				if err == nil && nextPieceID.Uint64() > 0 {
					cidInfo, err := s.pdpVerifier.GetPieceCid(&bind.CallOpts{Context: ctx}, big.NewInt(int64(dataID)), big.NewInt(0))
					if err == nil && len(cidInfo.Data) > 0 {
						if c, err := cid.Cast(cidInfo.Data); err == nil {
							pieceCID = c.String()
						} else {
							pieceCID = fmt.Sprintf("0x%x", cidInfo.Data)
						}
					}
				}

				mu.Lock()
				datasets = append(datasets, DatasetInfo{
					ID:         dataID,
					ProviderID: providerID,
					PieceCID:   pieceCID,
				})
				mu.Unlock()
			}
		}(id)
	}

	wg.Wait()
	return datasets, nil
}
