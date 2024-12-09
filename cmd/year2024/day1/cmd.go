package day1

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  "day1",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/1.txt", parent, command))
	if err != nil {
		logrus.Fatalf("error reading input: %v", err)
	}

	logrus.Infof("distance: %d", part1(string(b)))
	logrus.Infof("similiarity score: %d", part2(string(b)))
}

func part1(s string) int {
	var distance int
	row1, row2 := parse(s)

	for i := 0; i < len(row1); i++ {
		distance += int(math.Abs(float64((row1[i] - row2[i]))))
	}

	return distance
}

func part2(s string) int {
	var score int
	row1, row2 := parse(s)

	for i := 0; i < len(row1); i++ {
		for j, num := range row2 {
			if row1[i] == row2[j] {
				score += num
			}
		}
	}

	return score
}

func parse(s string) ([]int, []int) {
	var lines [][]int
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		var lineNums []int
		for _, num := range strings.Split(line, "   ") {
			i, _ := strconv.Atoi(num)
			lineNums = append(lineNums, i)
		}
		lines = append(lines, lineNums)
	}

	var row1 []int
	var row2 []int

	for _, line := range lines {
		row1 = append(row1, line[0])
		row2 = append(row2, line[1])
	}

	sort.IntSlice(row1).Sort()
	sort.IntSlice(row2).Sort()

	return row1, row2
}
