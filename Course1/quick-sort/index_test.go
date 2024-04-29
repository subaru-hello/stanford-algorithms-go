package main

import (
	"log"
	"testing"
)

// TestQuickSort tests the merge sort function and compare with the already sorted result.
func TestQuickSort(t *testing.T) {
	hugeCase := "stringNumbers.txt"
	case1 := "Case1.txt"
	case2 := "Case2.txt"
	case3 := "Case3.txt"

	testCases := []struct {
		fileName string
		rule     string
		expect   int
	}{
		{fileName: case1, rule: "start", expect: 6},
		{fileName: case1, rule: "middle", expect: 10},
		{fileName: case1, rule: "end", expect: 6},
		{fileName: case2, rule: "start", expect: 7},
		{fileName: case2, rule: "middle", expect: 8},
		{fileName: case2, rule: "end", expect: 6},
		{fileName: case3, rule: "start", expect: 71},
		{fileName: case3, rule: "middle", expect: 73},
		{fileName: case3, rule: "end", expect: 56},
		{fileName: hugeCase, rule: "start", expect: 100},
		{fileName: hugeCase, rule: "middle", expect: 499},
		{fileName: hugeCase, rule: "end", expect: 888},
	}

	for _, tc := range testCases {
		targetArray, err := ReadIntegersFromFile(tc.fileName)
		if err != nil {
			log.Fatal(err)
		}
		sum := 0
		result := quickSort(targetArray, tc.rule, sum)
		directSorting := tc.expect

		if result != directSorting {
			t.Errorf("%v want %v", result, directSorting)
		}
	}
}
