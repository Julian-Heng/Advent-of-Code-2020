package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)


func main() {
    str_file, err := read("input")

    if err != nil {
        os.Exit(1)
    }

    delim := getNewline(str_file)
    lines := strings.Split(strings.TrimSpace(str_file), delim + delim)
    fmt.Println(solve1(lines))
    fmt.Println(solve2(lines))
}


func solve1(ln []string) int {
    result := 0

    for _, l := range ln {
        for i := 'a'; i <= 'z'; i++ {
            for _, j := range l {
                if i == j {
                    result++
                    break
                }
            }
        }
    }

    return result
}


func solve2(ln []string) int {
    result := 0

    delim := getNewline(ln[0])
    for _, l := range ln {
        split := strings.Split(l, delim)
        table := countLetters(split[0])

        for _, s := range split[1:] {
            table2 := countLetters(s)
            for j, _ := range table {
                table[j] = min(table[j], table2[j])
            }
        }

        for _, v := range table {
            if v > 0 {
                result++
            }
        }
    }

    return result
}


func countLetters(s string) []int {
    table := make([]int, 26)
    for _, v := range s {
        table[v-'a']++
    }
    return table
}


func min(a int, b int) int {
    if a <= b {
        return a
    }
    return b
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
