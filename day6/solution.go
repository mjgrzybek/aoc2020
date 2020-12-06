package day6

import (
	"bufio"
	"io"
	"os"
	"strings"
)



func loadData(reader io.Reader) []map[rune]struct{} {
	scanner := bufio.NewScanner(reader)

	groups := make([]map[rune]struct{}, 0)
	firstLoop := true
	for scanner.Scan() {
		text := scanner.Text()
		if firstLoop || strings.EqualFold(text, "") {
			groups = append(groups, make(map[rune]struct{}))
			firstLoop = false
		}

		group := groups[len(groups)-1]

		for _, r := range text {
			group[r] = struct{}{}
		}
	}

	return groups
}

func Solve1() int {
	f, _ := os.Open("day6/input")
	groups := loadData(f)

	sum :=0

	for _, group := range groups {
		sum += len(group)
	}
	return sum
}

func Solve2() int {
	return -1
}
