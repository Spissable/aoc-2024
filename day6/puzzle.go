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
	path, _ := field.simulateWalk()
	return len(path)
}

func SolvePuzzle2(input string) (result int) {
	field := NewField(input)
	path, _ := field.simulateWalk()

	// remove the player position since we can't put an obstacle there
	delete(path, field.currentPos)

	for candidate := range path {
		if field.simulateWithNewObstacle(candidate) {
			result++
		}
	}
	return result
}

func (f Field) simulateWithNewObstacle(newObstacle Coord) bool {
	// set the new obstacle
	f.obstacles = append(f.obstacles, newObstacle)

	_, isLoop := f.simulateWalk()
	return isLoop
}

func (f Field) simulateWalk() (result map[Coord][]rune, isLoop bool) {
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
				f.direction = '^'
			}
		}

		if f.currentPos.row < 0 || f.currentPos.row >= f.width || f.currentPos.col < 0 || f.currentPos.col >= f.height {
			// Job's done - out of bounds
			isLoop = false
			return
		}

		if slices.Contains(result[f.currentPos], f.direction) {
			// Job's done - closed the loop
			isLoop = true
			return
		}

		result[f.currentPos] = append(result[f.currentPos], f.direction)
	}
}
