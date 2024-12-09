package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  "day3",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/1.txt", parent, command))
	if err != nil {
		logrus.Fatalf("error reading input: %v", err)
	}

	logrus.Infof("mult result part 1: %d", part1(string(b)))
	logrus.Infof("mult result part 2: %d", part2(string(b)))
}

func part1(s string) int {
	prod := 0
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(string(s), -1)

	for _, match := range matches {
		nums := strings.Split((match[4 : len(match)-1]), ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		prod += num1 * num2
	}

	return prod
}

func part2(s string) int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)|do[n't]*\(\)`)
	matches := r.FindAllString(string(s), -1)
	prod := 0
	do := true

	for _, match := range matches {
		if match == "don't()" {
			do = false
			continue
		}

		if match == "do()" {
			do = true
			continue
		}

		if do {
			nums := strings.Split((match[4 : len(match)-1]), ",")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			prod += num1 * num2
		}
	}

	return prod
}
