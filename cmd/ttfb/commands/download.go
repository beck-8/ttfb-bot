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
	Usage:     "Test download from a URL or Provider",
	ArgsUsage: "[url]",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:    "provider-id",
			Aliases: []string{"p", "pid"},
			Usage:   "Provider ID to auto-discover a piece from",
		},
		&cli.BoolFlag{
			Name:  "include-dev",
			Usage: "Include providers marked as 'dev' (when auto-selecting)",
		},
	},
	Action: func(c *cli.Context) error {
		url := c.Args().First()
		pid := c.Uint64("provider-id")

		// Case 1: Direct URL provided
		if url != "" {
			return runDownloadTest(url)
		}

		// Case 2: Provider ID or Auto-Discovery needed
		s, err := SetupServices(c)
		if err != nil {
			return err
		}
		defer s.Client.Close()

		if pid == 0 {
			fmt.Println("No URL or Provider ID specified, auto-selecting a provider...")
			includeDev := c.Bool("include-dev")
			providers, err := s.Discovery.GetActiveProviders(c.Context, &pdp.DiscoveryOptions{IncludeDev: includeDev})
			if err != nil || len(providers) == 0 {
				return fmt.Errorf("no active providers found")
			}
			pid = providers[0].ID
			fmt.Printf("Selected Provider %d (%s)\n", pid, providers[0].Name)

			// Update url prefix
			if providers[0].ServiceURL != "" {
				url = providers[0].ServiceURL
			}
		} else {
			// Need to resolve service URL for this PID if we don't have it?
			// The DatasetService needs PID to find datasets, but to construct URL we need ServiceURL.
			// Discovery Service gives us ServiceURL.
			// Let's refetch provider info
			// This is a bit inefficient but safe.
			includeDev := c.Bool("include-dev")
			providers, err := s.Discovery.GetActiveProviders(c.Context, &pdp.DiscoveryOptions{IncludeDev: includeDev})
			if err != nil {
				return err
			}
			found := false
			for _, p := range providers {
				if p.ID == pid {
					url = p.ServiceURL
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("provider %d not found or inactive", pid)
			}
		}

		// Find a piece
		fmt.Printf("Finding a valid piece for Provider %d...\n", pid)
		datasets, err := s.Dataset.GetDatasetsForProvider(context.Background(), pid, 50)
		if err != nil {
			return err
		}

		var targetPiece string
		for _, ds := range datasets {
			if ds.PieceCID != "" && !strings.HasPrefix(ds.PieceCID, "0x") {
				targetPiece = ds.PieceCID
				fmt.Printf("Found valid piece: %s (Dataset %d)\n", ds.PieceCID, ds.ID)
				break
			}
		}

		if targetPiece == "" {
			return fmt.Errorf("no valid pieces found for provider %d", pid)
		}

		finalURL := fmt.Sprintf("%s/piece/%s", url, targetPiece)
		return runDownloadTest(finalURL)
	},
}

func runDownloadTest(url string) error {
	fmt.Printf("Testing Download: %s\n", url)

	t := tester.NewTester()
	m := t.DownloadPiece(url)

	if m.Success {
		fmt.Println("SUCCESS!")
		fmt.Printf("TTFB: %v\n", m.TTFB)
		fmt.Printf("Total: %v\n", m.TotalTime)
		fmt.Printf("Size: %d bytes\n", m.Size)
		fmt.Printf("Speed: %.2f MB/s\n", m.Speed)
	} else {
		fmt.Println("FAILED!")
		fmt.Printf("Error: %s\n", m.Error)
		fmt.Printf("Status: %d\n", m.StatusCode)
	}
	return nil
}
