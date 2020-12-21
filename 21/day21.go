package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)


type StringSet map[string]struct{}
type Food struct {
    Ingredients StringSet
    Allergens   StringSet
}


func main() {
    lines, err := ReadLines("input")

    if err != nil {
        os.Exit(1)
    }

    foods, tally := parse(lines)

    p1, p2 := Solve(foods, tally)
    fmt.Println(p1)
    fmt.Println(p2)
}


func Solve(foods []Food, tally map[string]int) (int, string) {
    i_only := []StringSet{}
    a_only := []StringSet{}

    for _, food := range foods {
        i_only = append(i_only, food.Ingredients)
        a_only = append(a_only, food.Allergens)
    }

    i_all := StringSetReduce(StringSetUnion, i_only, nil)
    a_all := StringSetReduce(StringSetUnion, a_only, nil)

    p := []Food{}
    for a_a := range a_all {
        valid := []StringSet{}
        for _, food := range foods {
            if StringSetContains(food.Allergens, a_a) {
                valid = append(valid, food.Ingredients)
            }
        }

        p = append(p, Food{
            Ingredients: StringSetReduce(StringSetIntersect, valid, nil),
            Allergens: StringSetMake([]string{a_a}),
        })
    }

    p_i_only := []StringSet{}
    for _, food := range p {
        p_i_only = append(p_i_only, food.Ingredients)
    }

    i_no := StringSetReduce(StringSetDifference, p_i_only, i_all)

    p1 := 0
    for i := range i_no {
        p1 += tally[i]
    }

    mapping := map[string]string{}
    for StringSetAny(func(s StringSet) bool { return len(s) > 0 }, p_i_only) {
        for _, food := range p {
            if len(food.Ingredients) == 1 {
                a := StringSetFirst(food.Allergens)
                i := StringSetFirst(food.Ingredients)
                mapping[a] = i

                for j := range p {
                    p[j].Ingredients = StringSetDifference(
                        p[j].Ingredients,
                        food.Ingredients,
                    )
                }
            }
        }

        p_i_only = []StringSet{}
        for _, food := range p {
            p_i_only = append(p_i_only, food.Ingredients)
        }
    }

    keys := []string{}
    for k := range mapping {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    values := []string{}
    for _, k := range keys {
        values = append(values, mapping[k])
    }
    p2 := strings.Join(values, ",")

    return p1, p2
}


func ReadLines(path string) ([]string, error) {
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


func parse(lines []string) ([]Food, map[string]int) {
    foods := []Food{}
    tally := map[string]int{}

    for _, l := range lines {
        split := strings.SplitN(strings.Trim(l, ")"), " (contains ", 2)
        i := StringSetMake(strings.SplitN(split[0], " ", -1))
        a := StringSetMake(strings.SplitN(split[1], ", ", -1))
        foods = append(foods, Food{Ingredients: i, Allergens: a})

        for k := range i {
            tally[k]++
        }
    }

    return foods, tally
}


func StringSetMake(a []string) StringSet {
    s := StringSet{}
    for _, e := range a {
        s[e] = struct{}{}
    }
    return s
}


func StringSetContains(s StringSet, e string) bool {
    if _, ok := s[e]; ok {
        return true
    }
    return false
}


func StringSetUnion(s1 StringSet, s2 StringSet) StringSet {
    s := StringSet{}
    for e := range s1 {
        s[e] = struct{}{}
    }
    for e := range s2 {
        s[e] = struct{}{}
    }
    return s
}


func StringSetIntersect(s1 StringSet, s2 StringSet) StringSet {
    s := StringSet{}
    for e := range s1 {
        if StringSetContains(s2, e) {
            s[e] = struct{}{}
        }
    }
    return s
}


func StringSetDifference(s1 StringSet, s2 StringSet) StringSet {
    s := StringSet{}
    for e := range s1 {
        if ! StringSetContains(s2, e){
            s[e] = struct{}{}
        }
    }
    return s
}


func StringSetAny(f func(StringSet) bool, sa []StringSet) bool {
    for _, s := range sa {
        if f(s) {
            return true
        }
    }
    return false
}


func StringSetFirst(s StringSet) string {
    for e := range s {
        return e
    }
    return ""
}


func StringSetReduce(
    f func(StringSet, StringSet) StringSet,
    sa []StringSet,
    start StringSet) StringSet {

    result := start
    if result == nil {
        result = sa[0]
    }

    for _, s := range sa[1:] {
        result = f(result, s)
    }

    return result
}
