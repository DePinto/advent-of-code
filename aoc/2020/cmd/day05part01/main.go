package main

import (
	"fmt"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

const (
	maxRows          = 127
	maxColumns       = 7
	rowCharacters    = 7
	columnCharacters = 3
)

func main() {
	ss, err := utils.LoadInputFile("day05.txt")
	if err != nil {
		panic(err)
	}
	currentMax := 0

	m := map[int]bool{} // part 2

	for _, s := range ss {
		id := findSeatID(s)
		m[id] = true // part 2
		if id > currentMax {
			currentMax = id
		}
	}

	fmt.Printf("max:\t%v\n", currentMax)

	// part 2
	for i, _ := range m {
		_, me := m[i+1]
		_, fatty := m[i+2]
		if me == false && fatty == true {
			fmt.Printf("I'm at seat %v\n", (i + 1))
		}
	}
}

func findSeatID(s string) int {
	row := findRow(s)
	column := findColumn(s)

	return calculateSeatID(row, column)
}

func findRow(s string) int {
	l := 0
	r := maxRows
	rs := []rune(s)

	for i := 0; i < rowCharacters-1; i++ {
		if rs[i] == 'F' {
			r = (r + l) / 2
		} else {
			l = (r + l + 1) / 2
		}
	}

	if rs[rowCharacters-1] == 'F' {
		return l
	}

	return r
}

func findColumn(s string) int {
	l := 0
	r := maxColumns
	rs := []rune(s)
	rs = rs[(rowCharacters):]

	for i := 0; i < columnCharacters-1; i++ {
		if rs[i] == 'L' {
			r = (r + l) / 2
		} else {
			l = (r + l + 1) / 2
		}
	}

	if rs[columnCharacters-1] == 'L' {
		return l
	}

	return r
}

func calculateSeatID(r, c int) int {
	return (r * 8) + c
}
