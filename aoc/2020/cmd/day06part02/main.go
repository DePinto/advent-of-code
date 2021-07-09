package main

import (
	"fmt"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

type person string
type group []person
type manifest []group

func main() {
	ss, err := utils.LoadInputFile("day06.txt")
	if err != nil {
		panic(err)
	}

	m := newManifestFromStringSlice(ss)
	//fmt.Println(m)

	fmt.Println(countManifest(m))
}

func newManifestFromStringSlice(ss []string) manifest {
	m := manifest{}
	g := group{}

	for _, s := range ss {
		if len(s) == 0 {
			m = append(m, g)
			g = group{}
		} else {
			g = append(g, person(s)) // unexpected
		}

	}

	return m
}

func countManifest(m manifest) int {
	count := 0

	for _, g := range m {
		i := countGroup(g)
		count = count + i
	}

	return count
}

func countGroup(g group) int {
	m := map[rune]int{}

	for _, p := range g {
		for _, r := range p {
			m[r] = m[r] + 1
		}
	}

	count := 0
	for _, v := range m {
		if v == len(g) {
			count++
		}
	}

	return count
}
