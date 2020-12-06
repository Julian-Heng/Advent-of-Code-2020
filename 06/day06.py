#!/usr/bin/env python3

import string

from collections import Counter
from functools import reduce


def main():
    with open("input", "r") as f:
        lines = f.read().strip().split("\n\n")

    print(solve_1(lines))
    print(solve_2(lines))


def solve_1(ln):
    l = string.ascii_lowercase
    return sum(map(lambda x: sum([i in x for i in l]), ln))


def solve_2(ln):
    c = lambda x, y: "".join(set(x).intersection(y))
    ln = [i.splitlines() for i in ln]
    return sum(len(i) for i in (reduce(c, i) for i in ln))


if __name__ == "__main__":
    main()
