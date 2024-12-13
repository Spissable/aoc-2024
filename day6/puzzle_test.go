package day6_test

import (
	"aoc-2024/day6"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(6)

	result := day6.SolvePuzzle1(input)

	if result != 5269 {
		t.Error(result)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(6)

	result := day6.SolvePuzzle2(input)

	if result != 1957 {
		t.Error(result)
	}
}
