package main

import (
	"log"
	"testing"
)

func createRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Helper function to compare two slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestMergeSort tests the merge sort function and compare with the already sorted result.
func TestMergeSort(t *testing.T) {
    fileName := "integerArray.txt"
    loadedValues,err := readIntegersFromFile(fileName)
    if err != nil {
		log.Fatal(err)
	}

	testCases := []struct {
		numbers []int
		expect  []int
	}{
		{numbers: []int{3, 6, 2, 1, 5, 8, 10, 0, 4, 7, 9}, expect: createRange(0, 10)},
		{numbers: []int{1, 6, 2, 10, 5, 8, 3, 4, 7, 0, 9}, expect: createRange(0, 10)},
		{numbers: []int{0, 3, 10, 2, 1, 5, 8, 6, 4, 7, 9}, expect: createRange(0, 10)},
		{numbers: loadedValues, expect: createRange(1, 100000)},
	}

	for _, tc := range testCases {
		x := tc.numbers

		result := mergeSort(x)
		directSorting := tc.expect

		if !slicesEqual(result, directSorting) {
			t.Errorf("%v want %v", result, directSorting)
		}
	}
}
