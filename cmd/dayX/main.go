package main

import (
	"fmt"
	"github.com/mjgrzybek/aoc2020/day1"
	"os"
)

func main() {
	dayNo := os.Args[1]

	switch dayNo {
	case "1":
		fmt.Println(day1.Solve1())
		fmt.Println(day1.Solve2())
	}
}
