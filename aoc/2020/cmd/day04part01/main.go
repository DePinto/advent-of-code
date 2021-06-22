package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

type passport struct {
	BirthYear      string `json:"byr"`
	IssueYear      string `json:"iyr"`
	ExpirationYear string `json:"eyr"`
	HeightCM       string `json:"hgt"`
	HairColor      string `json:"hcl"`
	EyeColor       string `json:"ecl"` //enumeration?
	PassportID     string `json:"pid"`
	CountryID      string `json:"cid"`
}

func main() {
	// read input
	ss, err := utils.LoadInputFile("day04.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(ss) = %v\n", len(ss)) //debug

	inputs := newInputSliceFromStringSlice(ss)
	fmt.Printf("len(inputs) = %v\n", len(inputs)) //debug

	ps := newPassportSliceFromInputSlice(inputs)

	fmt.Println(countValidPassports(ps))
}

func newInputSliceFromStringSlice(ss []string) [][]string {
	inputs := [][]string{}
	input := []string{}

	for _, s := range ss {
		if len(s) == 0 {
			inputs = append(inputs, input) // what about the end
			input = []string{}
		} else {
			input = append(input, s)
		}
	}

	return inputs
}

func newPassportSliceFromInputSlice(inputs [][]string) []passport {
	ps := []passport{}

	for _, ss := range inputs {
		p := newPassportFromInput(ss)
		ps = append(ps, p)
	}

	return ps
}

func newPassportFromInput(input []string) passport {
	m := map[string]string{}

	for _, s := range input {
		words := strings.Split(s, " ")
		for _, w := range words {
			kv := strings.Split(w, ":")
			m[kv[0]] = kv[1]
			//fmt.Printf("k: %v, v: %v\n", kv[0], kv[1])
		}
	}

	p := passport{
		BirthYear:      m["byr"],
		IssueYear:      m["iyr"],
		ExpirationYear: m["eyr"],
		HeightCM:       m["hgt"],
		HairColor:      m["hcl"],
		EyeColor:       m["ecl"],
		PassportID:     m["pid"],
		CountryID:      m["cid"],
	}

	return p
}

func countValidPassports(ps []passport) int {
	validCount := 0
	for _, p := range ps {
		if isValidPassport(p) {
			validCount++
		}
	}
	return validCount
}

func isValidPassport(p passport) bool {
	if len(p.BirthYear) == 0 || len(p.ExpirationYear) == 0 || len(p.EyeColor) == 0 || len(p.HairColor) == 0 || len(p.HeightCM) == 0 || len(p.IssueYear) == 0 || len(p.PassportID) == 0 {
		debugg, err := json.MarshalIndent(p, "", " ")
		if err != nil {
			fmt.Printf("error unmarshalling: %v", err)
		}

		fmt.Printf("Failed passport:\n%s\n----------\n", string(debugg))
		return false
	}

	return true
}
