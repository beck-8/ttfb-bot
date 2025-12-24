package commands

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/filoz/ttfb-bot/pkg/pdp"
	"github.com/filoz/ttfb-bot/pkg/tester"
	"github.com/urfave/cli/v2"
)

type AggregateResult struct {
	ProviderID   uint64 `json:"provider_id"`
	ProviderName string `json:"provider_name"`
	Region       string `json:"region"`
	ServiceURL   string `json:"service_url"`

	Samples      int     `json:"samples"`
	SuccessCount int     `json:"success_count"`
	SuccessRate  float64 `json:"success_rate"`

	AvgTTFB_MS    int64   `json:"avg_ttfb_ms"`
	P95TTFB_MS    int64   `json:"p95_ttfb_ms"`
	AvgSpeed_MBps float64 `json:"avg_speed_mbps"`

	LocalRegion string `json:"local_region"` // Metadata
	Error       string `json:"error,omitempty"`
}

var RunCmd = &cli.Command{
	Name:  "run",
	Usage: "Run automated retrieval tests across providers",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Value:   "results.json",
			Usage:   "Output file path (supports .json or .csv)",
		},
		&cli.IntFlag{
			Name:    "concurrency",
			Aliases: []string{"c"},
			Value:   5,
			Usage:   "Number of concurrent provider tests",
		},
		&cli.Uint64Flag{
			Name:    "limit-providers",
			Aliases: []string{"limit"},
			Usage:   "Limit number of providers to test (0 = all)",
		},
		&cli.Uint64Flag{
			Name:  "scan-limit",
			Value: 1000,
			Usage: "Number of global datasets to scan for candidates",
		},
		&cli.Uint64Flag{
			Name:    "provider-id",
			Aliases: []string{"pid"},
			Usage:   "Run test on a specific provider only",
		},
		&cli.BoolFlag{
			Name:  "include-dev",
			Usage: "Include providers marked as 'dev'",
		},
		&cli.IntFlag{
			Name:    "samples",
			Aliases: []string{"s"},
			Value:   3,
			Usage:   "Number of distinct files to test per provider",
		},
		&cli.BoolFlag{
			Name:    "ping",
			Aliases: []string{"p"},
			Usage:   "Run a simple ping test to /pdp/ping instead of downloading pieces",
		},
	},
	Action: func(c *cli.Context) error {
		rand.Seed(time.Now().UnixNano())

		s, err := SetupServices(c)
		if err != nil {
			return err
		}
		defer s.Client.Close()

		// Get Local Region
		localRegion := getLocalRegion()
		fmt.Printf("Tester Region detected: %s\n", localRegion)

		fmt.Println("Step 1: Discovering active providers...")

		opts := &pdp.DiscoveryOptions{
			IncludeDev: c.Bool("include-dev"),
		}

		providers, err := s.Discovery.GetActiveProviders(c.Context, opts)
		if err != nil {
			return err
		}
		fmt.Printf("Found %d active providers.\n", len(providers))

		// Filter if --provider-id is set
		targetPid := c.Uint64("provider-id")
		if targetPid > 0 {
			var filtered []pdp.ProviderInfo
			for _, p := range providers {
				if p.ID == targetPid {
					filtered = append(filtered, p)
					break
				}
			}
			providers = filtered
			if len(providers) == 0 {
				return fmt.Errorf("provider %d not found in active list (check --include-dev)", targetPid)
			}
			fmt.Printf("Targeting single provider: %d (%s)\n", providers[0].ID, providers[0].Name)
		} else {
			limit := c.Uint64("limit-providers")
			if limit > 0 && limit < uint64(len(providers)) {
				providers = providers[:limit]
			}
		}

		samples := c.Int("samples")
		if samples < 1 {
			samples = 1
		}

		scanLimit := c.Uint64("scan-limit")
		if scanLimit == 0 {
			scanLimit = 1000
		}

		// Channel for results
		resultsChan := make(chan AggregateResult, len(providers))
		var wg sync.WaitGroup
		sem := make(chan struct{}, c.Int("concurrency"))

		fmt.Println("Step 2: Starting tests...")
		startTime := time.Now()

		for _, p := range providers {
			wg.Add(1)
			go func(prov pdp.ProviderInfo) {
				defer wg.Done()
				sem <- struct{}{}        // Acquire
				defer func() { <-sem }() // Release

				fmt.Printf("Testing Provider %d (%s) [%d samples]...\n", prov.ID, prov.Name, samples)

				agg := AggregateResult{
					ProviderID:   prov.ID,
					ProviderName: prov.Name,
					Region:       prov.Region,
					ServiceURL:   prov.ServiceURL,
					Samples:      samples,
					LocalRegion:  localRegion,
				}

				if c.Bool("ping") {
					targetURL := fmt.Sprintf("%s/pdp/ping", prov.ServiceURL)

					t := tester.NewTester()
					var ttfbs []float64
					var errors []string

					for i := 0; i < samples; i++ {
						metrics := t.DownloadPiece(targetURL)
						if metrics.Success {
							agg.SuccessCount++
							ttfbs = append(ttfbs, float64(metrics.TTFB.Milliseconds()))
						} else {
							errors = append(errors, metrics.Error)
						}
						// Small delay
						if i < samples-1 {
							time.Sleep(200 * time.Millisecond)
						}
					}

					if agg.SuccessCount > 0 {
						agg.AvgTTFB_MS = int64(avg(ttfbs))
						agg.P95TTFB_MS = int64(percentile(ttfbs, 95))
						agg.SuccessRate = float64(agg.SuccessCount) / float64(samples)
					}

					if len(errors) > 0 {
						agg.Error = strings.Join(removeDuplicateStrings(errors), "; ")
					}

					resultsChan <- agg
					return
				}

				// Scan for potential pieces
				datasets, err := s.Dataset.GetDatasetsForProvider(context.Background(), prov.ID, scanLimit)
				if err != nil {
					agg.Error = fmt.Sprintf("Scan failed: %v", err)
					resultsChan <- agg
					return
				}

				// Filter candidates with valid CIDs
				var candidates []pdp.DatasetInfo
				for _, ds := range datasets {
					if ds.PieceCID != "" && !strings.HasPrefix(ds.PieceCID, "0x") {
						candidates = append(candidates, ds)
					}
				}

				if len(candidates) == 0 {
					agg.Error = "No valid sample pieces found"
					resultsChan <- agg
					return
				}

				// Shuffle and pick N
				rand.Shuffle(len(candidates), func(i, j int) { candidates[i], candidates[j] = candidates[j], candidates[i] })

				targetCount := samples
				if len(candidates) < targetCount {
					targetCount = len(candidates)
					fmt.Printf("Warning: Provider %d only has %d valid pieces (requested %d)\n", prov.ID, len(candidates), samples)
				}

				selected := candidates[:targetCount]
				agg.Samples = targetCount

				// Perform samples
				var ttfbs []float64
				var speeds []float64
				var errors []string

				t := tester.NewTester()

				for _, ds := range selected {
					targetURL := fmt.Sprintf("%s/piece/%s", prov.ServiceURL, ds.PieceCID)

					metrics := t.DownloadPiece(targetURL)
					if metrics.Success {
						agg.SuccessCount++
						ttfbs = append(ttfbs, float64(metrics.TTFB.Milliseconds()))
						speeds = append(speeds, metrics.Speed)
					} else {
						errors = append(errors, metrics.Error)
					}
					// Small delay between samples
					time.Sleep(200 * time.Millisecond)
				}

				if agg.SuccessCount > 0 {
					agg.AvgTTFB_MS = int64(avg(ttfbs))
					agg.P95TTFB_MS = int64(percentile(ttfbs, 95))
					agg.AvgSpeed_MBps = avg(speeds)
					agg.SuccessRate = float64(agg.SuccessCount) / float64(agg.Samples)
				}

				if len(errors) > 0 {
					// Join unique errors
					uniqueErrors := removeDuplicateStrings(errors)
					agg.Error = strings.Join(uniqueErrors, "; ")
				}

				resultsChan <- agg
			}(p)
		}

		wg.Wait()
		close(resultsChan)

		// Collect results
		var results []AggregateResult
		for r := range resultsChan {
			results = append(results, r)
		}

		// Output
		fmt.Printf("\nTests completed in %v\n", time.Since(startTime))
		return saveResults(results, c.String("output"))
	},
}

