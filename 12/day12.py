#!/usr/bin/env python3

from collections import deque


def main():
    with open("input", "r") as f:
        lines = [(l[0], int(l[1:])) for l in f.read().strip().splitlines()]

    print(solve1(lines))
    print(solve2(lines))


def solve1(lines):
    directions = ("N", "E", "S", "W")
    curr_direction = "E"
    curr_index = directions.index(curr_direction)

    dist = {i: 0 for i in directions}

    for d, m in lines:
        if d in ("N", "S", "E", "W"):
            dist[d] += m
        elif d in ("L", "R"):
            index_offset = (m // 90) * (-1 if d == "L" else 1)
            curr_index += index_offset
            curr_index %= len(directions)
            curr_direction = directions[curr_index]
        elif d == "F":
            dist[curr_direction] += m

    return abs(dist["N"] - dist["S"]) + abs(dist["E"] - dist["W"])


def solve2(lines):
    directions = ("N", "E", "S", "W")
    dist = {i: 0 for i in directions}
    waypoint= {i: 0 for i in directions}
    waypoint["N"] = 1
    waypoint["E"] = 10

    for d, m in lines:
        if d in directions:
            waypoint[d] += m
        elif d in ("L", "R"):
            vals = deque(waypoint.values())
            index_offset = (m // 90) * (-1 if d == "L" else 1)
            vals.rotate(index_offset)
            waypoint = {k: v for k, v in zip(waypoint.keys(), vals)}
        elif d == "F":
            for k, v in waypoint.items():
                dist[k] += m * v

    return abs(dist["N"] - dist["S"]) + abs(dist["E"] - dist["W"])


if __name__ == "__main__":
    main()
