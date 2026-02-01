package stats

import (
	"math"
	"github.com/imranullahkhann/golt/internal/types"
	"cmp"
	"slices"
)

func Aggregatedata(result <-chan types.Result, collection chan<- []types.Result) {
	responseData := make([]types.Result, 0, 100)

	for res := range result {
		responseData = append(responseData, res)
	}

	slices.SortFunc(responseData, func(a, b types.Result) int {
		return cmp.Compare(a.Latency, b.Latency)
	})

	collection <- responseData
	close(collection)
}

func Getpercentile(sorted []types.Result, percentile float64) int32 {
	if len(sorted) == 0 {
		panic("Empty dataset")
	}
	if percentile <= 0 || percentile > 100 {
		panic("Percentile must be in (0, 100]")
	}

	// Nearest Rank Method
	n := len(sorted)
	index := int(math.Ceil(percentile / 100 * float64(n))) - 1

	return sorted[index].Latency
}
