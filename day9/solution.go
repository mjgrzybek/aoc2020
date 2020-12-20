package day9

import (
	"bufio"
	"os"
	"strconv"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Sum2Map struct {
	Map map[int]int
}

func NewSum2Map() *Sum2Map {
	m := &Sum2Map{}
	m.Map = make(map[int]int)
	return m
}

func (m *Sum2Map) isValid(num int) bool {
	// a + b = c
	// a + b - c = 0
	
	return m.Map[]
}

func (s *Sum2Map) Update(removed, added int) {

}

type CyclicBuffer struct {
	buffer  []int
	lastIdx int
	size    int
	cap     int
}

func NewCyclicBuffer(size int) *CyclicBuffer {
	b := &CyclicBuffer{size: size}
	b.buffer = make([]int, size, size)
	return b
}

func (b *CyclicBuffer) Push(val int) {
	b.cap = Min(b.cap+1, b.size)
	b.buffer[b.lastIdx] = val
	b.lastIdx = (b.lastIdx + 1) % b.size
}

func (b *CyclicBuffer) Get(idx int) int {
	if b.cap < b.size {
		return b.buffer[idx]
	}
	return b.buffer[(b.lastIdx+idx)%b.size]
}

func (b *CyclicBuffer) Size() int {
	return b.size
}

func Solve1() int {
	f, _ := os.Open("day9/simple.txt")
	reader := bufio.NewReader(f)
	const bufferSize = 5
	return firstInvalidValue(reader, bufferSize)
}

func firstInvalidValue(r *bufio.Reader, size int) int {
	b := NewCyclicBuffer(size)

	// preamble
	scanner := bufio.NewScanner(r)
	for i := 0; i<size && scanner.Scan(); i++ {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		b.Push(num)
	}

	sum2map := NewSum2Map()
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		if sum2map.isValid(num) == false {
			return num
		}
		sum2map.Update(b.Get(b.Size()-1), num)
		b.Push(num)
	}
	return -1
}

func Solve2() int {
	return -1
}
