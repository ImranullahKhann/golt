package stats

import (
	"math"
	"github.com/imranullahkhann/golt/internal/types"
	"cmp"
	"slices"
)

func Getpercentile(responseData []types.Result, percentile float64) int32 {
	if len(responseData) == 0 {
		panic("Empty dataset")
	}
	if percentile <= 0 || percentile > 100 {
		panic("Percentile must be in (0, 100]")
	}
	slices.SortFunc(responseData, func(a, b types.Result) int {
		return cmp.Compare(a.Latency, b.Latency)
	})

	// Nearest Rank Method
	n := len(responseData)
	index := int(math.Ceil(percentile / 100 * float64(n))) - 1

	return responseData[index].Latency
}
