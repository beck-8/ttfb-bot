package main

import (
	"log"
	"os"

	"github.com/filoz/ttfb-bot/cmd/ttfb/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ttfb",
		Usage: "Filecoin Retrieval Performance Tester",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "network",
				Aliases: []string{"n"},
				Value:   "calibration",
				Usage:   "Network to use (calibration, mainnet)",
			},
			&cli.StringFlag{
				Name:  "rpc",
				Usage: "Custom RPC URL (optional)",
			},
			&cli.StringFlag{
				Name:  "warm-storage",
				Usage: "Custom WarmStorage contract address (optional)",
			},
		},
		Commands: []*cli.Command{
			commands.ListCmd,
			commands.DownloadCmd,
			commands.RunCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
