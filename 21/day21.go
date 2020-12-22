package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type stringSet map[string]struct{}
type food struct {
	ingredients stringSet
	allergens   stringSet
}

func main() {
	lines, err := readLines("input")

	if err != nil {
		os.Exit(1)
	}

	foods, tally := parse(lines)

	p1, p2 := solve(foods, tally)
	fmt.Println(p1)
	fmt.Println(p2)
}

func solve(foods []food, tally map[string]int) (int, string) {
	iOnly := []stringSet{}
	aOnly := []stringSet{}

	for _, f := range foods {
		iOnly = append(iOnly, f.ingredients)
		aOnly = append(aOnly, f.allergens)
	}

	iAll := stringSetReduce(stringSetUnion, iOnly, nil)
	aAll := stringSetReduce(stringSetUnion, aOnly, nil)

	p := []food{}
	for _, a := range stringSetToSortedStringSlice(aAll) {
		valid := []stringSet{}
		for _, f := range foods {
			if stringSetContains(f.allergens, a) {
				valid = append(valid, f.ingredients)
			}
		}

		p = append(p, food{
			ingredients: stringSetReduce(stringSetIntersect, valid, nil),
			allergens:   stringSetMake([]string{a}),
		})
	}

	pIOnly := []stringSet{}
	for _, f := range p {
		pIOnly = append(pIOnly, f.ingredients)
	}

	iNo := stringSetReduce(stringSetDifference, pIOnly, iAll)

	p1 := 0
	for i := range iNo {
		p1 += tally[i]
	}

	mapping := map[string]string{}
	for stringSetAny(func(s stringSet) bool { return len(s) > 0 }, pIOnly) {
		for _, f := range p {
			if len(f.ingredients) == 1 {
				a := stringSetFirst(f.allergens)
				i := stringSetFirst(f.ingredients)
				mapping[a] = i

				for j := range p {
					p[j].ingredients = stringSetDifference(
						p[j].ingredients,
						f.ingredients,
					)
				}
			}
		}

		pIOnly = []stringSet{}
		for _, f := range p {
			pIOnly = append(pIOnly, f.ingredients)
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

func readLines(path string) ([]string, error) {
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

func parse(lines []string) ([]food, map[string]int) {
	foods := []food{}
	tally := map[string]int{}

	for _, l := range lines {
		split := strings.SplitN(strings.Trim(l, ")"), " (contains ", 2)
		i := stringSetMake(strings.SplitN(split[0], " ", -1))
		a := stringSetMake(strings.SplitN(split[1], ", ", -1))
		foods = append(foods, food{ingredients: i, allergens: a})

		for k := range i {
			tally[k]++
		}
	}

	return foods, tally
}

func stringSetMake(a []string) stringSet {
	s := stringSet{}
	for _, e := range a {
		s[e] = struct{}{}
	}
	return s
}

func stringSetToSortedStringSlice(s stringSet) []string {
	a := []string{}
	for k := range s {
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}

func stringSetContains(s stringSet, e string) bool {
	if _, ok := s[e]; ok {
		return true
	}
	return false
}

func stringSetUnion(s1 stringSet, s2 stringSet) stringSet {
	s := stringSet{}
	for e := range s1 {
		s[e] = struct{}{}
	}
	for e := range s2 {
		s[e] = struct{}{}
	}
	return s
}

func stringSetIntersect(s1 stringSet, s2 stringSet) stringSet {
	s := stringSet{}
	for e := range s1 {
		if stringSetContains(s2, e) {
			s[e] = struct{}{}
		}
	}
	return s
}

func stringSetDifference(s1 stringSet, s2 stringSet) stringSet {
	s := stringSet{}
	for e := range s1 {
		if !stringSetContains(s2, e) {
			s[e] = struct{}{}
		}
	}
	return s
}

func stringSetAny(f func(stringSet) bool, sa []stringSet) bool {
	for _, s := range sa {
		if f(s) {
			return true
		}
	}
	return false
}

func stringSetFirst(s stringSet) string {
	for e := range s {
		return e
	}
	return ""
}

func stringSetReduce(
	f func(stringSet, stringSet) stringSet,
	sa []stringSet,
	start stringSet) stringSet {

	result := start
	if result == nil {
		result = sa[0]
	}

	for _, s := range sa[1:] {
		result = f(result, s)
	}

	return result
}
