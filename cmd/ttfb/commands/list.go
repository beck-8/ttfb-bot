package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List providers or datasets",
	Subcommands: []*cli.Command{
		{
			Name:  "providers",
			Usage: "List active PDP providers",
			Action: func(c *cli.Context) error {
				s, err := SetupServices(c)
				if err != nil {
					return err
				}
				defer s.Client.Close()

				fmt.Println("Fetching active PDP providers...")
				providers, err := s.Discovery.GetActiveProviders(c.Context)
				if err != nil {
					return err
				}

				fmt.Printf("Found %d active providers:\n", len(providers))
				for _, p := range providers {
					fmt.Printf("- [%d] %s (%s) URL: %s Region: %s\n", p.ID, p.Name, p.Address.Hex(), p.ServiceURL, p.Region)
				}
				return nil
			},
		},
		{
			Name:  "datasets",
			Usage: "List datasets for a provider",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "provider-id",
					Aliases: []string{"p", "pid"},
					Usage:   "Provider ID to scan",
				},
				&cli.Uint64Flag{
					Name:    "limit",
					Aliases: []string{"l"},
					Value:   200,
					Usage:   "Number of recent datasets to scan",
				},
			},
			Action: func(c *cli.Context) error {
				s, err := SetupServices(c)
				if err != nil {
					return err
				}
				defer s.Client.Close()

				pid := c.Uint64("provider-id")
				limit := c.Uint64("limit")

				if pid == 0 {
					fmt.Println("No provider ID specified, picking first active provider...")
					providers, err := s.Discovery.GetActiveProviders(c.Context)
					if err != nil || len(providers) == 0 {
						return fmt.Errorf("no active providers found: %v", err)
					}
					pid = providers[0].ID
					fmt.Printf("Selected Provider %d (%s)\n", pid, providers[0].Name)
				}

				fmt.Printf("Scanning datasets for Provider %d (limit: last %d)...\n", pid, limit)
				datasets, err := s.Dataset.GetDatasetsForProvider(context.Background(), pid, limit)
				if err != nil {
					return err
				}

				fmt.Printf("Found %d datasets:\n", len(datasets))
				for _, ds := range datasets {
					fmt.Printf("- Dataset %d | Piece %s\n", ds.ID, ds.PieceCID)
				}
				return nil
			},
		},
	},
}
