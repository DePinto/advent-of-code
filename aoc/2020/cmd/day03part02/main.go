package main

import (
	"fmt"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

type tree bool
type position struct {
	x int
	y int
}

//type vector position
type grid map[position]tree
type slopes []position

func main() {
	// read input
	ss, err := utils.LoadInputFile("day03.txt")
	if err != nil {
		panic(err)
	}
	width := len(ss[0])

	slopeInputs := slopes{
		position{
			x: 1,
			y: 1,
		},
		position{
			x: 3,
			y: 1,
		},
		position{
			x: 5,
			y: 1,
		},
		position{
			x: 7,
			y: 1,
		},
		position{
			x: 1,
			y: 2,
		},
	}

	// create grid
	g, err := newGridFromSliceString(ss)
	if err != nil {
		panic(err)
	}

	// navigate path
	totalCrashes := 1
	for _, v := range slopeInputs {
		crashes, err := g.navigate(width, v.x, v.y)
		if err != nil {
			panic(err)
		}
		fmt.Println(crashes)
		totalCrashes = totalCrashes * crashes
	}

	// show yo moves
	fmt.Println(totalCrashes)
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
