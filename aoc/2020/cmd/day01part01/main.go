package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

const (
	targetSum = 2020
)

func main () {
	inputString, err := utils.LoadInputFile("day01.txt")
	if err != nil {
		panic(err)
	}

	inputInt, err := stringstoInts(inputString)
	if err != nil {
		panic(err)
	}

	iPtr, err := findProduct(inputInt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", *iPtr)
}

func stringstoInts(input []string) ([]int, error) {
	ints := []int{}

	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	return ints, nil
}

func findProduct(input []int) (*int, error) {
	for i, a := range input {
		for _, b := range input[i+1:] { 
			if a + b == targetSum {
				product := a * b
				return &product, nil
			}
		}
	} 
	return nil, errors.New("no solution found")
}

func debug(input interface{}) {
	fmt.Printf("%v\n",input)
}