package main

import (
	"strconv"
	"strings"
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

type rule struct {
	Bag     string
	Contain map[string]int
}

func newRuleSliceFromStringSlice(ss []string) []rule {
	rs := []rule{}

	for _, s := range ss {
		r := newRuleFromString(s)
		rs = append(rs, r)
	}

	return rs
}

func newRuleFromString(s string) rule {
	s1 := strings.Split(s, " ")
	bag := s1[0] + s1[1]
	m := newContainFromStringSlice(s1[4:])
	r := rule{
		Bag:     bag,
		Contain: m,
	}

	return r
}

func newContainFromStringSlice(ss []string) map[string]int {
	m := map[string]int{}

	if len(ss) == 3 {
		return m
	}

	for i := 0; i < len(ss); i = i + 4 {
		str, i := newContainEntryFromStringSlice(ss[i : i+3])
		m[str] = m[str] + i
	}

	return m
}

func newContainEntryFromStringSlice(ss []string) (string, int) {
	i, _ := strconv.Atoi(ss[0])
	return ss[1] + ss[2], i
}
