package dynamic_test

import (
	"fmt"
	"testing"

	"github.com/TheAlgorithms/Go/dynamic"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	td := []struct {
		elements    []int
		expectedLen int
	}{
		{[]int{1, 2, 3, 4, 5, 10}, 6},
		{[]int{1, 7, 3, 4, 5}, 4}, // 1,3,4,5
		{[]int{1, 3, 5}, 3},
		{[]int{7, 1, 6}, 2},
		{[]int{4, 1, 6, 2}, 2},
		{[]int{11, 9, 6}, 1},
	}
	for _, tc := range td {
		t.Run(fmt.Sprint("test with", tc.elements), func(t *testing.T) {
			calculatedLen := dynamic.LongestIncreasingSubsequence(tc.elements)
			calculatedLenGreedy := dynamic.LongestIncreasingSubsequenceGreedy(tc.elements)
			if tc.expectedLen != calculatedLen {
				t.Fatalf("expecting a sequence of len %d to be found but the actual len was %d; input: %v", tc.expectedLen, calculatedLen, tc.elements)
			}
			if tc.expectedLen != calculatedLenGreedy {
				t.Fatalf("greedy approach failed, expecting a sequence of len %d to be found but the actual len was %d; input: %v", tc.expectedLen, calculatedLenGreedy, tc.elements)
			}
		})
	}
}
