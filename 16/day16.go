package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    "strconv"
    "strings"
)


func main() {
    str_file, err := read("input")

    if err != nil {
        os.Exit(1)
    }

    rules, ticket, nearby := parse(str_file)
    fmt.Println(solve1(rules, nearby))
    fmt.Println(solve2(rules, ticket, nearby))
}


func solve1(rules map[string][][]int, nearby [][]int) int {
    sum := 0

    for _, i := range nearby {
        for _, j := range i {
            if ! inRules(rules, j) {
                sum += j
            }
        }
    }

    return sum
}


func solve2(rules map[string][][]int, ticket []int, nearby [][]int) int {
    valid := [][]int{}
    for _, i := range nearby {
        if ticketValid(rules, i) {
            valid = append(valid, i)
        }
    }

    choices := [][]string{}
    for i := 0; i < len(valid[0]); i++ {
        choice := []string{}
        for k, v := range rules {
            check := true
            for _, t := range valid {
                if ! inRule(v, t[i]) {
                    check = false
                    break
                }
            }

            if check {
                choice = append(choice, k)
            }
        }
        choices = append(choices, choice)
    }

    departure := []int{}
    for ! allEmpty(choices) {
        for i, c := range choices {
            if len(c) == 1 {
                choice := c[0]

                for j, c2 := range choices {
                    if _, ok := contains(c2, choice); ok {
                        choices[j] = remove(c2, choice)
                    }
                }

                if strings.HasPrefix(choice, "departure") {
                    departure = append(departure, i)
                }

                break
            }
        }
    }

    prod := 1
    for _, d := range departure {
        prod *= ticket[d]
    }
    return prod
}


func parse(s string) (map[string][][]int, []int, [][]int) {
    lines := [][]string{}
    delim := getNewline(s)

    for _, l := range strings.SplitN(strings.TrimSpace(s), delim + delim, -1) {
        lines = append(lines, strings.SplitN(l, delim, -1))
    }

    reg := regexp.MustCompile(`(\d+)-(\d+)`)
    rules := map[string][][]int{}
    for _, r := range lines[0] {
        split := strings.SplitN(r, ": ", 2)
        rules[split[0]] = [][]int{}
        matches := reg.FindAllStringSubmatch(split[1], -1)
        for _, m := range matches {
            ranges := []int{}
            low, err1 := strconv.Atoi(m[1])
            high, err2 := strconv.Atoi(m[2])
            if err1 != nil || err2 != nil {
                return nil, nil, nil
            }
            ranges = append(ranges, low)
            ranges = append(ranges, high)
            rules[split[0]] = append(rules[split[0]], ranges)
        }
    }

    ticket := []int{}
    for _, i := range strings.SplitN(lines[1][1], ",", -1) {
        n, err := strconv.Atoi(i)
        if err != nil {
            return nil, nil, nil
        }
        ticket = append(ticket, n)
    }

    nearby := [][]int{}
    for _, l := range lines[2][1:] {
        nearby_ticket := []int{}
        for _, i := range strings.SplitN(l, ",", -1) {
            n, err := strconv.Atoi(i)
            if err != nil {
                return nil, nil, nil
            }
            nearby_ticket = append(nearby_ticket, n)
        }
        nearby = append(nearby, nearby_ticket)
    }

    return rules, ticket, nearby
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


func ticketValid(r map[string][][]int, t []int) bool {
    for _, i := range t {
        if ! inRules(r, i) {
            return false
        }
    }
    return true
}


func inRules(r map[string][][]int, n int) bool {
    for _, v := range r {
        if inRule(v, n) {
            return true
        }
    }
    return false
}


func inRule(r [][]int, n int) bool {
    for _, v:= range r {
        if v[0] <= n && n <= v[1] {
            return true
        }
    }
    return false
}


func allEmpty(a [][]string) bool {
    for _, v := range a {
        if len(v) > 0 {
            return false
        }
    }
    return true
}


func contains(a []string, s string) (int, bool) {
    for i, v := range a {
        if v == s {
            return i, true
        }
    }
    return -1, false
}


func remove(s []string, r string) []string {
    for i, v := range s {
        if v == r {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}