func saveResults(results []AggregateResult, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if strings.HasSuffix(filename, ".csv") {
		w := csv.NewWriter(f)
		defer w.Flush()

		// Header
		w.Write([]string{
			"ProviderID", "Name", "Region", "LocalRegion",
			"Samples", "SuccessRate",
			"AvgTTFB(ms)", "P95TTFB(ms)", "AvgSpeed(MB/s)",
			"Error",
		})
		for _, r := range results {
			w.Write([]string{
				fmt.Sprintf("%d", r.ProviderID),
				r.ProviderName,
				r.Region,
				r.LocalRegion,
				fmt.Sprintf("%d", r.Samples),
				fmt.Sprintf("%.2f", r.SuccessRate),
				fmt.Sprintf("%d", r.AvgTTFB_MS),
				fmt.Sprintf("%d", r.P95TTFB_MS),
				fmt.Sprintf("%.2f", r.AvgSpeed_MBps),
				r.Error,
			})
		}
	} else {
		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		return enc.Encode(results)
	}

	fmt.Printf("Results saved to %s\n", filename)
	return nil
}

// Helpers

func avg(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

func percentile(nums []float64, p float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sort.Float64s(nums)
	idx := int(math.Ceil(float64(len(nums))*p/100)) - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(nums) {
		idx = len(nums) - 1
	}
	return nums[idx]
}

func removeDuplicateStrings(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getLocalRegion() string {
	// Simple HTTP request to ip-api.com
	// Returns JSON: {"countryCode":"US", "regionName":"California", ...}
	// We'll format it similar to our encoded regions: C=US;ST=California
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json")
	if err != nil {
		return "Unknown (Network Error)"
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "Unknown (Parse Error)"
	}

	cc, _ := data["countryCode"].(string)
	region, _ := data["regionName"].(string)
	city, _ := data["city"].(string)

	return fmt.Sprintf("C=%s;ST=%s;L=%s", cc, region, city)
}
