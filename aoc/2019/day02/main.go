package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	day          = 2
	targetOutput = 19690720
)

func main() {
	input := getProblemInput(day)
	input[1] = 12
	input[2] = 2
	result, err := Intcode(input, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result:\n%v\n", result)
	trueResult := IntToCSV(result) // READING IS HARD
	fmt.Printf("Formatted Result:\n%v\n", trueResult)

	fmt.Println("\nPART 2:")
	mem := getProblemInput(day)
	noun, verb, err := FindNounVerb(mem, targetOutput)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Noun: %v\nVerb: %v\n", noun, verb)
	fmt.Printf("Forthem: %v\n", 100*noun+verb)
}

func Intcode(input []int, pos int) ([]int, error) {
	switch input[pos] {
	case 1:
		val1 := input[input[pos+1]]
		val2 := input[input[pos+2]]
		input[input[pos+3]] = val1 + val2
	case 2:
		val1 := input[input[pos+1]]
		val2 := input[input[pos+2]]
		input[input[pos+3]] = val1 * val2
	case 99:
		return input, nil
	default:
		err := fmt.Errorf("Inccorect OP Code\nValue: %v\nPosition:%v", input[pos], pos)
		return input, err
	}
	return Intcode(input, pos+4)
}

func FindNounVerb(input []int, target int) (int, int, error) {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			thisRun := make([]int, len(input))
			copy(thisRun, input)
			thisRun[1] = i
			thisRun[2] = j
			output, _ := Intcode(thisRun, 0)
			if output[0] == target {
				return i, j, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("no verb and noun found for targeted value")
}

func getProblemInput(day int) []int {
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

	// intBody := xsByteToInt(body)
	intBody := csvToInt(body)

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

func csvToInt(input []byte) []int {
	result := []int{}
	s := string(input)
	s = strings.TrimSuffix(s, "\n")
	xs := strings.Split(s, ",")

	for _, v := range xs {
		newInt, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("error on strcon - %v\n", err)
			continue
		}
		result = append(result, newInt)
	}

	return result
}

func IntToCSV(input []int) string {
	xs := []string{}
	for _, v := range input {
		s := strconv.Itoa(v)
		xs = append(xs, s)
	}
	str := strings.Join(xs, ",")

	return str
}
