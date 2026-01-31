package engine

import (
	"github.com/imranullahkhann/golt/internal/requester"
	"github.com/imranullahkhann/golt/internal/types"
	"slices"
	"cmp"
)

func Startload(n int, url string) []types.Result {
	responseData := make([]types.Result, 0, 100)

	for i := 0; i < n; i++ {
		latency, status, err := requester.Makereq(url)

		resResult := types.Result{Latency: latency, Status: status, Err: err}
		
		responseData = append(responseData, resResult)
	}	

	slices.SortFunc(responseData, func(a, b types.Result) int {
		return cmp.Compare(a.Latency, b.Latency)
	})
	
	return responseData
}
