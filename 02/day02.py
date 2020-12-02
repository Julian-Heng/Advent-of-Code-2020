#!/usr/bin/env python3

import re


def main():
    with open("input", "r") as f:
        lines = f.read().strip().split("\n")

    spec = re.compile(r"(\d+)-(\d+)\s(\w):\s(.*)")

    passwords = list()
    for l in lines:
        m = spec.match(l)
        passwords.append(
            (int(m.group(1)), int(m.group(2)), m.group(3), m.group(4))
        )

    part1(passwords)
    part2(passwords)


def part1(passwords):
    print(sum(i[0] <= i[3].count(i[2]) <= i[1] for i in passwords))


def part2(passwords):
    def check(s, t, i, j):
        return (s[i] == t and s[j] != t) or (s[i] != t and s[j] == t)
    print(sum(check(i[3], i[2], i[0] - 1, i[1] - 1) for i in passwords))


if __name__ == "__main__":
    main()
