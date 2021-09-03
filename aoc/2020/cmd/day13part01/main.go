package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

type system struct{
	startTime int
	busSchedule []int
}

type part1 struct {
	startTime int
	departureTime int
	busRoute int
}

func main() {
	ss, err := utils.LoadInputFile("day13.txt")
	mustNot(err)

	newSystemFromStringSlice(ss).
	findPart1().
	calc()
}

func newSystemFromStringSlice(ss []string) *system {
	startTime, err := strconv.Atoi(ss[0])
	mustNot(err)

	busSchedule := newIntSliceFromString(ss[1])

	return &system{
		startTime: startTime,
		busSchedule: busSchedule,
	}
}

func newIntSliceFromString(s string) []int {
	ss := strings.Split(s, ",")
	
	is := []int{}
	for _, s := range ss {
		if s != "x" {
			i, err := strconv.Atoi(s)
			mustNot(err)

			is = append(is, i)
		}
	}

	return is
}

func (self *system) findPart1() (*part1) {
	i := self.startTime
	for  {
		for _, val := range self.busSchedule {
			if i % val == 0 {
				return &part1{
					startTime: self.startTime,
					departureTime: i,
					busRoute: val,
				}
			}
		}
		i++
	}
}


func (self *part1) calc() {
	fmt.Println((self.departureTime - self.startTime) * self.busRoute)
}

func mustNot(err error) {
	if err != nil {
		panic(err)
	}
}