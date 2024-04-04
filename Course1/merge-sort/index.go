package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mergeSort(items []int)[]int{
	// base case
if(len(items) < 2){
	return items
}
first := mergeSort(items[:len(items)/2])
second := mergeSort(items[len(items)/2:])
return merge(first, second)
}

// after sorted and devided into one piece, merge them all into one array.
func merge(x,y []int) []int {
	finalOutput := []int{}
	i := 0
	j := 0

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			finalOutput = append(finalOutput, x[i])
			i++
		} else {
			finalOutput = append(finalOutput, y[j])
			j++
		}
	}

	// Append remaining elements of x
	for ; i < len(x); i++ {
		finalOutput = append(finalOutput, x[i])
	}
	// Append remaining elements of y (fixed: changed i to j)
	for ; j < len(y); j++ {
		finalOutput = append(finalOutput, y[j])
	}
	return finalOutput
}

func readIntegersFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var integers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		integers = append(integers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return integers, nil
}

func main(){
	fileName := "integerArray.txt" // Replace with your file's path
	integers, err := readIntegersFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

    // unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
    sorted := mergeSort(integers)
	fmt.Println(sorted)
}