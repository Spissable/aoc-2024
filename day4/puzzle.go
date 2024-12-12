package day4

import (
	"slices"
	"strings"
)

type Coord struct {
	LineNum int
	ElemNum int
}

type CharCoords map[rune][]Coord

type Direction struct {
	lineAdd int
	elemAdd int
}

var directions = []Direction{
	{
		// right
		lineAdd: 0,
		elemAdd: 1,
	},
	{
		// left
		lineAdd: 0,
		elemAdd: -1,
	},
	{
		// down
		lineAdd: 1,
		elemAdd: 0,
	},
	{
		// up
		lineAdd: -1,
		elemAdd: 0,
	},
	{
		// up right
		lineAdd: -1,
		elemAdd: 1,
	},
	{
		// up left
		lineAdd: -1,
		elemAdd: -1,
	},
	{
		// down right
		lineAdd: 1,
		elemAdd: 1,
	},
	{
		// down left
		lineAdd: 1,
		elemAdd: -1,
	},
}

func NewCharCoords(input string) (result CharCoords) {
	result = map[rune][]Coord{}
	lines := strings.Split(input, "\n")

	for l, line := range lines {
		for e, elem := range line {
			result[elem] = append(result[elem], Coord{LineNum: l, ElemNum: e})
		}
	}

	return
}

func SolvePuzzle1(input string) (result int) {
	searchString := "XMAS"
	base := NewCharCoords(input)

	startCoords := base[rune(searchString[0])]
	for _, x := range startCoords {
		for _, dir := range directions {
			match := true
			for i := 1; i < len(searchString); i++ {
				if !slices.Contains(base[rune(searchString[i])], Coord{LineNum: x.LineNum + i*dir.lineAdd, ElemNum: x.ElemNum + i*dir.elemAdd}) {
					match = false
					break
				}
			}

			if match {
				result++
			}
		}
	}

	return
}
