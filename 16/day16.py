#!/usr/bin/env python3

import re

from functools import reduce
from itertools import chain

try:
    from math import prod
except ImportError:
    def prod(iter):
        return reduce(lambda x, y: x * y, iter)


def main():
    with open("input", "r") as f:
        lines = [i.split("\n") for i in f.read().strip().split("\n\n")]

    rules, ticket, nearby = parse(lines)

    print(solve1(rules, nearby))
    print(solve2(rules, ticket, nearby))


def solve1(rules, nearby):
    r = list(reduce(set.union, map(set, rules.values())))
    return sum(i for i in chain(*nearby) if i not in r)


def solve2(rules, ticket, nearby):
    r = list(reduce(set.union, map(set, rules.values())))
    valid = [i for i in nearby if all(j in r for j in i)]

    column_choices = list()
    for col in zip(*valid):
        choices = [k for k, v in rules.items() if all(i in v for i in col)]
        column_choices.append(choices)

    departure = list()
    while any(i for i in column_choices):
        for i, c in enumerate(column_choices):
            if len(c) == 1:
                choice = c[0]
                for j in column_choices:
                    if choice in j:
                        j.remove(choice)
                if choice.startswith("departure"):
                    departure.append(i)
                break

    return prod(ticket[i] for i in departure)


def parse(lines):
    _rules, _tickets, _nearby = lines

    _rules = [i.split(": ") for i in _rules]
    _tickets = _tickets[1:]
    _nearby = _nearby[1:]

    rules = {k: list() for k, _ in _rules}
    for field, value in _rules:
        for m in re.finditer(r"(\d+)-(\d+)", value):
            rules[field] += range(int(m.group(1)), int(m.group(2)) + 1)

    tickets = list(map(int, ",".join(_tickets).split(",")))
    nearby = list(map(lambda i: list(map(int, i.split(","))), _nearby))

    return rules, tickets, nearby


if __name__ == "__main__":
    main()
