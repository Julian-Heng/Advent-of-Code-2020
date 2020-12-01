#!/usr/bin/env python3

from itertools import combinations as combine
from math import prod


def main():
    with open("input", "r") as f:
        nums = list(map(int, f.read().split()))
    print(prod(next(i for i in combine(nums, 2) if sum(i) == 2020)))
    print(prod(next(i for i in combine(nums, 3) if sum(i) == 2020)))


if __name__ == "__main__":
    main()
