package day9

import (
	"bufio"
	"os"
	"sort"
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

type SortedSlice struct {
	Array []int
}

func NewSortedSlice(array []int) *SortedSlice {
	return &SortedSlice{Array: array}
}

func (s *SortedSlice) Update(removed, added int) {
	// naive impl O(n) because of no generics in go yet..
	found := -1
	for i, v := range s.Array {
		if v == removed {
			found = i
			break
		}
	}

	if found > -1 {
		s.Array = append(s.Array[:found], s.Array[found+1:]...)
	}

	s.Array = append(s.Array, added)
	sort.Ints(s.Array)
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

	// losf preamble
	scanner := bufio.NewScanner(r)
	for i := 0; i<size && scanner.Scan(); i++ {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		b.Push(num)
	}

	sortedView := SortedSlice{append([]int(nil), b.buffer...)}
	for scanner.Scan() {
		txt := scanner.Text()
		num, _ := strconv.Atoi(txt)
		if isValid(num, sortedView) == false {
			return num
		}
		sortedView.Update(b.Get(b.Size()-1), num)
		b.Push(num)
	}
	return -1
}

func isValid(num int, b SortedSlice) bool {
	for i, v := range b.Array {
		
	}
	return true
}

func Solve2() int {
	return -1
}
