package main

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntegersFromFile(fileName string) ([]int, error) {
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
