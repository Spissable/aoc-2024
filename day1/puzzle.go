package day1

import (
	"aoc-2024/util"
	"sort"
	"strings"
)

type Locations struct {
	left  []int
	right []int
}

func SolvePuzzle1(input string) int {
	l := NewLocations(input)
	return l.getTotalDistance()
}

func SolvePuzzle2(input string) int {
	l := NewLocations(input)
	return l.getSimilarity()
}

func NewLocations(input string) (result Locations) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		result.left = append(result.left, util.StringToNum(parts[0]))
		result.right = append(result.right, util.StringToNum(parts[1]))
	}

	return
}

func (l *Locations) sort() {
	sort.Slice(l.left, func(i, j int) bool {
		return l.left[i] < l.left[j]
	})

	sort.Slice(l.right, func(i, j int) bool {
		return l.right[i] < l.right[j]
	})
}

func (l *Locations) getTotalDistance() (result int) {
	l.sort()

	for i, left := range l.left {
		result += util.IntDiff(left, l.right[i])
	}

	return
}

func (l *Locations) getSimilarity() (result int) {
	l.sort()

	for _, left := range l.left {
		counter := 0
		for _, right := range l.right {
			if left == right {
				counter++
			} else if left < right {
				break
			}
		}
		result += left * counter
	}

	return
}
