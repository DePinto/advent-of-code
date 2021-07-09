package main

import (
	"fmt"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

func main() {
	fmt.Println("vim-go")
	inputString, err := utils.LoadInputFileV2("day06_test.txt")
	if err != nil {
		panic(err)
	}

	total := eachGroup(countUniqueRunes, inputString)
	homo := eachGroup(countHomoRunes, inputString)

	fmt.Println(total)
	fmt.Println(homo)
}

func eachGroup(f func(string) int, xs []string) int {
	total := 0

	for _, s := range xs {
		i := f(s)
		total = total + i
	}

	return total
}

func countUniqueRunes(s string) int {
	m := map[rune]bool{}

	for _, r := range s {
		m[r] = true
	}

	return len(m)
}

func countHomoRunes(s string) int {
	m := map[rune]int{}

	for _, r := range s {
		m[r] = m[r] + 1
	}

	count := 0
	for _, i := range m {
		if i == len(s) {
			count++
		}
	}

	return count
}
