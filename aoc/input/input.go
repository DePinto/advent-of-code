package input

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const session = ""

func getProblemInput(day int) []int {

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
