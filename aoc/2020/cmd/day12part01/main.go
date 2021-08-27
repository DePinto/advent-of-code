package main

import (
	"fmt"
	"strconv"
	"math"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

// type direction int

const (
	north = 0
	east = north + 90
	south = east + 90
	west = south + 90
)

type ship struct {
	x int
	y int
	direction int
}

type instruction struct {
	action string
	value int
}

func main() {
	ss, err := utils.LoadInputFile("day12.txt")
	if err != nil {
		panic(err)
	}

	is := newInstructionSliceFromStringSlice(ss)

	// Part 1
	newShipFromStartingPosition(0,0).navigate(is).manhattan()

	fmt.Printf("\n=====\n\n")

	// Part 2
	newShipFromStartingPosition(0,0).navigateWayPoint(is).manhattan()
}

func newInstructionSliceFromStringSlice(ss []string) []instruction {
	is := []instruction{}

	for _, s := range ss {
		is = append(is,newInstructionFromString(s))
	}

	return is
}

func newInstructionFromString(s string) instruction {
	action := s[:1]
	value, _ := strconv.Atoi(s[1:]) //share when Scotty isn't looking

	return instruction{
		action: action,
		value: value,
	}
}

func newShipFromStartingPosition(x, y int) *ship {
	ship := ship{
		x: x,
		y: y,
		direction: east,  // passing in a third parameter is hard
	}

	return &ship
}

func (s *ship) navigate(is []instruction) *ship {
	for _, i := range is {
		switch i.action {
		case "N":
			s.y += i.value
		case "S":
			s.y -= i.value
		case "E":
			s.x += i.value
		case "W":
			s.x -= i.value
		case "L":
			s.direction -= i.value
			if s.direction < 0 {
				s.direction += 360
			}
		case "R":
			s.direction = (s.direction + i.value) % 360
		case "F":
			switch s.direction {
			case north:
				s.y += i.value
			case south:
				s.y -= i.value
			case east:
				s.x += i.value
			case west:
				s.x -= i.value
			}
		}
	}

	return s
}

func (s *ship) manhattan() *ship {
	md := math.Abs(float64(s.x)) + math.Abs(float64(s.y))
	fmt.Println(md)

	return s
}

// part 2
/*
10 units east and 4 units north 10, 4
R90 
4 units east and 10 units south 4, -10
R90
10 units west and 4 units south -10, -4
R90
4 units west and 10 units north -4, 10
R90
10 units east and 4 units north 10, 4
*/

type waypoint struct {
	x int
	y int
}

func (s *ship) navigateWayPoint(is []instruction) *ship {
	wp := waypoint{
		x: 10,  // hardcoded like a champ
		y: 1,
	}
	
	for _, i := range is {
		switch i.action {
		case "N":
			wp.y += i.value
		case "S":
			wp.y -= i.value
		case "E":
			wp.x += i.value
		case "W":
			wp.x -= i.value
		case "L":
			switch i.value {
			case 90:
				tmp := wp.x
				wp.x = wp.y * -1
				wp.y = tmp
			case 180:
				wp.x = wp.x * -1
				wp.y = wp.y * -1
			case 270:
				tmp := wp.x
				wp.x = wp.y
				wp.y = tmp * -1
			}
		case "R":
			switch i.value {
			case 90:
				tmp := wp.x
				wp.x = wp.y
				wp.y = tmp * -1
			case 180:
				wp.x = wp.x * -1
				wp.y = wp.y * -1
			case 270:
				tmp := wp.x
				wp.x = wp.y * -1
				wp.y = tmp
			}
		case "F":
			s.x += i.value * wp.x
			s.y += i.value * wp.y
		}
	}

	return s
}