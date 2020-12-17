#!/usr/bin/env python3

import sys

from itertools import product


def main():
    with open("input", "r") as f:
        coords = [(x, y) for y, i in enumerate(f.read().strip().splitlines())
                         for x, j in enumerate(i) if j == "#"]
    print(solve(coords, ndim=3))
    print(solve(coords, ndim=4))


def solve(start, ndim=3, iterations=6):
    coords = {i + ((0,) * (ndim - len(i))) for i in start}

    _min = sys.maxsize
    _max = -_min
    for a in zip(*coords):
        _min = min(min(a) - 1, _min)
        _max = max(max(a) + 1, _max)

    for r in range(iterations):
        _coords = list()
        for point in product(range(_min, _max + 1), repeat=ndim):
            adj = set(product(*[range(a - 1, a + 2) for a in point]))
            num_active = len(coords.intersection(adj))

            if (point in coords and num_active in (3, 4)) or num_active == 3:
                _min = min(min(point) - 1, _min)
                _max = max(max(point) + 1, _max)
                _coords.append(point)

        coords = set(_coords)

    return len(coords)


if __name__ == "__main__":
    main()
