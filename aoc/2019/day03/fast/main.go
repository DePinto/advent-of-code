package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// part 1: 248
// part 2: 28580

const (
	day = 3
	R   = 82
	D   = 68
	L   = 76
	U   = 85
)

type wires []wire

type wire []movement

type movement struct {
	direction byte
	length    int
}

type positions []position

type position struct {
	x    int
	y    int
	dist int
}

type intersections []intersection

type intersection struct {
	x      int
	y      int
	w1Dist int
	w2Dist int
}

type measure func(intersection) int

type key struct {
	x int
	y int
}

type value struct {
	ocr     int
	sumDist int
}

type interMap map[key]value

func main() {
	w1 := wire{
		{
			direction: R,
			length:    8,
		},
		{
			direction: U,
			length:    5,
		},
		{
			direction: L,
			length:    5,
		},
		{
			direction: D,
			length:    3,
		},
	}

	w2 := wire{
		{
			direction: U,
			length:    7,
		},
		{
			direction: R,
			length:    6,
		},
		{
			direction: D,
			length:    4,
		},
		{
			direction: L,
			length:    4,
		},
	}

	testPart1 := solveManhattan(wires{w1, w2})
	testPart2 := solveShortestPath(wires{w1, w2})
	fmt.Printf("--Test Cases--\nPart 1 Manhattan:\t%v\nPart 2 Shortest Path:\t%v\n\n", testPart1, testPart2)

	inputWires := getProblemInput(day)
	part1 := solveManhattan(inputWires)
	part2 := solveShortestPath(inputWires)
	fmt.Printf("--AoC Cases--\nPart 1 Manhattan:\t%v\nPart 2 Shortest Path:\t%v\n\n", part1, part2)

}

func solveManhattan(ws wires) int { return getClosest(prepProblem(ws), getManhattan) }

func getManhattan(inter intersection) int {
	return int(math.Abs(float64(inter.x)) + math.Abs(float64(inter.y)))
}

func solveShortestPath(ws wires) int { return getClosest(prepProblem(ws), getSteps) }

func getSteps(inter intersection) int { return inter.w1Dist + inter.w2Dist }

func prepProblem(ws wires) intersections {
	w1 := ws[0]
	w2 := ws[1]

	pos1 := getPositions(w1)
	pos2 := getPositions(w2)

	return getIntersections(pos1, pos2)
}

func getPositions(w wire) positions {

	pos := positions{}
	p := position{0, 0, 0}

	for _, mov := range w {
		switch mov.direction {
		case R:
			for i := 0; i < mov.length; i++ {
				p.x++
				p.dist++
				pos = append(pos, p)
			}
		case D:
			for i := 0; i < mov.length; i++ {
				p.y--
				p.dist++
				pos = append(pos, p)
			}
		case L:
			for i := 0; i < mov.length; i++ {
				p.x--
				p.dist++
				pos = append(pos, p)
			}
		case U:
			for i := 0; i < mov.length; i++ {
				p.y++
				p.dist++
				pos = append(pos, p)
			}
		}
	}

	return pos
}

func getIntersections(pos1, pos2 positions) intersections {
	inters := intersections{}

	points := createMap(pos1, pos2)

	for i, v := range points {
		if v.ocr > 1 {
			inters = append(inters, intersection{
				x:      i.x,
				y:      i.y,
				w1Dist: v.sumDist,
				w2Dist: 0,
			})
		}
	}

	return inters
}

func createMap(pos1, pos2 positions) interMap {
	points := make(interMap)

	points = mapLoop(pos1, points)
	points = mapLoop(pos2, points)

	return points
}

func mapLoop(pos positions, m map[key]value) map[key]value {
	for _, p := range pos {
		k := key{
			p.x,
			p.y,
		}
		v := value{
			ocr:     1 + m[k].ocr,
			sumDist: p.dist + m[k].sumDist,
		}
		// if m[k].ocr < 2 {
		m[k] = v
		// }
	}

	return m
}

func getClosest(inters intersections, fn measure) int {
	dist := fn(inters[0])

	for _, inter := range inters {
		newDist := fn(inter)
		if newDist < dist {
			dist = newDist
		}
	}

	return dist
}

func getProblemInput(day int) wires {
	const session = ""
	url := fmt.Sprintf("https://adventofcode.com/2019/day/%v/input", day)
	cookie := http.Cookie{
		Name:   "session",
		Value:  session,
		Domain: ".adventofcode.com",
		Path:   "/",
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error creating http req - %v\n", err)
		return nil
	}
	req.AddCookie(&cookie)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error on http req - %v\n", err)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error on http body - %v\n", err)
		return nil
	}

	wires := csvToWires(body)

	return wires
}

func csvToWires(input []byte) wires {
	w := wires{}
	xm1 := []movement{}
	xm2 := []movement{}
	s := string(input)
	s = strings.TrimSuffix(s, "\n")
	s1 := strings.Split(s, "\n")[0]
	s2 := strings.Split(s, "\n")[1]
	xs1 := strings.Split(s1, ",")
	xs2 := strings.Split(s2, ",")

	for _, v := range xs1 {
		d := v[0]
		ms := v[1:]
		mi, err := strconv.Atoi(ms)
		if err != nil {
			fmt.Printf("error on strcon - %v\n", err)
			continue
		}
		m := movement{
			direction: d,
			length:    mi,
		}
		xm1 = append(xm1, m)
	}

	for _, v := range xs2 {
		d := v[0]
		ms := v[1:]
		mi, err := strconv.Atoi(ms)
		if err != nil {
			fmt.Printf("error on strcon - %v\n", err)
			continue
		}
		m := movement{
			direction: d,
			length:    mi,
		}
		xm2 = append(xm2, m)
	}

	w1 := xm1

	w2 := xm2

	w = append(w, w1)
	w = append(w, w2)
	return w
}
