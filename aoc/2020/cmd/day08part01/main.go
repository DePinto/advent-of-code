package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

const (
	pass = iota
	fail
	loop
)

type program struct {
	accumulator    int
	instructions   []instruction
	currentStep    int
	executionCount map[int]int //v2
}

type instruction struct {
	operation string
	argument  int
	//executionCount int //v1
}

func main() {
	ss, err := utils.LoadInputFile("day08.txt")
	if err != nil {
		panic(err)
	}

	p := newProgramFromStringSlice(ss)

	p.Run(0) //p1
	p.Reset()
	p.RunWithFix() //p2

}

func newProgramFromStringSlice(ss []string) *program {
	is := newInstructionSliceFromStringSlice(ss)
	ec := map[int]int{}

	return &program{
		accumulator:    0,
		instructions:   is,
		currentStep:    0,
		executionCount: ec,
	}
}

func newInstructionSliceFromStringSlice(ss []string) []instruction {
	is := []instruction{}

	for _, s := range ss {
		i := newInstructionFromString(s)
		is = append(is, i)
	}

	return is
}

func newInstructionFromString(s string) instruction {
	ss := strings.Split(s, " ")
	argument, _ := strconv.Atoi(ss[1])

	return instruction{
		operation: ss[0],
		argument:  argument,
		//executionCount: 0, //v1
	}
}

func (p *program) Run(step int) int {
	p.currentStep = p.currentStep + step

	if p.currentStep == len(p.instructions) && step == 1 {
		fmt.Printf("Winnner winner, chicken dinner.\nAccumulator: %v\n", p.accumulator)
		return pass
	}

	if p.currentStep >= len(p.instructions) {
		fmt.Printf("ERROR: STEP EXCEEDS CURRENT INSTRUCTION SET. \nAccumulator: %v\n", p.accumulator)
		return fail
	}

	//if p.instructions[p.currentStep].executionCount > 0 { //v1
	if p.executionCount[p.currentStep] > 0 {
		fmt.Printf("Loop detected.  Ending execution.\nAccumulator: %v\n", p.accumulator)
		return loop
	}

	//p.instructions[p.currentStep].executionCount++ //v1
	p.executionCount[p.currentStep]++

	switch p.instructions[p.currentStep].operation {
	case "acc":
		p.accumulator += p.instructions[p.currentStep].argument
		return p.Run(1)
	case "jmp":
		return p.Run(p.instructions[p.currentStep].argument)
	case "nop":
		return p.Run(1)
	}

	return -1
}

// Part 2
func (p *program) RunWithFix() {
	for i, v := range p.instructions {
		switch v.operation {
		case "nop":
			p.instructions[i].operation = "jmp"
			p.Reset()
			result := p.Run(0)

			if result == pass {
				return
			}

			if result != pass {
				p.instructions[i].operation = "nop"
			}
		case "jmp":
			p.instructions[i].operation = "nop"
			p.Reset()
			result := p.Run(0)

			if result == pass {
				return
			}

			if result != pass {
				p.instructions[i].operation = "jmp"
			}
		}
	}
}

func (p *program) Reset() {
	ec := map[int]int{}

	p.accumulator = 0
	p.currentStep = 0
	p.executionCount = ec
	//for i, _ := range p.instructions {
	//	p.instructions[i].executionCount = 0
	//}
}
