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
    ACTION_MASK  action = 0
    ACTION_WRITE action = 1
    ACTION_NOP   action = 2
)


type instruction struct {
    action action
    mask string
    address int
    value int
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
        case ACTION_MASK:
            mask = i.mask

        case ACTION_WRITE:
            switch version {
            case 1:
                mem[i.address] = apply_mask(mask, i.value)

            case 2:
                for _, v := range calculate_floats(mask, i.address) {
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


func apply_mask(mask string, value int) int {
    if len(mask) == 0 {
        return value
    }

    n := apply_mask(mask[:len(mask)-1], value >> 1)

    if mask[len(mask)-1] == 'X' {
        n <<= 1
        n ^= value & 1
    } else {
        n <<= 1
        n ^= atoi(mask[len(mask)-1:])
    }

    return n
}


func calculate_floats(mask string, address int) ([]int) {
    if len(mask) == 1 {
        if mask[0] == 'X' {
            return []int{0, 1}
        } else {
            m := atoi(mask[len(mask)-1:])
            if m == 0 {
                m = address & 1
            }
            return []int{m}
        }
    }

    a := calculate_floats(mask[:len(mask)-1], address >> 1)

    if mask[len(mask)-1] == 'X' {
        a_size := len(a)
        for i := 0; i < a_size; i++ {
            v := a[i] << 1
            a[i] = v ^ 0
            a = append(a, v ^ 1)
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

        action := ACTION_NOP
        mask := ""
        address := 0
        value := 0

        if strings.HasPrefix(line, "mask") {
            action = ACTION_MASK
            mask = line[len(line)-36:len(line)]
        } else {
            action = ACTION_WRITE
            vals := re.FindAllString(line, 2)
            address = atoi(vals[0])
            value = atoi(vals[1])
        }

        instructions = append(instructions, instruction{
            action: action,
            mask: mask,
            address: address,
            value: value,
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
