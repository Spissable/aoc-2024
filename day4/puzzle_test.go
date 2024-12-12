package day4_test

import (
	"aoc-2024/day4"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(4)

	result := day4.SolvePuzzle1(input)

	if result != 2483 {
		t.Error(result)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(4)

	result := day4.SolvePuzzle2(input)

	if result != 1925 {
		t.Error(result)
	}
}
