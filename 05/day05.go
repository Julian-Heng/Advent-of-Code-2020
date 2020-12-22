package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines, err := readLines("input")

	if err != nil {
		os.Exit(1)
	}

	seats := solve(lines)
	sort.Ints(seats[:])

	// Part 1
	fmt.Println(seats[len(seats)-1])

	// Part 2
	first := seats[0]
	for _, second := range seats[1:] {
		if second-first != 1 {
			break
		}
		first = second
	}
	fmt.Println(first + 1)
}

func solve(lines []string) []int {
	seats := []int{}

	for _, l := range lines {
		fb := l[:7]
		lr := l[len(l)-3:]

		row := traverse(fb, 0, 127, 'B', 'F')
		col := traverse(lr, 0, 7, 'R', 'L')

		seats = append(seats, (row*8)+col)
	}

	return seats
}

func traverse(p string, l int, h int, upper byte, lower byte) int {
	if len(p) == 0 {
		return l
	}

	offset := ((h - l) / 2) + 1

	l2 := l
	h2 := h

	if p[0] == upper {
		l2 += offset
	} else {
		h2 -= offset
	}

	return traverse(p[1:], l2, h2, upper, lower)
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
