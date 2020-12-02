package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)


type password struct {
    min int
    max int
    target string
    password string
}


func main() {
    lines, err := readlines("input")

    if err != nil {
        os.Exit(1)
    }

    r, err := regexp.Compile(`(\d+)-(\d+)\s(\w):\s(.*)`)

    if err != nil {
        os.Exit(1)
    }

    passwords := []password{}

    for _, l := range lines {
        m := r.FindStringSubmatch(l)
        min, err1 := strconv.Atoi(m[1])
        max, err2 := strconv.Atoi(m[2])

        if err1 != nil || err2 != nil {
            os.Exit(1)
        }

        password := password{min: min, max: max, target: m[3], password: m[4]}
        passwords = append(passwords, password)
    }

    fmt.Println(part1(passwords))
    fmt.Println(part2(passwords))
}


func part1(passwords []password) (int) {
    valid := 0
    for _, p:= range passwords {
        count := strings.Count(p.password, p.target)
        if p.min <= count && count <= p.max {
            valid++
        }
    }
    return valid
}


func part2(passwords []password) (int) {
    valid := 0
    for _, p:= range passwords {
        a := p.password[p.min-1:p.min]
        b := p.password[p.max-1:p.max]
        t := p.target
        if (a == t && b != t) || (b == t && a != t) {
            valid++
        }
    }
    return valid
}


func readlines(path string) ([]string, error) {
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
