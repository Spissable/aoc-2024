package day1_test

import (
	"aoc-2024/day1"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(1)

	result := day1.SolvePuzzle1(input)

	if result != 2000468 {
		t.Error(result)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(1)

	result := day1.SolvePuzzle2(input)

	if result != 18567089 {
		t.Error(result)
	}
}
