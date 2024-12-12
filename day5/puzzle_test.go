package day5_test

import (
	"aoc-2024/day5"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(5)

	result := day5.SolvePuzzle1(input)

	if result != 4135 {
		t.Error(result)
	}
}
