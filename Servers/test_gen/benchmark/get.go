package benchmark

import (
	"fmt"
	"net/http"
	"sync"
	"test_gen/common"
	"test_gen/types"
	"time"
)

func GetRequestBenchMark(numRequests int) {
	apiURL := "http://localhost:9010/profile/get?profileID="
	numRequests += 500

	var (
		wg               sync.WaitGroup
		successfulBursts int
		successCount     int
		clientErrorCount int
		serverErrorCount int
		totalLatency     time.Duration
		maxLatency       time.Duration
		minLatency       time.Duration
		averageLatency   time.Duration
		individualBursts []types.BurstStats
	)

	burstSize := 30
	numBursts := (numRequests + burstSize - 1) / burstSize

	startTime := time.Now()
	for burst := 0; burst < numBursts; burst++ {
		latencies := make([]time.Duration, burstSize)
		stats := types.BurstStats{BurstNumber: burst + 1}

		wg.Add(burstSize)
		for i := 0; i < burstSize && burst*burstSize+i < numRequests; i++ {
			go func(index int) {
				defer wg.Done()
				reqStart := time.Now()
				resp, err := http.Get(fmt.Sprintf("%s%d", apiURL, burst*burstSize+index+1))
				if err != nil {
					fmt.Printf("Error sending request: %v\n", err)
					return
				}
				latency := time.Since(reqStart)

				// Update burst statistics
				switch {
				case resp.StatusCode >= 200 && resp.StatusCode < 300:
					stats.SuccessfulCount++
				case resp.StatusCode >= 400 && resp.StatusCode < 500:
					stats.ClientErrors++
				case resp.StatusCode >= 500:
					stats.ServerErrors++
				}
				latencies[index] = latency

				resp.Body.Close()
			}(i)
		}
		wg.Wait()

		// Calculate individual latency statistics for this burst
		stats.TotalLatency, stats.MaxLatency, stats.MinLatency, stats.AverageLatency = common.CalculateLatencyStats(latencies)

		// Update overall statistics
		successCount += stats.SuccessfulCount
		clientErrorCount += stats.ClientErrors
		serverErrorCount += stats.ServerErrors
		totalLatency += stats.TotalLatency
		if stats.MaxLatency > maxLatency {
			maxLatency = stats.MaxLatency
		}
		if stats.MinLatency < minLatency || minLatency == 0 {
			minLatency = stats.MinLatency
		}
		averageLatency += stats.AverageLatency
		individualBursts = append(individualBursts, stats)

		// Check if the burst was successful
		if stats.SuccessfulCount == burstSize {
			successfulBursts++
		}
	}

	// Calculate overall latency statistics
	averageLatency /= time.Duration(numRequests)

	elapsed := time.Since(startTime)

	// Print individual burst statistics
	fmt.Println("Individual Burst Statistics:")
	for _, burst := range individualBursts {
		fmt.Printf("Burst %d: Successful: %d, Client Errors: %d, Server Errors: %d, Total Latency: %v, Max Latency: %v, Min Latency: %v, Average Latency: %v\n",
			burst.BurstNumber, burst.SuccessfulCount, burst.ClientErrors, burst.ServerErrors, burst.TotalLatency, burst.MaxLatency, burst.MinLatency, burst.AverageLatency)

		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

	// Print overall performance metrics
	fmt.Printf("\nTotal requests sent: %d\n", numRequests)
	fmt.Printf("Total time taken: %v\n", elapsed)
	fmt.Printf("Average time per request: %v\n", elapsed/time.Duration(numRequests))
	fmt.Printf("Successful bursts: %d\n", successfulBursts)
	fmt.Printf("Successful responses: %d\n", successCount)
	fmt.Printf("Client error responses: %d\n", clientErrorCount)
	fmt.Printf("Server error responses: %d\n", serverErrorCount)
	fmt.Printf("Max latency: %v\n", maxLatency)
	fmt.Printf("Min latency: %v\n", minLatency)
	fmt.Printf("Average latency: %v\n", averageLatency)

	err := common.WriteBurstStatsToCSV(individualBursts, "burst_stats.csv")
	if err != nil {
		fmt.Printf("Error writing CSV file: %v\n", err)
	}
}
