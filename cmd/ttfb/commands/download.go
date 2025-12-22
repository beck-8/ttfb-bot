package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/filoz/ttfb-bot/pkg/pdp"
	"github.com/filoz/ttfb-bot/pkg/tester"
	"github.com/urfave/cli/v2"
)

var DownloadCmd = &cli.Command{
	Name:      "download",
	Usage:     "Test download performance",
	ArgsUsage: "[url]",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "provider-id",
			Usage: "Auto-find a piece from this provider ID",
		},
	},
	Action: func(c *cli.Context) error {
		targetURL := c.Args().First()
		pid := c.Uint64("provider-id")

		if targetURL == "" && pid == 0 {
			return fmt.Errorf("must provide url OR --provider-id")
		}

		if targetURL == "" && pid > 0 {
			// Auto discovery logic
			s, err := SetupServices(c)
			if err != nil {
				return err
			}
			defer s.Client.Close()

			// Get Provider Info
			// TODO: Optimise by adding GetProviderById to discovery
			provs, _ := s.Discovery.GetActiveProviders(context.Background())
			var selectedProv *pdp.ProviderInfo
			for _, p := range provs {
				if p.ID == pid {
					selectedProv = &p
					break
				}
			}
			if selectedProv == nil {
				return fmt.Errorf("provider %d not found or inactive", pid)
			}

			fmt.Printf("Searching piece for provider %s (%s)...\n", selectedProv.Name, selectedProv.ServiceURL)
			datasets, err := s.Dataset.GetDatasetsForProvider(context.Background(), pid, 500)
			if err != nil || len(datasets) == 0 {
				return fmt.Errorf("no datasets found for provider")
			}

			for _, ds := range datasets {
				if ds.PieceCID != "" && !strings.HasPrefix(ds.PieceCID, "0x") {
					targetURL = fmt.Sprintf("%s/piece/%s", selectedProv.ServiceURL, ds.PieceCID)
					fmt.Printf("Selected Piece %s from Dataset %d\n", ds.PieceCID, ds.ID)
					break
				}
			}
			if targetURL == "" {
				return fmt.Errorf("no valid pieces found for provider (all CIDs were raw hex?)")
			}
		}

		fmt.Printf("Testing Download: %s\n", targetURL)
		t := tester.NewTester()
		metrics := t.DownloadPiece(targetURL)

		if metrics.Success {
			fmt.Printf("SUCCESS!\nTTFB: %v\nTotal: %v\nSize: %d bytes\nSpeed: %.2f MB/s\n",
				metrics.TTFB, metrics.TotalTime, metrics.Size, metrics.Speed)
		} else {
			fmt.Printf("FAILURE: %s\nTotal: %v\n", metrics.Error, metrics.TotalTime)
			return fmt.Errorf("download failed: %s", metrics.Error)
		}

		return nil
	},
}
