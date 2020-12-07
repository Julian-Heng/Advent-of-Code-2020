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
    lines, err := readlines("input")

    if err != nil {
        os.Exit(1)
    }

    bags := make_bags(lines)
    fmt.Println(solve1(bags))
    fmt.Println(solve2(bags))
}


func solve1(bags map[string]map[string]int) int {
    n := 0
    for k := range bags {
        if traverse(bags, k) {
            n++
        }
    }
    return n
}


func solve2(bags map[string]map[string]int) int {
    return traverse2(bags, "shiny gold")
}


func traverse(bags map[string]map[string]int, bag string) bool {
    for k := range bags[bag] {
        if k == "shiny gold" {
            return true
        }
    }

    for k := range bags[bag] {
        if traverse(bags, k) {
            return true
        }
    }

    return false
}


func traverse2(bags map[string]map[string]int, bag string) int {
    n := 0
    for k := range bags[bag] {
        v := bags[bag][k]
        n += v + (v * traverse2(bags, k))
    }
    return n
}


func make_bags(lines []string) map[string]map[string]int {
    bags := map[string]map[string]int{}
    re := regexp.MustCompile(`(\d+)\s(\w+\s\w+)`)
    for _, l := range lines {
        split := strings.SplitN(l, " bags contain ", 2)
        match := re.FindAllStringSubmatch(split[1], -1)

        sub_bags := map[string]int{}

        if len(match) > 0 {
            for _, m := range match {
                sub_bags[m[2]] = atoi(m[1])
            }
        }

        bags[split[0]] = sub_bags
    }

    return bags
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


func atoi(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return n
}
