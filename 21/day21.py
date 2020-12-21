#!/usr/bin/env python3

from collections import Counter
from functools import reduce


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()
    foods, tally = parse(lines)
    print(solve(foods, tally, 1))
    print(solve(foods, tally, 2))


def solve(foods, tally, part):
    i_all = reduce(set.union, (i[0] for i in foods))
    a_all = reduce(set.union, (i[1] for i in foods))

    p = list()
    for a_a in a_all:
        valid = iter(i_f for i_f, a_f in foods if a_a in a_f)
        p.append((a_a, reduce(set.intersection, valid)))

    if part == 1:
        i_no = reduce(set.difference, (i[1] for i in p), i_all)
        return sum(tally[i] for i in i_no)
    else:
        mapping = list()
        while any(i[1] for i in p):
            for n, (a_p, i_p) in enumerate(p):
                if len(i_p) == 1:
                    mapping.append((a_p, next(iter(i_p))))
                    for m in range(len(p)):
                        p[m] = (p[m][0], p[m][1].difference(i_p))

        return ",".join([i[1] for i in sorted(mapping)])


def parse(lines):
    result = list()
    count = Counter()
    for l in lines:
        _ingredients, _allergens = l.rstrip(")").split(" (contains ", 2)
        ingredients = set(_ingredients.split())
        allergens = set(_allergens.split(", "))

        result.append((ingredients, allergens))
        count.update(ingredients)

    return result, count


if __name__ == "__main__":
    main()
