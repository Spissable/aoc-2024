package day5

import (
	"aoc-2024/util"
	"slices"
	"strconv"
	"strings"
)

type Rules map[int][]int

type Update []int

type Input struct {
	beforeRules Rules
	updates     []Update
}

func NewInput(input string) (result Input) {
	result = Input{
		beforeRules: map[int][]int{},
	}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		rules := strings.Split(line, "|")
		if len(rules) == 2 {
			a, _ := strconv.Atoi(rules[0])
			b, _ := strconv.Atoi(rules[1])

			result.beforeRules[a] = append(result.beforeRules[a], b)
		} else {
			updateLine := strings.Split(line, ",")
			if len(updateLine) > 1 {
				update := make(Update, 0, len(updateLine))
				for _, updateStr := range updateLine {
					update = append(update, util.StringToNum(updateStr))
				}
				result.updates = append(result.updates, update)
			}
		}
	}

	return
}

func SolvePuzzle1(input string) int {
	data := NewInput(input)

	return sumMiddlePage(data.updatesInOrder())
}

func (i Input) updatesInOrder() (result []Update) {
	for _, update := range i.updates {
		if i.updateOk(update) {
			result = append(result, update)
		}
	}

	return
}

func (i Input) updateOk(update Update) bool {
	for numI, num := range update {
		if !i.beforeOk(update, num, numI) {
			return false
		}
	}

	return true
}

func (i Input) beforeOk(update Update, num, numI int) bool {
	for _, rule := range i.beforeRules[num] {
		ruleI := slices.Index(update, rule)

		if ruleI != -1 && ruleI < numI {
			return false
		}
	}

	return true
}

func sumMiddlePage(updates []Update) (result int) {
	for _, update := range updates {
		result += update[len(update)/2]
	}

	return
}
