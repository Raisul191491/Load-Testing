package benchmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"test_gen/generator"
	"time"
)

func PostRequestBenchMark(profileSet []generator.Profile) {
	apiURL := "http://localhost:9010/profile/create"
	numRequests := len(profileSet)

	total, maxLatency, minLatency := 0.0, math.Inf(-1), math.Inf(0)

	for i := 0; i < numRequests; i++ {
		startTime := time.Now()
		jsonData, err := json.Marshal(profileSet[i])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		_, err = http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			return
		}
		elapsed := time.Since(startTime)

		maxLatency = math.Max(float64(elapsed), maxLatency)
		minLatency = math.Min(minLatency, float64(elapsed))

		total += float64(elapsed)
	}

	// Performance metrics
	fmt.Printf("Total requests sent: %d\n", numRequests)
	// fmt.Printf("Total time taken: %v\n", elapsed)
	// fmt.Printf("Average time per request: %v\n", elapsed/time.Duration(numRequests))
}
