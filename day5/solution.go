package day5

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	Col, Row string
}

func NewCoord(s string) *Coord {
	split := strings.IndexAny(s, "RL")

	return &Coord{
		Row: s[:split],
		Col: s[split:],
	}
}

func loadData(reader io.Reader) []*Coord {
	scanner := bufio.NewScanner(reader)

	data := make([]*Coord, 0)
	for scanner.Scan() {
		data = append(data, NewCoord(scanner.Text()))
	}

	return data
}

func (c *Coord) String() string {
	ret := c.Row + " " + c.Col
	ret += " = " + strconv.Itoa(BinarySpacePartitioning(c.Row, 'F', 'B'))
	ret += " " + strconv.Itoa(BinarySpacePartitioning(c.Col, 'L', 'R'))
	ret += " = SID: " + strconv.Itoa(c.getSeatID())
	return ret
}

func (c *Coord) getSeatID() int {
	return BinarySpacePartitioning(c.Row, 'F', 'B') * 8 + BinarySpacePartitioning(c.Col, 'L', 'R')
}

func BinarySpacePartitioning(inputValue string, loKey rune, hiKey rune) int {
	inputValue = strings.Replace(inputValue, string(loKey), "0", -1)
	inputValue = strings.Replace(inputValue, string(hiKey), "1", -1)
	i, err := strconv.ParseInt(inputValue, 2, 32)

	if err != nil {
		panic(err)
	}

	return int(i)
}

func Solve1() int {
	f, _ := os.Open("day5/input")
	coords := loadData(f)

	max := math.MinInt32

	for _, coord := range coords {
		//fmt.Println(coord)
		sid := coord.getSeatID()
		if sid > max {
			max = sid
		}
		if sid == 1023 {
			fmt.Println(coord)
		}
	}

	return max
}

func Solve2() int {
	f, _ := os.Open("day5/input")
	coords := loadData(f)

	seats := [128][8]bool{}

	for _, coord := range coords {
		c := BinarySpacePartitioning(coord.Col, 'L', 'R')
		r := BinarySpacePartitioning(coord.Row, 'F', 'B')

		seats[r][c] = true
	}

	for r := range seats {
		for c := range seats[r] {
			if seats[r][c] == false {
				fmt.Println(r, c, "...", r*8+c)
			}
		}
	}

	return 0
}
