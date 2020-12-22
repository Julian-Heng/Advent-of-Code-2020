package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	target, buses, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve1(target, buses))
	fmt.Println(solve2(buses))
}

func solve1(target int, buses []int) int {
	bestTime := int(^uint(0) >> 1)
	busID := 0

	for _, b := range buses {
		if b == -1 {
			continue
		}

		time := (b * int(math.Round(float64(target)/float64(b)))) - target
		if time < 0 {
			continue
		}

		if time < bestTime {
			bestTime = time
			busID = b
		}
	}

	return bestTime * busID
}

func solve2(buses []int) int {
	n := []int{}
	a := []int{}

	for i, j := range buses {
		if j == -1 {
			continue
		}

		n = append(n, j)
		a = append(a, j-i)
	}

	return crt(n, a)
}

func crt(n []int, a []int) int {
	sum := 0
	prod := prod(n)

	for i := range n {
		ni := n[i]
		ai := a[i]
		p := prod / ni
		sum += ai * modInv(p, ni) * p
	}

	return sum % prod
}

func modInv(a int, b int) int {
	b0 := b
	x0 := 0
	x1 := 1

	if b == 1 {
		return 1
	}

	for a > 1 {
		q := a / b

		tmp := a
		a = b
		b = tmp % b

		tmp = x0
		x0 = x1 - q*x0
		x1 = tmp
	}

	if x1 < 0 {
		x1 += b0
	}

	return x1
}

func parse(path string) (int, []int, error) {
	target := 0
	buses := []int{}

	file, err := os.Open(path)

	if err != nil {
		return target, buses, err
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return target, buses, scanner.Err()
	}

	target = atoi(lines[0])
	for _, s := range strings.SplitN(lines[1], ",", -1) {
		buses = append(buses, atoi(s))
	}

	return target, buses, nil
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}

func prod(a []int) int {
	return reduce(a, func(x int, y int) int { return x * y })
}

func reduce(a []int, callback func(int, int) int) int {
	result := a[0]
	for _, i := range a[1:] {
		result = callback(result, i)
	}
	return result
}
