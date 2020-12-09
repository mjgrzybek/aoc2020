package day8

import (
	"bufio"
	"github.com/bradfitz/iter"
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
	//fmt.Println(i)
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
	var listing []Instruction

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

func detectCycle(listing []Instruction, processor *Processor) *Instruction {
	visited := make([]bool, len(listing))

	for processor.Ip < len(listing) {
		instruction := listing[processor.Ip]

		if visited[processor.Ip] != true {
			visited[processor.Ip] = true
		} else {
			return &instruction
		}

		processor.execute(instruction)
	}

	return nil
}

func Solve1() int {
	f, _ := os.Open("day8/input.txt")
	listing := loadData(f)
	r, _ := execute(listing)
	return r
}

func execute(listing []Instruction) (int, *Instruction) {
	processor := &Processor{}
	instr := detectCycle(listing, processor)

	return processor.Acc, instr
}

func switchNopJmp(op *operation) {
	//fmt.Println(*op)
	switch *op {
	case "nop":
		*op = "jmp"
	case "jmp":
		*op = "nop"
	}
}

func Solve2() int {
	f, _ := os.Open("day8/input.txt")
	listing := loadData(f)

	for x := range iter.N(len(listing)) {
		if x>0 {
			// revert previous change
			switchNopJmp(&listing[x-1].operation)
		}
		switchNopJmp(&listing[x].operation)

		r, instruction := execute(listing)
		if instruction == nil {
			return r
		}
	}

	panic("no solution found")
}
