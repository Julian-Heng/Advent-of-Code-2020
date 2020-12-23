#!/usr/bin/env python3

from collections import deque


def main():
    with open("input", "r") as f:
        num = f.read().strip()

    print(solve(num, limit=100, part=1))
    print(solve(num, limit=10_000_000, part=2))


def solve(num, limit=100, part=1):
    num = list(map(int, map(str, num)))

    _min = min(num)
    _max = max(num)
    current = num[0]

    if part == 2:
        num += range(_max + 1, 1_000_000 + 1)
        _min = min(num)
        _max = max(num)

    cups = {i: j for i, j in zip(num, num[1:] + [num[0]])}

    for r in range(limit):
        pick_up = list()
        _next = current
        for i in range(3):
            pick_up.append(cups[_next])
            _next = cups[_next]
        cups[current] = cups[_next]
        dest = current - 1
        while dest not in cups.keys() or dest in pick_up:
            dest -= 1
            if dest < _min:
                dest = _max

        tmp = cups[dest]
        cups[dest] = pick_up[0]
        cups[pick_up[-1]] = tmp
        current = cups[current]

    if part == 1:
        result = ""
        start = 1
        _next = cups[start]
        while _next != start:
            result += str(_next)
            _next = cups[_next]
        return result
    else:
        return cups[1] * cups[cups[1]]


if __name__ == "__main__":
    main()
