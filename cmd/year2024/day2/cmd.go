package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type reportsState int

const (
	increasing reportsState = iota
	decreasing
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  "day2",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/1.txt", parent, command))
	if err != nil {
		logrus.Fatalf("error reading input: %v", err)
	}

	logrus.Infof("part 1 safe reports: %d", part1(string(b)))
	logrus.Infof("part 2 safe reports: %d", part2(string(b)))
}

func part1(s string) int {
	var score int
	reports := parse(s)

	for _, report := range reports {
		safe := true
		var isIncreasing bool
		var isDecreasing bool
		for i := 0; i < len(report)-1; i++ {
			if report[i+1] < report[i] {
				isIncreasing = true
			}
			if report[i+1] > report[i] {
				isDecreasing = true
			}

			if isIncreasing && isDecreasing {
				safe = false
				break
			}

			difference := math.Abs(float64(report[i+1] - report[i]))
			if difference < 1 || difference > 3 {
				safe = false
				break
			}
		}

		if safe {
			score++
		}
	}

	return score
}

func part2(s string) int {
	var score int
	reports := parse(s)

	for _, report := range reports {
		if existsSafePermutation(report) {
			score++
		}
	}

	return score
}

func existsSafePermutation(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	// Not allocating an entire new array results in the original being modified,
	// weird Go behaviour, I guess?
	for i := range nums {
		newNums := make([]int, len(nums))
		copy(newNums, nums)

		if isSafe(append(newNums[:i], newNums[i+1:]...)) {
			return true
		}
	}

	return false
}

func isSafe(nums []int) bool {
	state := increasing
	if nums[0] > nums[1] {
		state = decreasing
	}

	prev := nums[0]
	for _, num := range nums[1:] {
		if num == prev ||
			(state == increasing && (num < prev || num > prev+3)) ||
			(state == decreasing && (num > prev || num < prev-3)) {
			return false
		}

		prev = num
	}

	return true
}

func parse(s string) [][]int {
	var lines [][]int
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		var lineNums []int
		for _, num := range strings.Split(line, " ") {
			i, _ := strconv.Atoi(num)
			lineNums = append(lineNums, i)
		}
		lines = append(lines, lineNums)
	}

	return lines
}
