package day7_test

import (
	"aoc-2024/day7"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(7)

	result := day7.SolvePuzzle1(input)

	if result != 1038838357795 {
		t.Error(result)
	}
}
