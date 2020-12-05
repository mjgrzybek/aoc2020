package day3

import (
	"bufio"
	"io"
	"os"
)

func loadData(reader io.Reader) [][]bool {
	scanner := bufio.NewScanner(reader)

	lines := make([][]bool, 0)
	for scanner.Scan() {
		line := make([]bool, len(scanner.Text()))
		for pos, char := range scanner.Text() {
			isTree := char == '#'
			line[pos] = isTree
		}
		lines = append(lines, line)
	}
	return lines
}

func Solve1() int {
	f, _ := os.Open("day3/input")
	lines := loadData(bufio.NewReader(f))

	trees := calculateCollisions(lines, 1, 3)

	return trees
}

func Solve2() uint64 {
	f, _ := os.Open("day3/input")
	lines := loadData(bufio.NewReader(f))

	slopes := []struct {
		x int
		y int
	}{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	var tressMultiplied uint64 = 1
	for _, xy := range slopes {
		mul := calculateCollisions(lines, xy.y, xy.x)
		tressMultiplied *= uint64(mul)
	}

	return tressMultiplied
}

func calculateCollisions(lines [][]bool, y, x int) int {
	height := len(lines)
	width := len(lines[0])

	trees := 0
	jumps := 0
	for i := 0; i < height; i += y {
		idx := (jumps * x) % width
		if lines[i][idx] == true {
			trees++
		}
		jumps++
	}
	return trees
}
