package day6

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type RuneSetT = map[rune]struct{}

func loadData1(reader io.Reader) []RuneSetT {
	scanner := bufio.NewScanner(reader)

	groups := make([]RuneSetT, 0)
	firstLoop := true
	for scanner.Scan() {
		text := scanner.Text()
		if firstLoop || strings.EqualFold(text, "") {
			groups = append(groups, make(RuneSetT))
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
	groups := loadData1(f)

	sum := 0

	for _, group := range groups {
		sum += len(group)
	}
	return sum
}

func processData2(reader io.Reader) []RuneSetT {
	scanner := bufio.NewScanner(reader)

	intersections := make([]RuneSetT, 0)

	inGroupResponses := make([]RuneSetT, 0)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			intersections = append(intersections, calculateIntersection(inGroupResponses))

			// reset state
			inGroupResponses = make([]RuneSetT, 0)
			continue
		}

		m := make(RuneSetT)
		for _, letter := range text {
			m[letter] = struct{}{}
		}

		inGroupResponses = append(inGroupResponses, m)
	}

	return append(intersections, calculateIntersection(inGroupResponses))
}


func calculateIntersection(responses []RuneSetT) RuneSetT {
	// end of group, run intersection on lines in group
	intersection := responses[0]
	for i := 1; i < len(responses); i++ {
		for r, _ := range intersection {
			_, exists := responses[i][r]
			if !exists {
				delete(intersection, r)
			}
		}
	}
	return intersection
}

func Solve2() int {
	f, _ := os.Open("day6/input")
	intersections := processData2(f)

	sum := 0
	for _, intersection := range intersections {
		sum += len(intersection)
	}

	return sum
}
