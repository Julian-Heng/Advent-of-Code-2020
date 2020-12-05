#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()
    s = list(map(solve, lines))

    # Part 1
    s.sort()
    print(s[-1])

    # Part 2
    d = [j - i for i, j in zip(s[:-1], s[1:])]
    print(s[d.index(next(i for i in d if i != 1))] + 1)


def solve(s):
    return int("".join([str(int(i in ["B", "R"])) for i in s]), 2)


if __name__ == "__main__":
    main()
