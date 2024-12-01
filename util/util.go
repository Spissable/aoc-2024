package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(day int) string {
	folder, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/../day%d/input.txt", folder, day)
	content, err := os.ReadFile(filePath)

	if err != nil {
		panic(fmt.Sprintf("File not found for path: %s", err))
	}

	return string(content)
}

func StringToNum(str string) int {
	num, _ := strconv.Atoi(strings.Trim(str, " "))
	return num
}

func StringsToNums(strs []string) []int {
	var nums []int

	for _, str := range strs {
		nums = append(nums, StringToNum(str))
	}

	return nums
}

func IntDiff(a, b int) int {
	if b > a {
		return b - a
	}

	return a - b
}
