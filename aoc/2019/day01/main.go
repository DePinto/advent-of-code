package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	const session = ""
	input := getProblemInput(session, 1)

	// Part 1
	reqFuel := moduleFuelCounter(input)
	fmt.Printf("Part 1: %v\n", reqFuel)

	// Part 2
	trueFuel := trueModuleFuelCounter(input)
	fmt.Printf("Part 2: %v\n", trueFuel)
}

func getProblemInput(session string, day int) []int {
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

	intBody := xsByteToInt(body)

	return intBody
}

func xsByteToInt(input []byte) []int {
	result := []int{}
	str := string(input)
	xsStr := strings.Split(str, "\n")

	for _, v := range xsStr {
		newInt, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("error on strcon - %v\n", err)
			continue
		}
		result = append(result, newInt)
	}

	return result
}

func fuelForMass(mass int) int {
	const fuelRatio = 3
	const fuelOffset = 2

	return (mass / fuelRatio) - fuelOffset
}

func moduleFuelCounter(modules []int) int {
	fuelTotal := 0

	for _, v := range modules {
		fuelTotal += fuelForMass(v)
	}

	return fuelTotal
}

func trueModuleFuelCounter(input []int) int {
	fuelTotal := 0

	for _, v := range input {
		for remainingMass := v; remainingMass > 0; {
			newFuel := fuelForMass(remainingMass)
			if newFuel > 0 {
				fuelTotal += newFuel
			}
			remainingMass = newFuel
		}
	}

	return fuelTotal
}
