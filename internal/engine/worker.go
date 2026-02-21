package engine

import (
	"github.com/imranullahkhann/golt/internal/types"  
	"github.com/imranullahkhann/golt/internal/requester"  
	"time"
	"sync"
)

func Startload(url string, n int, c int, t int) []types.Result {
	// Initialize a pool of c go routines, and make them wait for request cue via a channel
	// Send n cues per second
	// For t seconds
	var wg sync.WaitGroup
	done := make(chan bool)
	message := make(chan uint8, n)
	result := make(chan types.Result, n*t)
	results := make([]types.Result, 0, n*t)

	for i := 0; i < c; i++ {
		wg.Add(1)
		go worker(result, message, url, &wg)
	} 

	ticker := time.NewTicker(100 * time.Millisecond);
	go func() {
		loop:
		for true {
			select {
				case <-done:
					close(message)
					break loop
				case <-ticker.C:
					for i := 0; i < n/10; i++ {
						message <- 1
					}
			}
		}	
	}()
	
	time.Sleep(time.Duration(t) * time.Second + time.Duration(100) * time.Millisecond)
	done <- true
	wg.Wait()
	close(done)
	close(result)
	for res := range result {
		results = append(results, res)
	}	

	return results
}


func worker(result chan<- types.Result, message <-chan uint8, url string, wg* sync.WaitGroup) {
	for range message {
		latency, status, err := requester.MakeReq(url)
		res := types.Result{Latency: latency, Status: status, Err: err} 
		result <- res
	}
	wg.Done()
}
