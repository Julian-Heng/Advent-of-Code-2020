#!/usr/bin/env python3

from functools import reduce


def main():
    with open("input", "r") as f:
        l = f.read().strip().split("\n\n")
        l = [[set(j) for j in i.splitlines()] for i in l]
    print(sum(map(len, (reduce(set.union, i) for i in l))))
    print(sum(map(len, (reduce(set.intersection, i) for i in l))))


if __name__ == "__main__":
    main()
