package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input
	strLines, err := read("input")

	if err != nil {
		os.Exit(1)
	}

	lines, err := sliceAtoi(strLines)

	if err != nil {
		os.Exit(1)
	}

	// Part 1
	fmt.Println(solve(lines, 2))

	// Part 2
	fmt.Println(solve(lines, 3))
}

func solve(nums []int, n int) int {
	result := 0

	// Define callback function for combinations
	check := func(i []int) bool {
		// If they sum to 2020, set the result to the product
		if sum(i) == 2020 {
			result = product(i)
			return true
		}
		return false
	}

	// Run combinations
	combinations(nums, n, check)
	return result
}

func combinations(l []int, n int, callback func([]int) bool) {
	s := make([]int, n)
	last := n - 1
	quit := false

	var helper func(int, int)
	helper = func(i int, next int) {
		if quit {
			return
		}

		for j := next; j < len(l); j++ {
			s[i] = l[j]
			if i == last {
				if callback(s) {
					quit = true
					break
				}
			} else {
				helper(i+1, j+1)
			}
		}
	}

	helper(0, 0)
}

func read(path string) ([]string, error) {
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

func sliceAtoi(a []string) ([]int, error) {
	result := make([]int, 0, len(a))
	for _, s := range a {
		i, err := strconv.Atoi(s)

		if err != nil {
			return result, err
		}

		result = append(result, i)
	}

	return result, nil
}

func sum(a []int) int {
	return reduce(a, func(x int, y int) int { return x + y })
}

func product(a []int) int {
	return reduce(a, func(x int, y int) int { return x * y })
}

func reduce(a []int, callback func(int, int) int) int {
	result := a[0]
	for _, i := range a[1:] {
		result = callback(result, i)
	}
	return result
}
