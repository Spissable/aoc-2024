package day2_test

import (
	"aoc-2024/day2"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(2)

	result := day2.SolvePuzzle1(input)

	if result != 572 {
		t.Error(result)
	}
}
