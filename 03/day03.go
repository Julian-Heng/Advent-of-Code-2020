package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, err := readLines("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve(lines, 3, 1))

	p2 := 1
	p2 *= solve(lines, 1, 1)
	p2 *= solve(lines, 3, 1)
	p2 *= solve(lines, 5, 1)
	p2 *= solve(lines, 7, 1)
	p2 *= solve(lines, 1, 2)

	fmt.Println(p2)
}

func solve(lines []string, roffset int, doffset int) int {
	r := 0
	d := 0
	trees := 0
	for d < len(lines) {
		if lines[d][r] == '#' {
			trees++
		}

		r = (r + roffset) % len(lines[d])
		d += doffset
	}
	return trees
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
