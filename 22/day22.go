package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := read("input")

	if err != nil {
		os.Exit(1)
	}

	deck := parse(file)
	p1 := deck[0]
	p2 := deck[1]

	fmt.Println(solve1(p1, p2))
	fmt.Println(solve2(p1, p2))
}

func solve1(_p1 []int, _p2 []int) int {
	p1 := intSliceCopy(_p1)
	p2 := intSliceCopy(_p2)

	for len(p1) > 0 && len(p2) > 0 {
		a := p1[0]
		b := p2[0]
		p1 = p1[1:]
		p2 = p2[1:]

		if a > b {
			p1 = intSliceInsert(p1, len(p1), a)
			p1 = intSliceInsert(p1, len(p1), b)
		} else {
			p2 = intSliceInsert(p2, len(p2), b)
			p2 = intSliceInsert(p2, len(p2), a)
		}
	}

	r := p1
	if len(r) == 0 {
		r = p2
	}

	result := 0
	for i, e := range r {
		result += (len(r) - i) * e
	}

	return result
}

func solve2(_p1 []int, _p2 []int) int {
	p1, p2 := _solve2(_p1, _p2, [][2][]int{})

	r := p1
	if len(r) == 0 {
		r = p2
	}

	result := 0
	for i, e := range r {
		result += (len(r) - i) * e
	}

	return result
}

func _solve2(
	_p1 []int,
	_p2 []int,
	seen [][2][]int,
) ([]int, []int) {

	p1 := intSliceCopy(_p1)
	p2 := intSliceCopy(_p2)

	for len(p1) > 0 && len(p2) > 0 {
		k := [2][]int{}
		k[0] = intSliceCopy(p1)
		k[1] = intSliceCopy(p2)

		found := false
		for _, s := range seen {
			if intSliceEqual(s[0], k[0]) && intSliceEqual(s[1], k[1]) {
				found = true
				break
			}
		}

		if found {
			return p1, []int{}
		}
		seen = append(seen, k)

		a := p1[0]
		b := p2[0]
		p1 = p1[1:]
		p2 = p2[1:]

		p1Won := false
		if len(p1) >= a && len(p2) >= b {
			subP1, subP2 := _solve2(p1[:a], p2[:b], [][2][]int{})
			p1Won = len(subP1) > 0 && len(subP2) == 0
		} else {
			p1Won = a > b
		}

		if p1Won {
			p1 = intSliceInsert(p1, len(p1), a)
			p1 = intSliceInsert(p1, len(p1), b)
		} else {
			p2 = intSliceInsert(p2, len(p2), b)
			p2 = intSliceInsert(p2, len(p2), a)
		}
	}

	return p1, p2
}

func intSliceCopy(a []int) []int {
	c := make([]int, len(a))
	copy(c, a)
	return c
}

func intSliceInsert(s []int, k int, vs ...int) []int {
	n := len(s) + len(vs)

	if n <= cap(s) {
		s2 := s[:n]
		copy(s2[k+len(vs):], s[k:])
		copy(s2[k:], vs)
		return s2
	}

	s2 := make([]int, n)
	copy(s2, s[:k])
	copy(s2[k:], vs)
	copy(s2[k+len(vs):], s[k:])
	return s2
}

func intSliceEqual(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}

func parse(s string) [][]int {
	nl := getNewline(s)
	deck := [][]int{}

	split := strings.SplitN(strings.TrimSpace(s), nl+nl, -1)
	for _, p := range split {
		hand := []int{}
		lines := strings.SplitN(p, nl, -1)
		for _, l := range lines {
			if strings.Contains(l, "Player") {
				continue
			}

			n, err := strconv.Atoi(l)
			if err != nil {
				return nil
			}
			hand = append(hand, n)
		}
		deck = append(deck, hand)
	}

	return deck
}

func read(path string) (string, error) {
	bfile, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	file := string(bfile)
	return file, err
}

func getNewline(s string) string {
	for _, v := range s {
		if v == '\r' {
			return "\r\n"
		}
	}
	return "\n"
}
