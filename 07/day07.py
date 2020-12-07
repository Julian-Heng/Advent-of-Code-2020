#!/usr/bin/env python3

import itertools
import re

from functools import reduce


def main():
    with open("input", "r") as f:
        lines = f.read().splitlines()
    bags = dict(parse(i) for i in lines)
    print(solve1(bags))
    print(solve2(bags))


def solve1(bags):
    return sum(traverse(bags, k) for k in bags.keys())


def solve2(bags):
    return traverse2(bags, "shiny gold")


def traverse(bags, bag):
    if "shiny gold" in bags[bag].keys():
        return True

    for k, v in bags[bag].items():
        if traverse(bags, k):
            return True

    return False


def traverse2(bags, bag):
    return sum(v + (v * traverse2(bags, k)) for k, v in bags[bag].items())


def parse(s):
    main_colour, _sub_colours = s.split(" bags contain ", 2)
    reg = r"(\d+)\s(\w+\s\w+)"
    sub_colours = {col: int(n) for n, col in re.findall(reg, _sub_colours)}
    return main_colour, sub_colours


if __name__ == "__main__":
    main()
