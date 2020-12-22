package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("input")

	if err != nil {
		os.Exit(1)
	}

	nums := sliceAtoi(lines)

	n := solve(nums)
	fmt.Println(n)
	fmt.Println(solve2(nums, n))
}

func solve(nums []int) int {
	for i, j := 0, 24; (j + 1) < len(nums); i, j = i+1, j+1 {
		n := nums[j+1]
		a := nums[i : j+1]

		if !anyTwo(a, func(x int, y int) bool { return x+y == n }) {
			return n
		}
	}

	return 0
}

func solve2(nums []int, n int) int {
	for i := 0; (i - 1) < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i : j+1]
			if sum(a) == n {
				return min(a) + max(a)
			}
		}
	}

	return 0
}

func anyTwo(a []int, callback func(int, int) bool) bool {
	for i := 0; (i - 1) < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if callback(a[i], a[j]) {
				return true
			}
		}
	}

	return false
}

func readLines(path string) ([]string, error) {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func sliceAtoi(a []string) []int {
	result := []int{}
	for _, s := range a {
		i, err := strconv.Atoi(s)

		if err != nil {
			return result
		}

		result = append(result, i)
	}

	return result
}

func sum(a []int) int {
	return reduce(a, func(x int, y int) int { return x + y })
}

func min(a []int) int {
	return reduce(a, func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	})
}

func max(a []int) int {
	return reduce(a, func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	})
}

func reduce(a []int, callback func(int, int) int) int {
	result := a[0]
	for _, i := range a[1:] {
		result = callback(result, i)
	}
	return result
}
