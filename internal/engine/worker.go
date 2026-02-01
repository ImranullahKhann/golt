package engine

import (
	"github.com/imranullahkhann/golt/internal/requester"
	"github.com/imranullahkhann/golt/internal/types"
	"github.com/imranullahkhann/golt/internal/stats"
	"sync"
)

func Startload(n int, url string) []types.Result {
	result := make(chan types.Result, n)
	collection := make(chan []types.Result)
	var wg sync.WaitGroup

	go stats.Aggregatedata(result, collection)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(result, url, &wg)
	}

	wg.Wait()
	close(result)

	responseData := <- collection

	return responseData
}


func worker(result chan<- types.Result, url string, wg *sync.WaitGroup) {
	latency, status, err := requester.Makereq(url)
	resResult := types.Result{Latency: latency, Status: status, Err: err}

	result <- resResult
	wg.Done()
}
