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

	height := len(lines)
	width := len(lines[0])

	trees := 0
	for i:=0; i<height; i++ {
		idx := (i * 3) % width
		if lines[i][idx] == true {
			trees++
		}
	}

	return trees
}
