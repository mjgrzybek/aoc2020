package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type operation string
type argument string

type Instruction struct {
	operation
	argument
}

type Processor struct {
	Acc int
	Ip  int
}

func (p *Processor) execute(i Instruction) {
	fmt.Println(i)
	instrOffset := 1

	switch i.operation {
	case "nop":
	case "acc":
		arg, _ := strconv.Atoi(string(i.argument))
		p.Acc += arg
	case "jmp":
		arg, _ := strconv.Atoi(string(i.argument))
		instrOffset = arg
	default:
		panic("Instruction: " + i.operation + " is not supported")
	}

	p.Ip += instrOffset
}

func loadData(r io.Reader) []Instruction {
	listing := []Instruction{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Fields(txt)
		listing = append(listing, Instruction{
			operation: operation(fields[0]),
			argument:  argument(fields[1]),
		})
	}

	return listing
}

func Solve1() int {
	f, _ := os.Open("day8/input.txt")
	listing := loadData(f)

	processor := Processor{}

	visited := make([]bool, len(listing))

	for processor.Ip < len(listing) {
		if visited[processor.Ip] != true {
			visited[processor.Ip] = true
		} else {
			break
		}

		instruction := listing[processor.Ip]
		processor.execute(instruction)
	}

	return processor.Acc
}

func Solve2() int {
	return -1
}
