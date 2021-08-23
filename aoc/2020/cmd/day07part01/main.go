package main

import (
	"encoding/json"
	"fmt"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

/*
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
*/

type collection map[string]map[string]int

func main() {
	ss, err := utils.LoadInputFile("day07_test.txt")
	if err != nil {
		panic(err)
	}

	rs := newRuleSliceFromStringSlice(ss)

	pretty, _ := json.MarshalIndent(rs, "", " ")
	fmt.Println(string(pretty))
}

func newCollectionFromRuleSlice(rs []rule) collection {
	c := collection{}
	return c
}

func stuff(rs []rule, bag string) map[string]int {
	return nil
}
