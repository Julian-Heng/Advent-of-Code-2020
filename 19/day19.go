package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)


type RuleType int
const (
    RULE_SUBRULES RuleType = 0
    RULE_ROOT_A   RuleType = 1
    RULE_ROOT_B   RuleType = 2
)


type Rule struct {
    SubRules [][]int
    Type RuleType
}


func main() {
    str_file, err := read("input")

    if err != nil {
        os.Exit(1)
    }

    rules1, rules2, lines:= parse(str_file)
    if rules1 == nil || rules2 == nil || lines == nil {
        os.Exit(1)
    }

    fmt.Println(solve(rules1, lines))
    fmt.Println(solve(rules2, lines))
}


func solve(rules map[int]Rule, lines []string) int {
    count := 0
    for _, l := range lines {
        ok, n := match(rules, 0, l)
        if ok && n == len(l) {
            count++
        }
    }
    return count
}


func match(rules map[int]Rule, rule int, s string) (bool, int) {
    if len(s) == 0 {
        return true, 0
    }

    rr := rules[rule]
    switch rr.Type {
    case RULE_ROOT_A: return s[0] == 'a', btoi(s[0] == 'a')
    case RULE_ROOT_B: return s[0] == 'b', btoi(s[0] == 'b')
    }

    valid := false
    offset := 0
    for _, r := range rr.SubRules {
        rule_valid := true
        offset = 0
        for i, sr := range r {
            result, n := match(rules, sr, s[offset:])
            offset += n
            if result && len(s[offset:]) == 0 
                if sr != rule && i != (len(r) - 1) {
                    result = false
                }
            }

            if ! result {
                rule_valid = false
                break
            }
        }

        if rule_valid {
            valid = true
            break
        }
    }

    return valid, offset
}


func parse(s string) (map[int]Rule, map[int]Rule, []string) {
    nl := getNewline(s)
    split := strings.SplitN(strings.TrimSpace(s), nl + nl, -1)
    rules1 := map[int]Rule{}
    rules2 := map[int]Rule{}
    lines := strings.SplitN(split[1], nl, -1)

    for _, l := range strings.SplitN(split[0], nl, -1) {
        ssplit := strings.SplitN(l, ": ", 2)
        rtype := RULE_SUBRULES

        rule, err := strconv.Atoi(ssplit[0])
        if err != nil {
            return nil, nil, nil
        }

        srules := [][]int{}
        for _, sr := range strings.SplitN(ssplit[1], " | ", -1) {
            ssrules := []int{}
            for _, r := range strings.SplitN(sr, " ", -1) {
                if r[0] == '"' {
                    switch r[1] {
                    case 'a': rtype = RULE_ROOT_A
                    case 'b': rtype = RULE_ROOT_B
                    }
                } else {
                    srule, err := strconv.Atoi(r)
                    if err != nil {
                        return nil, nil, nil
                    }
                    ssrules = append(ssrules, srule)
                }
            }
            srules = append(srules, ssrules)
        }

        rules1[rule] = Rule{SubRules: srules, Type: rtype}

        if rule == 8 {
            srules = [][]int{{42}, {42, 8}}
        } else if rule == 11 {
            srules = [][]int{{42, 31}, {42, 11, 31}}
        }
        rules2[rule] = Rule{SubRules: srules, Type: rtype}
    }

    return rules1, rules2, lines
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


func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}
