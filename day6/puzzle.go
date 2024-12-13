package day6

import (
	"slices"
	"strings"
)

type Coord struct {
	row int
	col int
}

type Field struct {
	width      int
	height     int
	direction  rune
	currentPos Coord
	obstacles  []Coord
}

func NewField(input string) (result Field) {
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		for col, field := range line {
			if field == '.' {
				// empty field
				continue
			}
			if field == '#' {
				// obstacle
				result.obstacles = append(result.obstacles, Coord{row: row, col: col})
			}
			if field == '^' || field == '>' || field == 'v' || field == '<' {
				result.direction = field
				result.currentPos = Coord{row: row, col: col}
			}
		}
	}
	result.width = len(lines[0])
	result.height = len(lines)

	return
}

func SolvePuzzle1(input string) int {
	field := NewField(input)
	path := field.simulateWalk()
	return len(path)
}

var direction = map[rune]struct {
	row int
	col int
}{
	'^': {row: -1, col: 0}, // up
	'v': {row: 1, col: 0},  // down
	'>': {row: 0, col: 1},  // right
	'<': {row: 0, col: -1}, // left
}

func (f Field) simulateWalk() (result map[Coord][]rune) {
	result = map[Coord][]rune{}
	// starting position needs to be there
	result[f.currentPos] = []rune{f.direction}
	for {
		switch f.direction {
		case '^':
			tmpPos := Coord{
				row: f.currentPos.row - 1,
				col: f.currentPos.col,
			}
			if !slices.Contains(f.obstacles, tmpPos) {
				// no collision, move up
				f.currentPos = Coord{
					row: tmpPos.row,
					col: tmpPos.col,
				}
			} else {
				// collision turn right
				f.currentPos = Coord{
					row: f.currentPos.row,
					col: f.currentPos.col + 1,
				}
				f.direction = '>'
			}
		case '>':
			tmpPos := Coord{
				row: f.currentPos.row,
				col: f.currentPos.col + 1,
			}
			if !slices.Contains(f.obstacles, tmpPos) {
				// no collision, move up
				f.currentPos = Coord{
					row: tmpPos.row,
					col: tmpPos.col,
				}
			} else {
				// collision turn right
				f.currentPos = Coord{
					row: f.currentPos.row + 1,
					col: f.currentPos.col,
				}
				f.direction = 'v'
			}
		case 'v':
			tmpPos := Coord{
				row: f.currentPos.row + 1,
				col: f.currentPos.col,
			}
			if !slices.Contains(f.obstacles, tmpPos) {
				// no collision, move up
				f.currentPos = Coord{
					row: tmpPos.row,
					col: tmpPos.col,
				}
			} else {
				// collision turn right
				f.currentPos = Coord{
					row: f.currentPos.row,
					col: f.currentPos.col - 1,
				}
				f.direction = '<'
			}
		case '<':
			tmpPos := Coord{
				row: f.currentPos.row,
				col: f.currentPos.col - 1,
			}
			if !slices.Contains(f.obstacles, tmpPos) {
				// no collision, move up
				f.currentPos = Coord{
					row: tmpPos.row,
					col: tmpPos.col,
				}
			} else {
				// collision turn right
				f.currentPos = Coord{
					row: f.currentPos.row - 1,
					col: f.currentPos.col,
				}
				f.direction = '^'
			}
		}

		if f.currentPos.row < 0 || f.currentPos.row >= f.width || f.currentPos.col < 0 || f.currentPos.col >= f.height {
			// Job's done - out of bounds
			return
		}

		if slices.Contains(result[f.currentPos], f.direction) {
			// Job's done - closed the loop
			return
		}

		result[f.currentPos] = []rune{f.direction}
	}
}
