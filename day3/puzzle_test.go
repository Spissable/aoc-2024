package day3_test

import (
	"aoc-2024/day3"
	"aoc-2024/util"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := util.ReadInput(3)

	result := day3.SolvePuzzle1(input)

	if result != 173419328 {
		t.Error(result)
	}
}

func TestSolvePuzzle2(t *testing.T) {
	input := util.ReadInput(3)

	result := day3.SolvePuzzle2(input)

	if result != 90669332 {
		t.Error(result)
	}
}
