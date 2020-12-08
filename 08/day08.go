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
    ACC opcode = 0
    JMP opcode = 1
    NOP opcode = 2
)


type instruction struct {
    op opcode
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
    acc := 0
    ip := 0

    for {
        if ip >= len(instructions) {
            return acc, true
        }

        if visited[ip] {
            return acc, false
        }

        visited[ip] = true
        op := instructions[ip].op
        val := instructions[ip].val

        switch op {
        case ACC:
            acc += val
            fallthrough

        case NOP:
            ip++

        case JMP:
            ip += val
        }
    }
}


func solve2(instructions []instruction) int {
    nopjmp := []int{}

    for n, i := range instructions {
        if i.op == JMP || i.op == NOP {
            nopjmp = append(nopjmp, n)
        }
    }

    for _, i := range nopjmp {
        original := instructions[i]
        modified := original
        switch original.op {
        case NOP:
            modified.op = JMP
        case JMP:
            modified.op = NOP
        }

        instructions[i] = modified

        acc, ret := solve1(instructions)
        if ret {
            return acc
        } else {
            instructions[i] = original
        }
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
        op := NOP
        val := atoi(split[1])

        switch split[0] {
        case "acc":
            op = ACC

        case "jmp":
            op = JMP

        case "nop":
            op = NOP
        }

        instructions = append(instructions, instruction{
            op: op,
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
