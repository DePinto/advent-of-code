package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mlhoyt/aoc2020/go/pkg/utils"
)

type passwordEntry struct {
	min      int
	max      int
	r        rune
	password string
}

func main() {
	ss, err := utils.LoadInputFile("day02.txt")
	if err != nil {
		panic(err)
	}

	pes, err := newPasswordEntrySliceFromStringSlice(ss)
	if err != nil {
		panic(err)
	}

	totalValidPasswords := calculateTotalValidPasswords(pes)

	fmt.Printf("Total valid password(s): %v\n", totalValidPasswords)
}

func newPasswordEntrySliceFromStringSlice(input []string) ([]passwordEntry, error) {
	pes := []passwordEntry{}
	for _, s := range input {
		pe, err := newPasswordEntryFromString(s)
		if err != nil {
			return nil, err
		}
		pes = append(pes, *pe)
	}
	return pes, nil
}

func newPasswordEntryFromString(input string) (*passwordEntry, error) {
	//  1-3 a: abcde
	ss := strings.Split(input, " ")
	errs := []string{}

	min, err := getMinFromString(ss[0])
	if err != nil {
		errs = append(errs, err.Error())
	}
	max, err := getMaxFromString(ss[0])
	if err != nil {
		errs = append(errs, err.Error())
	}
	r, err := getRuneFromString(ss[1])
	if err != nil {
		errs = append(errs, err.Error())
	}
	password, err := getPasswordFromString(ss[2])
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) != 0 {
		return nil, errors.New(strings.Join(errs, ", "))
	}

	pe := passwordEntry{
		min:      *min,
		max:      *max,
		r:        *r,
		password: *password,
	}
	return &pe, nil
}

func getMinFromString(input string) (*int, error) {
	return getIntFromStringRange(input, 0)
}

func getMaxFromString(input string) (*int, error) {
	return getIntFromStringRange(input, 1)
}

func getIntFromStringRange(input string, pos int) (*int, error) {
	ss := strings.Split(input, "-")

	i, err := strconv.Atoi(ss[pos])
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func getRuneFromString(input string) (*rune, error) {
	for _, r := range input {
		return &r, nil
	}
	return nil, errors.New("BUT HOWWWW?!?")
}

func getPasswordFromString(input string) (*string, error) {
	return &input, nil
}

func calculateTotalValidPasswords(pes []passwordEntry) int {
	totalValidPasswords := 0
	for _, pe := range pes {
		if pe.validate() {
			totalValidPasswords += 1
		}
	}
	return totalValidPasswords
}

func (self passwordEntry) validate() bool {
	matches := 0
	for i, v := range self.password {

		if i+1 == self.min || i+1 == self.max {
			if v == self.r {
				matches += 1
			}
		}
	}

	if matches == 1 {
		return true
	}
	return false
}
