package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLines("input")

	if err != nil {
		os.Exit(1)
	}

	bags := makeBags(lines)
	fmt.Println(solve1(bags))
	fmt.Println(solve2(bags))
}

func solve1(bags map[string]map[string]int) int {
	n := 0
	seen := map[string]bool{}
	for k := range bags {
		if traverse(bags, k, seen) {
			n++
		}
	}
	return n
}

func solve2(bags map[string]map[string]int) int {
	seen := map[string]int{}
	return traverse2(bags, "shiny gold", seen)
}

func traverse(
	bags map[string]map[string]int,
	bag string,
	seen map[string]bool,
) bool {
	if v, ok := seen[bag]; ok {
		return v
	}

	for k := range bags[bag] {
		if k == "shiny gold" {
			seen[bag] = true
			return true
		}
	}

	for k := range bags[bag] {
		if traverse(bags, k, seen) {
			seen[bag] = true
			return true
		}
	}

	seen[bag] = false
	return false
}

func traverse2(
	bags map[string]map[string]int,
	bag string,
	seen map[string]int,
) int {
	if v, ok := seen[bag]; ok {
		return v
	}

	n := 0
	for k := range bags[bag] {
		v := bags[bag][k]
		n += v + (v * traverse2(bags, k, seen))
	}

	seen[bag] = n
	return n
}

func makeBags(lines []string) map[string]map[string]int {
	bags := map[string]map[string]int{}
	re := regexp.MustCompile(`(\d+)\s(\w+\s\w+)`)
	for _, l := range lines {
		split := strings.SplitN(l, " bags contain ", 2)
		match := re.FindAllStringSubmatch(split[1], -1)

		subBags := map[string]int{}

		if len(match) > 0 {
			for _, m := range match {
				subBags[m[2]] = atoi(m[1])
			}
		}

		bags[split[0]] = subBags
	}

	return bags
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

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
