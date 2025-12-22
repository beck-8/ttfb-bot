# Filecoin PDP Benchmark Tool (ttfb)

A powerful CLI tool to benchmark Time-To-First-Byte (TTFB) and download speeds of Filecoin Storage Providers via the PDP (Provable Data Possession) protocol.

## Features

*   **Automated Discovery**: Automatically discovers active PDP providers on-chain.
*   **Performance Benchmarking**: Measures TTFB, Download Speed, and Connectivity.
*   **Multi-Sampling**: Performs multiple random sample tests per provider to ensure statistical accuracy (Avg, P95).
*   **Advanced Filtering**: Filter by Provider ID, Environment (Dev/Prod), and more.
*   **Visualization**: Includes a Python script to generate professional performance charts.
*   **Region Awareness**: Detects and records the tester's geographic location for context.

## Build

Prerequisites: **Go 1.21+**

```bash
# Clone the repo
git clone https://github.com/beck-8/pdp-bot.git
cd pdp-bot

# Build the binary (outputs to ./ttfb)
make build
```

## Usage

### 1. Run Automated Benchmarks

The `run` command is the main entry point. It discovers providers, selects random pieces, and runs download tests.

```bash
# Basic run (default: 5 concurrent, 3 samples per provider)
./ttfb run

# Output to CSV (recommended for visualization)
./ttfb run -o report.csv

# Test a specific provider (e.g., Provider ID 1)
./ttfb run --pid 1

# Include "Dev" nodes (nodes marked as 'dev' status are hidden by default)
./ttfb run --include-dev

# Increase scan limit for providers with few recent uploads
# (Scans last 5000 datasets to find valid pieces)
./ttfb run --scan-limit 5000
```

### 2. List Resources

```bash
# List all active providers
./ttfb list providers --include-dev

# List recent datasets for a specific provider
./ttfb list datasets --pid 1 --limit 100
```

### 3. Single Download Test

```bash
# Test downloadability of a specific Piece URL
./ttfb download http://provider-url/piece/bafk...

# Auto-discover a piece from a provider and test it
./ttfb download --pid 1
```

## Visualization

A Python script is provided to visualize the benchmark results (`.csv`).

**Prerequisites**: Python 3, pandas, matplotlib, seaborn.

```bash
# 1. Create a virtual environment (optional but recommended)
python3 -m venv venv
source venv/bin/activate

# 2. Install dependencies
pip install pandas matplotlib seaborn

# 3. Generate Chart
# This will produce 'performance_summary_clean.png'
python3 plot_results.py report.csv
```

## Flags Reference

| Flag | Commands | Description |
|------|----------|-------------|
| `--pid`, `--provider-id` | `run`, `list`, `download` | Target a specific Provider ID. |
| `--include-dev` | `run`, `list`, `download` | Include providers with `service_status="dev"`. |
| `--samples`, `-s` | `run` | Number of distinct random files to test per provider. |
| `--scan-limit` | `run` | Number of global datasets to scan for candidates (default: 1000). |
| `--concurrency`, `-c` | `run` | Number of providers to test in parallel. |

## License

MIT
