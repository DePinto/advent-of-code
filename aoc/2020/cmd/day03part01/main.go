package main

import (
	"fmt"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

const tobogganX = 3
const tobogganY = 1

type tree bool
type position struct {
	x int
	y int
}

//type vector position
type grid map[position]tree

func main() {
	// read input
	ss, err := utils.LoadInputFile("day03.txt")
	if err != nil {
		panic(err)
	}
	width := len(ss[0])

	// create grid
	g, err := newGridFromSliceString(ss)
	if err != nil {
		panic(err)
	}

	// navigate path
	crashes, err := g.navigate(width, tobogganX, tobogganY)
	if err != nil {
		panic(err)
	}
	fmt.Println(crashes)
}

func newGridFromSliceString(ss []string) (grid, error) {
	g := make(grid)

	for y, s := range ss {
		for x, r := range s {
			g[position{x: x, y: y}] = isTree(r)
		}
	}

	return g, nil
}

func isTree(r rune) tree {
	return r == rune('#')
}

func (self grid) navigate(width int, x int, y int) (int, error) {
	trees := 0
	for i, j := x, y; self.isDone(i, j); i, j = (i+x)%width, j+y {
		if self[position{x: i, y: j}] {
			trees++
		}
	}

	return trees, nil
}

func (self grid) isDone(x int, y int) bool {
	_, ok := self[position{x: x, y: y}]
	return ok
}
