package day2

import (
	"aoc-2024/util"
	"strings"
)

type Report []int

type Reports struct {
	Report []Report
}

func NewReports(input string) Reports {
	lines := strings.Split(input, "\n")
	reports := Reports{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		report := make(Report, 0)
		for _, part := range parts {
			report = append(report, util.StringToNum(part))
		}
		reports.Report = append(reports.Report, report)
	}

	return reports
}

func SolvePuzzle1(input string) int {
	reports := NewReports(input)

	result := 0
	for _, report := range reports.Report {
		if report.isSafe() {
			result++
		}
	}

	return result
}

func SolvePuzzle2(input string) int {
	reports := NewReports(input)

	result := 0
	for _, report := range reports.Report {
		if report.isSafe() {
			result++
			continue
		}

		for i := range report {
			damped := append(append(Report{}, report[:i]...), report[i+1:]...)
			if damped.isSafe() {
				result++
				break
			}
		}
	}

	return result
}

func (r Report) isSafe() bool {
	asc := r[1] > r[0]

	for i := 0; i < len(r)-1; i++ {
		diff := r[i+1] - r[i]
		if diff == 0 {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
		if asc && diff < 0 {
			return false
		}
		if !asc && diff > 0 {
			return false
		}
	}

	return true
}
