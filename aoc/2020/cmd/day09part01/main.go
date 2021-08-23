package main

import (
	"fmt"
	"strconv"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

const preamble = 25

func main() {
	ss, err := utils.LoadInputFile("day09.txt")
	if err != nil {
		panic(err)
	}

	is := newIntSliceFromStringSlice(ss)

	fmt.Println(findWeakness(is))
}

func newIntSliceFromStringSlice(ss []string) []int {
	is := []int{}

	for _, s := range ss {
		if i, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			is = append(is, i)
		}
	}

	return is
}

func findWeakness(is []int) int {
	for i := preamble; i < len(is); i++ {
		if !findSumInIntSlice(is[i-preamble:i], is[i]) {
			return is[i]
		}
	}

	return -1
}

func findSumInIntSlice(is []int, i int) bool {
	for ind, _ := range is {
		for ind2 := ind + 1; ind2 < len(is); ind2++ {
			if is[ind]+is[ind2] == i {
				return true
			}
		}
	}

	return false
}
