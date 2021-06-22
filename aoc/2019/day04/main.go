package main

import "fmt"

const (
	inputMin       = 245182
	inputMax       = 790572
	passwordLength = 6

	testPass           = 122345
	testFailDecreasing = 223450
	testFailDouble     = 123789
)

type password struct {
	min     int
	max     int
	current int
	xs      [passwordLength]int
	count   int
}

func main() {
	p := password{
		min:     inputMin,
		max:     inputMax,
		current: inputMin,
		xs:      [passwordLength]int{},
		count:   0,
	}
	fmt.Println(p.checkAll())

	// p.position = testPass
	// fmt.Printf("pass: %v\n", p.checkCurrentPos())

	// p.position = testFailDecreasing
	// fmt.Printf("fail: %v\n", p.checkCurrentPos())

	// p.position = testFailDouble
	// fmt.Printf("fail: %v\n", p.checkCurrentPos())
}

func (self *password) checkAll() int {
	for self.current <= self.max {
		if self.checkCurrentPos() {
			self.count++
		}
		self.increment()
	}

	return self.count
}

func (self *password) checkCurrentPos() bool {
	self.toSlice()

	// if !self.rule3() || !self.rule4() {
	// 	return false
	// }

	// return true
	return self.rule3() && self.rule4()
}

func (self *password) toSlice() {
	xs := [passwordLength]int{}
	myInt := self.current

	for i := passwordLength - 1; i >= 0; i-- {
		r := myInt % 10
		myInt = myInt / 10
		xs[i] = r
	}

	self.xs = xs
}

// Two adjacent digits are the same (like 22 in 122345).
func (self password) rule3() bool {
	for i, j := 0, 1; j < passwordLength; i, j = i+1, j+1 {
		if self.xs[i] == self.xs[j] {
			return true
		}
	}

	return false
}

// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
func (self password) rule4() bool {
	for i, j := 0, 1; j < passwordLength; i, j = i+1, j+1 {
		if self.xs[i] > self.xs[j] {
			return false
		}
	}

	return true
}

func (self *password) increment() {
	self.current++
}

func (self *password) toInt() {
	newInt := 0

	for i := 0; i < passwordLength; i++ {
		newInt = newInt * 10
		newInt = newInt + self.xs[i]
	}

	self.current = newInt
}
