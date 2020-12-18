package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


const (
    TOKENIZER_STATE_INSIDE  int = 0
    TOKENIZER_STATE_OUTSIDE int = 1
)


func main() {
    lines, err := readlines("input")

    if err != nil {
        os.Exit(1)
    }

    tokenized := [][]string{}
    for _, l := range lines {
        tokenized = append(tokenized, tokenize(strings.ReplaceAll(l, " ", "")))
    }

    for v := 1; v <= 2; v++ {
        sum := 0
        for _, t := range tokenized {
            sum += solve(t, v)
        }
        fmt.Println(sum)
    }
}


func solve(tokens []string, version int) int {
    first := true
    result := 0
    op := ""

    for i, t := range tokens {
        if any(t, func(i rune) bool { return '0' <= i && i <= '9' }) {
            n := 0
            if strings.HasPrefix(t, "(") && strings.HasSuffix(t, ")") {
                n = solve(tokenize(t[1:len(t)-1]), version)
            } else {
                _n, err := strconv.Atoi(t)
                if err != nil {
                    return 0
                }
                n = _n
            }

            if first {
                result = n
                first = false
            } else {
                switch op {
                case "+": result += n
                case "-": result -= n
                case "*": result *= n
                }
            }
        } else {
            if version == 2 && t == "*" {
                result *= solve(tokens[i+1:], version)
                return result
            } else {
                op = t
            }
        }
    }

    return result
}


func tokenize(eq string) []string {
    tokens := []string{}
    token := ""
    state := TOKENIZER_STATE_OUTSIDE
    level := 0

    for _, _t := range eq {
        t := string(_t)
        switch _t {
        case '(':
            level++
            if state == TOKENIZER_STATE_OUTSIDE {
                state = TOKENIZER_STATE_INSIDE
                if len(token) > 0 {
                    tokens = append(tokens, token)
                }
                token = ""
            }
            token += t

        case ')':
            level--
            token += t
            if level == 0 {
                state = TOKENIZER_STATE_OUTSIDE
                if len(token) > 0 {
                    tokens = append(tokens, token)
                }
                token = ""
            }

        case '+', '-', '*':
            if state != TOKENIZER_STATE_INSIDE {
                if len(token) > 0 {
                    tokens = append(tokens, token)
                }
                token = ""
                tokens = append(tokens, t)
            } else {
                token += t
            }

        default:
            token += t
        }
    }

    if len(token) > 0 {
        tokens = append(tokens, token)
    }

    return tokens
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


func any(s string, callback func(rune) bool) bool {
    for _, v := range s {
        if callback(v) {
            return true
        }
    }
    return false
}
