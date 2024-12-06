package main

import (
	"fmt"
	"log"
	// "time"
)

func main() {
	sum := 0
	fileName := "stringNumbers.txt"
	unsorted, err := ReadIntegersFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print("↓unsorted array↓")
	// fmt.Println(unsorted)
	// startTime := time.Now()
	// exec(unsorted, "start", sum)
	exec(unsorted, "middle", sum)
	// exec(unsorted, "end", sum)
	// elapsed := time.Since(startTime)
	// println("Execution time: ", elapsed.Microseconds(), "ms")
}

// select first element as a pivot
// if it selects min or max num in the array as a pivot, it takes O(n^2)
// if it selects middle, it takes O(n log n).

func exec(arr []int, rule string, sum int) int {
	res := quickSort(arr, rule, sum)
	return res
}

func quickSort(arr []int, rule string, sum int) int {
	N := len(arr)
	// base case
	if (N <= 1) {
		return sum
	}

	// 1 partition from pivot
	index := partition(arr, rule)
	// 2 quicksort left side elements
	quickSort(arr[0:index], rule, sum)
	// add index to sum
	sum += index
	// 3 quick sort right side elements
	quickSort(arr[index+1:], rule, sum)
	// add index to sum
	// sum += len(arr) - 1  -index
	// fmt.Print("↓result sum↓")
	fmt.Println(sum)
	// fmt.Print("↓sorted array↓")
	// fmt.Println(arr)
	return sum
}

func swap(arr []int, target, pivotPlace int) []int {
	tmp := arr[target]
	arr[target] = arr[pivotPlace]
	arr[pivotPlace] = tmp
	return arr
}

// partition
func partition(arr []int, rule string) int {
	N := len(arr)
	pivot_index := choose_pivot(arr, rule)
	pivot := arr[pivot_index]
	swap(arr, 0, pivot_index)
	// define i to store a pivot index
	i := 1
	for j := 1; j < N; j++ {
		// if pivot is bigger than target, swap
		if arr[j] < pivot {
			swap(arr, j, i)
			i += 1
		}
	}
	// return pivot to its original place.
	swap(arr, 0, i-1)
	 i -= 1
return i 
}

// TODO: define Pivot constants → First, Middle, Last
// decide which elements to be a pivot
// partice 2 parts from a pivot deviding by the following rule.
// 1. the left side elements of the pivot are always lower than the pivot.
// 2. the right side elements of the pivot are always bigger than the pivot
func choose_pivot(arr []int, rule string) int {
	start := 0
	end := len(arr) - 1
	middle := len(arr) / 2
	switch rule {
	case "start":
		return start
	case "end":
		return end
	case "middle":
		return find_median(start, middle, end)
	}
	return start
}

// get a medium element of the three(start, middle, end)
func find_median(start, middle, end int) int {
	if (middle < end && end < start) || (start < end && end < middle) {
		return end
	}
	// if a start number is a median
	if (start < middle && middle < end) || (end < middle && middle < start) {
		return middle
	}
	// if a last number is a median
	return start
}
