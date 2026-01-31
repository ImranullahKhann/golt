package stats

import (
	"math"
	"github.com/imranullahkhann/golt/internal/types"
)

func Getpercentile(sorted []types.Result, percentile float64) int32 {
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
