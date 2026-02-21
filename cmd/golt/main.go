package main

import (
	"flag"
	"fmt"
	"github.com/imranullahkhann/golt/internal/engine"
	"github.com/imranullahkhann/golt/internal/stats"
	"github.com/imranullahkhann/golt/internal/types"
)

func main() {
	u := flag.String("url", "-", "the URL of the endpoint to send requests to")
	n := flag.Int("n", 1000, "the number of requests to make per second")
	c := flag.Int("c", 100, "the number of go routines that will be concurrently sending GET requests")
	t := flag.Int("t", 10, "the number of seconds to make requests for")
	percentile := flag.Float64("p", 90, "the latency percentile to return")

	flag.Parse()
	url := *u
	num := *n
	conc := *c
	tim := *t

	if url == "-" {
		panic("No URL specified!")
	}
	if num % 10 != 0 {
		panic("Number of requests per second should be exactly divisible by 10")
	}
	if num < 100 || num > 50000 {
		panic("Number of requests per second should be in the range [100, 50000]")
	}
	if conc < 100 || conc > 2000 {
		panic("Number of go routines should be in the range [100, 2000]")
	}
	if tim < 10 || tim > 300 {
		panic("Duration should be in the range of [10, 300] seconds")
	}	

	var responseData []types.Result = engine.Startload(url, num, conc, tim)

	fmt.Printf("Responses: %d \n", len(responseData))
	fmt.Printf("P%v Latency: %v \n", *percentile, stats.Getpercentile(responseData, *percentile))
}
