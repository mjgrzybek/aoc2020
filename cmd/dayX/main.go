package main

import (
	"fmt"
	"github.com/mjgrzybek/aoc2020/day1"
	"github.com/mjgrzybek/aoc2020/day2"
	"github.com/mjgrzybek/aoc2020/day3"
	"github.com/mjgrzybek/aoc2020/day4"
	"github.com/mjgrzybek/aoc2020/day5"
	"github.com/mjgrzybek/aoc2020/day6"
	"github.com/mjgrzybek/aoc2020/day7"
	"os"
)

func main() {
	dayNo := "7"
	if len(os.Args) == 2 {
		dayNo = os.Args[1]
	}

	switch dayNo {
	case "1":
		fmt.Println(day1.Solve1())
		fmt.Println(day1.Solve2())
	case "2":
		fmt.Println(day2.Solve1())
		fmt.Println(day2.Solve2())
	case "3":
		fmt.Println(day3.Solve1())
		fmt.Println(day3.Solve2())
	case "4":
		fmt.Println(day4.Solve1())
		fmt.Println(day4.Solve2())
	case "5":
		fmt.Println(day5.Solve1())
		fmt.Println(day5.Solve2())
	case "6":
		fmt.Println(day6.Solve1())
		fmt.Println(day6.Solve2())
	case "7":
		fmt.Println(day7.Solve1())
		fmt.Println(day7.Solve2())
	}
}
