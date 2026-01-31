package main

import (
	"flag"
	"fmt"
	"github.com/imranullahkhann/golt/internal/stats"
	"github.com/imranullahkhann/golt/internal/engine"
	"github.com/imranullahkhann/golt/internal/types"
)

func main() {
	url := flag.String("url", "-", "the URL of the endpoint to send requests to")
	percentile := flag.Float64("p", 1, "the latency percentile to return")

	flag.Parse()

	if *url == "-" {
		panic("No URL specified!")
	}


	var responseData []types.Result = engine.Startload(10, *url)

	fmt.Printf("Responses: \n %v \n", responseData)
	fmt.Printf("P%v Latency: %v \n", *percentile,  stats.Getpercentile(responseData, *percentile))
}


