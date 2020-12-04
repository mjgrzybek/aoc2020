package day2

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Solve1() int {
	entries := loadData()

	valid := 0

	for _, e := range entries {
		fmt.Println(e)
		if e.isValid1() {
			valid++
		}
	}

	return valid
}

func Solve2() int {
	entries := loadData()

	valid := 0

	for _, e := range entries {
		//fmt.Println(e)
		if e.isValid2() {
			valid++
		}
	}

	return valid
}

type entry struct {
	Min, Max int
	Letter rune
	Password string
}

func (e entry) isValid1() bool {
	ctr:=0
	for _, v := range e.Password {
		if v==e.Letter {
			ctr++
		}
	}
	return ctr >= e.Min && ctr <= e.Max
}

func (e entry) isValid2() bool {
	if len(e.Password) < e.Min {
		return false
	}

	a := rune(e.Password[e.Min-1])
	b := rune(e.Password[e.Max-1])

	if len(e.Password) < e.Max {
		return a == e.Letter
	}

	return (a == e.Letter || b == e.Letter) && (a != b)
}

func (e entry) String() string {
	return fmt.Sprintf("%d-%d %c: %s %v", e.Min, e.Max, e.Letter, e.Password, e.isValid1())
}

func loadData() []entry {
	data, _ := ioutil.ReadFile("day2/input")

	entriesScanner := bufio.NewScanner(strings.NewReader(string(data)))

	entries := make([]entry, 0)
	for entriesScanner.Scan() {
		fields := strings.FieldsFunc(entriesScanner.Text(), func(r rune) bool { return r == ' ' } )
		minmax := strings.Split(fields[0], "-")
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		letter := rune(fields[1][0])
		password := fields[2]
		entries = append(entries, entry{min, max, letter, password})
	}

	return entries
}
