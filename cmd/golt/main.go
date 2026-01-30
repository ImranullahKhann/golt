package main

import (
	"flag"
	"github.com/imranullahkhann/golt/internal/requester"
	"fmt"
	"slices"
	"cmp"
	"math"
)

type Result struct {
	Latency int32
	Status int16
	Err error
}

func main() {
	url := flag.String("url", "-", "the URL of the endpoint to send requests to")
	percentile := flag.Float64("p", 1, "the latency percentile to return")

	flag.Parse()

	if *url == "-" {
		panic("No URL specified!")
	}

	responseData := make([]Result, 0, 100)
	
	for i := 0; i < 10; i++ {
		latency, status, err := requester.Makereq(*url)

		resResult := Result{Latency: latency, Status: status, Err: err}
		
		responseData = append(responseData, resResult)
	}

	slices.SortFunc(responseData, func(a, b Result) int {
		return cmp.Compare(a.Latency, b.Latency)
	})

	fmt.Printf("Responses: \n %v \n", responseData)
	fmt.Printf("P%v Latency: %v \n", *percentile,  getPercentile(responseData, *percentile))
}

func getPercentile(sorted []Result, percentile float64) int32 {
	if len(sorted) == 0 {
		panic("Empty dataset")
	}
	if percentile <= 0 || percentile > 100 {
		panic("Percentile must be in (0, 100]")
	}

	n := len(sorted)
	index := int(math.Ceil(percentile / 100 * float64(n))) - 1

	return sorted[index].Latency
}
