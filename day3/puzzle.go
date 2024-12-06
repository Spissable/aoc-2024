package day3

import (
	"regexp"
	"strconv"
)

type Multiplication struct {
	a int
	b int
}

type Calc struct {
	muls []Multiplication
}

func NewCalc(input string) (result Calc) {
	regex := `mul[(](\d+),\s*(\d+)[)]`
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		mul := Multiplication{
			a: a,
			b: b,
		}

		result.muls = append(result.muls, mul)
	}

	return
}

func SolvePuzzle1(input string) int {
	c := NewCalc(input)

	return c.Sum()
}

func (c Calc) Sum() (result int) {
	for _, mul := range c.muls {
		result += mul.a * mul.b
	}

	return
}
