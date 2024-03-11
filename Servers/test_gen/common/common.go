package common

import (
	"encoding/csv"
	"os"
	"strconv"
	"test_gen/types"
	"time"
)

func CalculateLatencyStats(latencies []time.Duration) (total, max, min, avg time.Duration) {
	if len(latencies) == 0 {
		return 0, 0, 0, 0
	}

	max = latencies[0]
	min = latencies[0]
	for _, l := range latencies {
		total += l
		if l > max {
			max = l
		}
		if l < min {
			min = l
		}
	}
	avg = total / time.Duration(len(latencies))
	return total, max, min, avg
}

func WriteBurstStatsToCSV(individualBursts []types.BurstStats, filename string) error {
	// Create or open the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := []string{"BurstNumber", "SuccessfulCount", "ClientErrors", "ServerErrors", "TotalLatency(ms)", "MaxLatency(ms)", "MinLatency(ms)", "AverageLatency(ms)"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data rows
	for _, burst := range individualBursts {
		row := []string{
			strconv.Itoa(burst.BurstNumber),
			strconv.Itoa(burst.SuccessfulCount),
			strconv.Itoa(burst.ClientErrors),
			strconv.Itoa(burst.ServerErrors),
			burst.TotalLatency.String(),
			burst.MaxLatency.String(),
			burst.MinLatency.String(),
			burst.AverageLatency.String(),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
