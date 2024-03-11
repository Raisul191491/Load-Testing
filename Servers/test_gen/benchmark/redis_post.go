package benchmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"test_gen/generator"
	"time"
)

func RedisPostRequestBenchMark(profileSet []generator.Profile) {
	apiURL := "http://localhost:9010/profile/redis/create"
	numRequests := len(profileSet)

	startTime := time.Now()
	for i := 0; i < numRequests; i++ {
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
	}

	elapsed := time.Since(startTime)

	// Performance metrics
	fmt.Printf("Total requests sent: %d\n", numRequests)
	fmt.Printf("Total time taken: %v\n", elapsed)
	fmt.Printf("Average time per request: %v\n", elapsed/time.Duration(numRequests))
}
