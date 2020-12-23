package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	num, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve(num[:], 100, 1))
	fmt.Println(solve(num[:], 10_000_000, 2))
}

func solve(num []int, limit int, part int) int {
	min := aMin(num)
	max := aMax(num)
	current := num[0]

	if part == 2 {
		for i := max + 1; i <= 1_000_000; i++ {
			num = append(num, i)
		}
		max = num[len(num)-1]
	}

	cups := map[int]int{}
	for i := 0; i < len(num); i++ {
		cups[num[i]] = num[(i+1)%len(num)]
	}

	for i := 0; i < limit; i++ {
		pickUp := []int{}
		next := current
		for j := 0; j < 3; j++ {
			pickUp = append(pickUp, cups[next])
			next = cups[next]
		}

		cups[current] = cups[next]
		dest := current - 1
		_, ok := cups[dest]
		for !ok || intSliceContains(pickUp, dest) {
			dest--
			if dest < min {
				dest = max
			}
			_, ok = cups[dest]
		}

		tmp := cups[dest]
		cups[dest] = pickUp[0]
		cups[pickUp[len(pickUp)-1]] = tmp
		current = cups[current]
	}

	result := 0
	if part == 1 {
		start := 1
		next := cups[start]
		for next != start {
			result = (result * 10) + next
			next = cups[next]
		}
	} else {
		result = cups[1] * cups[cups[1]]
	}

	return result
}

func parse(path string) ([]int, error) {
	file, err := os.Open(path)

	if err != nil {
		return []int{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret := []int{}
		s := scanner.Text()
		for _, c := range s {
			if '0' <= c && c <= '9' {
				ret = append(ret, int(c-'0'))
			} else {
				break
			}
		}

		if len(ret) > 0 {
			return ret, scanner.Err()
		}
	}

	return []int{}, scanner.Err()
}

func intSliceContains(a []int, n int) bool {
	for _, e := range a {
		if e == n {
			return true
		}
	}
	return false
}

func aMax(a []int) int {
	return reduce(func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}, a)
}

func aMin(a []int) int {
	return reduce(func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}, a)
}

func reduce(f func(int, int) int, a []int) int {
	ret := a[0]
	for _, e := range a[1:] {
		ret = f(ret, e)
	}
	return ret
}
