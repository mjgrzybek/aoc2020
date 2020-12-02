package day1

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

// O(nlogn)
func SolvePart1() int {
	nums := loadData()

	// O(n)
	numsMap := map[int]bool{} // mimic set datastructure
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

// O(n^2 logn)
func SolvePart2() int {
	nums := loadData()

	// O(n^2 logn)
	nums2product := map[int]int{}
	for i, n := range nums {
		for _, m := range nums[i+1:] {
			nums2product[n+m] = n*m
		}
	}
	// O(n)
	for _, n := range nums {
		diff := 2020 - n
		// O(logn)
		if nums2product[diff] > 0 {
			return n * nums2product[diff]
		}
	}

	return -1
}


func loadData() []int {
	data, _ := ioutil.ReadFile("day1/input")

	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	// load data, O(n)
	nums := make([]int, 0)
	for scanner.Scan() {
		atoi, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, atoi)
	}
	return nums
}
