#!/usr/bin/env python3

import re
import sys

from itertools import product


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()
    black = solve1(lines)
    print(len(black))
    print(solve2(black))


def solve1(lines):
    tiles = set()
    reg = re.compile(r"((?:s|n)?(?:e|w))")
    for l in lines:
        c = 0
        r = 0
        for d in reg.findall(l):
            dc, dr = get_delta(d, c, r)
            c += dc
            r += dr

        coords = (c, r)
        if coords in tiles:
            tiles.remove(coords)
        else:
            tiles.add(coords)

    return tiles


def solve2(tiles, iterations=100):
    _min = sys.maxsize
    _max = -_min
    for a in zip(*tiles):
        _min = min(min(a) - 1, _min)
        _max = max(max(a) + 1, _max)

    for r in range(iterations):
        _tiles = list()
        for point in product(range(_min, _max + 1), repeat=2):
            adj = get_adjacent(point)
            num_active = len(tiles.intersection(adj))

            if (point in tiles and num_active in (2, 3)) or num_active == 2:
                _min = min(min(point) - 1, _min)
                _max = max(max(point) + 1, _max)
                _tiles.append(point)

        tiles = set(_tiles)
    return len(tiles)


def get_adjacent(point):
    c, r = point
    dirs = ("e", "se", "sw", "w", "nw", "ne")
    adj = [(c, r)]
    for d in dirs:
        dc, dr = get_delta(d, c, r)
        adj.append((c + dc, r + dr))
    return set(adj)


def get_delta(d, c, r):
    dc = 0
    dr = 0
    if d == "w":
        dc = -1
    elif d == "e":
        dc = 1
    else:
        offset = (r % 2) != 0
        if "s" in d:
            dr = 1
        elif "n" in d:
            dr = -1
        if "w" in d:
            dc = 0 if offset else -1
        elif "e" in d:
            dc = 1 if offset else 0

    return dc, dr

if __name__ == "__main__":
    main()
