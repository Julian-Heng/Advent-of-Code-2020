package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type opcode int

const (
	acc opcode = 0
	jmp opcode = 1
	nop opcode = 2
)

type instruction struct {
	op  opcode
	val int
}

func main() {
	instructions, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	ret, _ := solve1(instructions)
	fmt.Println(ret)
	fmt.Println(solve2(instructions))
}

func solve1(instructions []instruction) (int, bool) {
	visited := make([]bool, len(instructions))
	totalAcc := 0
	ip := 0

	for {
		if ip >= len(instructions) {
			return totalAcc, true
		}

		if visited[ip] {
			return totalAcc, false
		}

		visited[ip] = true
		op := instructions[ip].op
		val := instructions[ip].val

		switch op {
		case acc:
			totalAcc += val
			fallthrough

		case nop:
			ip++

		case jmp:
			ip += val
		}
	}
}

func solve2(instructions []instruction) int {
	nopjmp := []int{}

	for n, i := range instructions {
		if i.op == jmp || i.op == nop {
			nopjmp = append(nopjmp, n)
		}
	}

	for _, i := range nopjmp {
		original := instructions[i]
		modified := original
		switch original.op {
		case nop:
			modified.op = jmp

		case jmp:
			modified.op = nop
		}

		instructions[i] = modified

		if totalAcc, ret := solve1(instructions); ret {
			return totalAcc
		}

		instructions[i] = original
	}

	return 0
}

func parse(path string) ([]instruction, error) {
	instructions := []instruction{}

	file, err := os.Open(path)

	if err != nil {
		return instructions, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.SplitN(scanner.Text(), " ", 2)
		op := nop
		val := atoi(split[1])

		switch split[0] {
		case "acc":
			op = acc

		case "jmp":
			op = jmp

		case "nop":
			op = nop
		}

		instructions = append(instructions, instruction{
			op:  op,
			val: val,
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
