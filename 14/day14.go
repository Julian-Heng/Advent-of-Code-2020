package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type action int

const (
	actionMask  action = 0
	actionWrite action = 1
	actionNop   action = 2
)

type instruction struct {
	action  action
	mask    string
	address int
	value   int
}

func main() {
	instructions, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve(instructions, 1))
	fmt.Println(solve(instructions, 2))
}

func solve(instructions []instruction, version int) int {
	mem := map[int]int{}
	mask := ""

	for _, i := range instructions {
		switch i.action {
		case actionMask:
			mask = i.mask

		case actionWrite:
			switch version {
			case 1:
				mem[i.address] = applymask(mask, i.value)

			case 2:
				for _, v := range calculateFloats(mask, i.address) {
					mem[v] = i.value
				}
			}
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	return sum
}

func applymask(mask string, value int) int {
	if len(mask) == 0 {
		return value
	}

	n := applymask(mask[:len(mask)-1], value>>1)

	if mask[len(mask)-1] == 'X' {
		n <<= 1
		n ^= value & 1
	} else {
		n <<= 1
		n ^= atoi(mask[len(mask)-1:])
	}

	return n
}

func calculateFloats(mask string, address int) []int {
	if len(mask) == 1 {
		if mask[0] == 'X' {
			return []int{0, 1}
		}

		m := atoi(mask[len(mask)-1:])
		if m == 0 {
			m = address & 1
		}
		return []int{m}
	}

	a := calculateFloats(mask[:len(mask)-1], address>>1)

	if mask[len(mask)-1] == 'X' {
		aLen := len(a)
		for i := 0; i < aLen; i++ {
			v := a[i] << 1
			a[i] = v ^ 0
			a = append(a, v^1)
		}
	} else {
		m := atoi(mask[len(mask)-1:])
		if m == 0 {
			m = address & 1
		}

		for i, v := range a {
			a[i] = (v << 1) ^ m
		}
	}

	return a
}

func parse(path string) ([]instruction, error) {
	instructions := []instruction{}
	file, err := os.Open(path)

	if err != nil {
		return instructions, err
	}

	defer file.Close()

	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		action := actionNop
		mask := ""
		address := 0
		value := 0

		if strings.HasPrefix(line, "mask") {
			action = actionMask
			mask = line[len(line)-36:]
		} else {
			action = actionWrite
			vals := re.FindAllString(line, 2)
			address = atoi(vals[0])
			value = atoi(vals[1])
		}

		instructions = append(instructions, instruction{
			action:  action,
			mask:    mask,
			address: address,
			value:   value,
		})
	}

	return instructions, scanner.Err()
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}
