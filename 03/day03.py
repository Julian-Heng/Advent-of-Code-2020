#!/usr/bin/env python3

try:
    from math import prod
except ImportError:
    from functools import reduce
    def prod(iter):
        return reduce(lambda x, y: x * y, iter)


def main():
    with open("input", "r") as f:
        lines = f.read().strip().split("\n")

    spec = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]

    print(solve(lines, 3, 1))
    print(prod([solve(lines, x, y) for x, y in spec]))


def solve(lines, r_offset, d_offset):
    r = 0
    d = 0
    trees = 0
    while d < len(lines):
        trees += 1 if lines[d][r] == "#" else 0
        r = (r + r_offset) % len(lines[d])
        d += d_offset
    return trees


if __name__ == "__main__":
    main()
