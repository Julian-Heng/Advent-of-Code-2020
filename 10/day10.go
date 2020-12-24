package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines, err := readLines("input")

	if err != nil {
		os.Exit(1)
	}

	nums := sliceAtoi(lines)
	sort.Ints(nums[:])
	nums = append(nums, aMax(nums)+3)
	nums = append([]int{0}, nums...)

	fmt.Println(solve1(nums))
	fmt.Println(solve2(nums))
}

func solve1(nums []int) int {
	one := 0
	three := 0
	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		switch nums[j] - nums[i] {
		case 1:
			one++

		case 3:
			three++
		}
	}

	return one * three
}

func solve2(nums []int) int {
	d := []int{}
	result := 1
	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		if (nums[j] - nums[i]) == 3 {
			d = append(d, i)
		}
	}

	i := -1
	for _, j := range d {
		n := max(1, j-i-2)
		e := 1
		if nums[j]-nums[i+1] == 3 || j-i == 3 {
			e = 0
		}
		result *= (1 << n) - e
		i = j
	}

	return result
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

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func aMax(a []int) int {
	return reduce(a, max)
}

func reduce(a []int, callback func(int, int) int) int {
	result := a[0]
	for _, i := range a[1:] {
		result = callback(result, i)
	}
	return result
}
