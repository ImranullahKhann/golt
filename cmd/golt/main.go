package main

import (
	"flag"
	"fmt"
	"github.com/imranullahkhann/golt/internal/engine"
	"github.com/imranullahkhann/golt/internal/stats"
	"github.com/imranullahkhann/golt/internal/types"
)

func main() {
	url := flag.String("url", "-", "the URL of the endpoint to send requests to")
	conc := flag.Int("c", 1, "the amount of concurrent threads that will make a GET request")
	percentile := flag.Float64("p", 1, "the latency percentile to return")

	flag.Parse()

	if *url == "-" {
		panic("No URL specified!")
	}

	var responseData []types.Result = engine.Startload(*conc, *url)

	fmt.Printf("Responses: %d \n", len(responseData))
	fmt.Printf("P%v Latency: %v \n", *percentile, stats.Getpercentile(responseData, *percentile))
}
