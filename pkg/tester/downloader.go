package tester

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type DownloadMetrics struct {
	TTFB       time.Duration
	TotalTime  time.Duration
	Size       int64
	Speed      float64 // MB/s
	Success    bool
	StatusCode int
	Error      string
}

type Tester struct {
	client *http.Client
}

func NewTester() *Tester {
	return &Tester{
		client: &http.Client{
			Timeout: 30 * time.Second, // Adjustable
			Transport: &http.Transport{
				DisableCompression: true,
			},
		},
	}
}

func (t *Tester) DownloadPiece(url string) DownloadMetrics {
	metrics := DownloadMetrics{
		Success: false,
	}

	start := time.Now()
	resp, err := t.client.Get(url)
	if err != nil {
		metrics.Error = err.Error()
		metrics.TotalTime = time.Since(start)
		return metrics
	}
	defer resp.Body.Close()

	metrics.StatusCode = resp.StatusCode
	if resp.StatusCode != http.StatusOK {
		metrics.Error = fmt.Sprintf("HTTP Status %d", resp.StatusCode)
		metrics.TotalTime = time.Since(start)
		return metrics
	}

	// TTFB: Time until Body.Read returns first byte?
	// Actually, client.Get returns after headers are received.
	// So time.Now() - start is roughly TTFB (connect + headers).
	// To be precise on "First Byte of Body", we need to read 1 byte.

	ttfb := time.Since(start)
	metrics.TTFB = ttfb

	// Read body
	// We can discard data if we don't save it.
	// Or we can hash it if verification needed. For now just drain.
	written, err := io.Copy(io.Discard, resp.Body)
	totalTime := time.Since(start)

	metrics.TotalTime = totalTime
	metrics.Size = written
	if err != nil {
		metrics.Error = fmt.Sprintf("Read error: %v", err)
		return metrics
	}

	metrics.Success = true

	// Speed in MB/s
	if totalTime.Seconds() > 0 {
		metrics.Speed = float64(written) / (1024 * 1024) / totalTime.Seconds()
	}

	return metrics
}
