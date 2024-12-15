package day5

import (
	"aoc-2024/util"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Rules map[int][]int

type Update []int

type Input struct {
	beforeRules Rules
	updates     []Update
}

type SingleUpdate struct {
	rules  Rules
	update Update
}

func (b SingleUpdate) Len() int {
	return len(b.update)
}

func (b SingleUpdate) Less(i, j int) bool {
	rulesBefore := b.rules[b.update[j]]
	if slices.Contains(rulesBefore, b.update[i]) {
		return false
	}

	rulesNotBefore := b.rules[b.update[i]]
	return slices.Contains(rulesNotBefore, b.update[j])
}

func (b SingleUpdate) Swap(i, j int) {
	b.update[i], b.update[j] = b.update[j], b.update[i]
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

	return sumMiddlePage(data.updatesInOrder(true))
}

func SolvePuzzle2(input string) int {
	data := NewInput(input)

	unordered := data.updatesInOrder(false)
	sorted := make([]Update, 0, len(unordered))
	for _, u := range unordered {
		tmp := SingleUpdate{
			rules:  data.beforeRules,
			update: u,
		}
		sort.Sort(tmp)
		sorted = append(sorted, tmp.update)
	}
	return sumMiddlePage(sorted)
}

func (i Input) updatesInOrder(order bool) (result []Update) {
	for _, update := range i.updates {
		if i.updateOk(update) == order {
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
