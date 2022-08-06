// author(s) [jo3zeph](https://github.com/jo3zeph)
// description: Find the median from a set of values
// see median_test.go

package math

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/sort"
)

func Median[T constraints.Number](values []T) float64 {
	// Sorting the values in order

	sort.Bubble(values)

	var median float64

	// Find the length of values
	l := len(values)

	switch {
	case l == 0:
		return 0

	case l%2 == 0:
		median = float64((values[l/2-1] + values[l/2]) / 2)

	default:
		median = float64(values[l/2])
	}

	return median
}
