package main

import (
	"fmt"
	"test_gen/benchmark"
	"test_gen/generator"
)

func main() {
	cnt := 1000
	profileSet := generator.Generator(cnt)

	// fmt.Println("POST request performance...")
	// fmt.Println()
	// benchmark.PostRequestBenchMark(profileSet)
	// fmt.Println("\n")

	fmt.Println("GET request performance...")
	fmt.Println()
	benchmark.GetRequestBenchMark(len(profileSet))

	// fmt.Println("Redis Post request performance...")
	// fmt.Println()
	// benchmark.RedisPostRequestBenchMark(profileSet)
	// fmt.Println("\n")

	fmt.Println("Redis GET request performance...")
	fmt.Println()
	benchmark.RedisGetRequestBenchMark(len(profileSet))

}
