package main

import (
	"fmt"
	"github.com/mjgrzybek/aoc2020/day1"
	"github.com/mjgrzybek/aoc2020/day2"
	"github.com/mjgrzybek/aoc2020/day3"
	"os"
)

func main() {
	dayNo := os.Args[1]

	switch dayNo {
	case "1":
		fmt.Println(day1.Solve1())
		fmt.Println(day1.Solve2())
	case "2":
		fmt.Println(day2.Solve1())
		fmt.Println(day2.Solve2())
	case "3":
		fmt.Println(day3.Solve1())
	}
}
