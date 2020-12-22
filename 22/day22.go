package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)


func main() {
    file, err := Read("input")

    if err != nil {
        os.Exit(1)
    }

    deck := Parse(file)
    p1 := deck[0]
    p2 := deck[1]

    fmt.Println(Solve1(p1, p2))
    fmt.Println(Solve2(p1, p2))
}


func Solve1(_p1 []int, _p2 []int) int {
    p1 := IntSliceCopy(_p1)
    p2 := IntSliceCopy(_p2)

    for len(p1) > 0 && len(p2) > 0 {
        a := p1[0]
        b := p2[0]
        p1 = p1[1:]
        p2 = p2[1:]

        if a > b {
            p1 = IntSliceInsert(p1, len(p1), a)
            p1 = IntSliceInsert(p1, len(p1), b)
        } else {
            p2 = IntSliceInsert(p2, len(p2), b)
            p2 = IntSliceInsert(p2, len(p2), a)
        }
    }

    r := p1
    if len(r) == 0 {
        r = p2
    }

    result := 0
    for i, e := range r {
        result += (len(r) - i) * e
    }

    return result
}


func Solve2(_p1 []int, _p2 []int) int {
    p1, p2 := solve2(_p1, _p2, [][2][]int{})

    r := p1
    if len(r) == 0 {
        r = p2
    }

    result := 0
    for i, e := range r {
        result += (len(r) - i) * e
    }

    return result
}


func solve2(
    _p1 []int,
    _p2 []int,
    seen [][2][]int,
) ([]int, []int) {

    p1 := IntSliceCopy(_p1)
    p2 := IntSliceCopy(_p2)

    for len(p1) > 0 && len(p2) > 0 {
        //fmt.Println(p1, p2)
        k := [2][]int{}
        k[0] = IntSliceCopy(p1)
        k[1] = IntSliceCopy(p2)

        found := false
        for _, s := range seen {
            if IntSliceEqual(s[0], k[0]) && IntSliceEqual(s[1], k[1]) {
                //fmt.Println(s, k)
                found = true
                break
            }
        }

        if found {
            return p1, []int{}
        } else {
            seen = append(seen, k)
        }

        a := p1[0]
        b := p2[0]
        p1 = p1[1:]
        p2 = p2[1:]

        p1_won := false
        if len(p1) >= a && len(p2) >= b {
            __p1, __p2 := solve2(p1[:a], p2[:b], [][2][]int{})
            p1_won = len(__p1) > 0 && len(__p2) == 0
        } else {
            p1_won = a > b
        }

        if p1_won {
            p1 = IntSliceInsert(p1, len(p1), a)
            p1 = IntSliceInsert(p1, len(p1), b)
        } else {
            p2 = IntSliceInsert(p2, len(p2), b)
            p2 = IntSliceInsert(p2, len(p2), a)
        }
    }

    return p1, p2
}


func IntSliceCopy(a []int) []int {
    c := make([]int, len(a))
    copy(c, a)
    return c
}


func IntSliceInsert(s []int, k int, vs ...int) []int {
    if n := len(s) + len(vs); n <= cap(s) {
        s2 := s[:n]
        copy(s2[k+len(vs):], s[k:])
        copy(s2[k:], vs)
        return s2
    } else {
        s2 := make([]int, n)
        copy(s2, s[:k])
        copy(s2[k:], vs)
        copy(s2[k+len(vs):], s[k:])
        return s2
    }
}


func IntSliceEqual(a1 []int, a2 []int) bool {
    if len(a1) != len(a2) {
        return false
    }

    for i, _ := range a1 {
        if a1[i] != a2[i] {
            return false
        }
    }

    return true
}


func Parse(s string) [][]int {
    nl := GetNewline(s)
    deck := [][]int{}

    split := strings.SplitN(strings.TrimSpace(s), nl + nl, -1)
    for _, p := range split {
        hand := []int{}
        lines := strings.SplitN(p, nl, -1)
        for _, l := range lines {
            if strings.Contains(l, "Player") {
                continue
            }

            n, err := strconv.Atoi(l)
            if err != nil {
                return nil
            }
            hand = append(hand, n)
        }
        deck = append(deck, hand)
    }

    return deck
}


func Read(path string) (string, error) {
    bfile, err := ioutil.ReadFile(path)

    if err != nil {
        return "", err
    }

    file := string(bfile)
    return file, err
}


func GetNewline(s string) string {
    for _, v := range s {
        if v == '\r' {
            return "\r\n"
        }
    }
    return "\n"
}
