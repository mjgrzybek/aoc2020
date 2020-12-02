package day1

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

// O(nlogn)
func Solve() int {
	data, _ := ioutil.ReadFile("day1/input")

	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	// load data, O(n)
	nums := make([]int, 0)
	for scanner.Scan() {
		atoi, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, atoi)
	}

	// O(n)
	numsMap := map[int]bool{} 		// mimic set datastructure
	for _, n := range nums {
		// O(logn)
		numsMap[n] = true
	}

	// O(n)
	for _, n := range nums {
		diff := 2020 - n
		// O(logn)
		if numsMap[diff] == true {
			return n * diff
		}
	}

	return -1
}
