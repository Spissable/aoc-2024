package day7

import (
	"aoc-2024/util"
	"fmt"
	"slices"
	"strings"
)

type Equation struct {
	TestValue int
	Values    []int
}

func NewEquations(input string) (result []Equation) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		colonSplit := strings.Split(line, ":")
		if len(colonSplit) < 2 {
			break
		}

		testValue := util.StringToNum(colonSplit[0])
		values := util.StringsToNums(strings.Split(strings.TrimSpace(colonSplit[1]), " "))
		result = append(result, Equation{TestValue: testValue, Values: values})
	}

	return
}

func SolvePuzzle1(input string) (result int) {
	eqs := NewEquations(input)

	for _, eq := range eqs {
		combs := make([]int, 0)
		eq.CalculateCombinations(1, eq.Values[0], &combs)
		fmt.Println(combs)
		if slices.Contains(combs, eq.TestValue) {
			result += eq.TestValue
		}
	}

	return
}

func (eq Equation) CalculateCombinations(index int, current int, results *[]int) {
	if index == len(eq.Values) {
		*results = append(*results, current)
		return
	}

	// Add the current number
	eq.CalculateCombinations(index+1, current+eq.Values[index], results)

	// Multiply the current number
	eq.CalculateCombinations(index+1, current*eq.Values[index], results)
}
