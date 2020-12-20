#!/usr/bin/env python3

import re

from math import isqrt
from itertools import chain


def main():
    with open("input", "r") as f:
        lines = f.read().strip().split("\n\n")

    tiles = parse(lines)

    grid, result = solve1(tiles)
    print(result)
    print(solve2(grid))


def parse(lines):
    tiles = dict()
    for t in lines:
        _id = int(re.search(r"\d+", t)[0])
        tile = tuple(map(lambda l: tuple(map(str, l)), t.splitlines()[1:]))
        tiles[_id] = all_tiles(tile)

    return tiles


def solve1(tiles):
    n = isqrt(len(tiles))
    _grid = [[0] * n for _ in range(n)]
    _solve(tiles, list(tiles.keys()), _grid, 0, 0)
    result = _grid[0][0][0]
    result *= _grid[0][-1][0]
    result *= _grid[-1][0][0] 
    result *= _grid[-1][-1][0]

    # Cleanup grid
    grid = list()
    for row in _grid:
        grid_row = list()
        for col in row:
            grid_row.append(col[1][0])
        grid.append(grid_row)

    return grid, result


def _solve(tiles, available, grid, x, y):
    if not available:
        return True

    for t_id in list(available):
        t_group = tiles[t_id]
        available.remove(t_id)
        for t in t_group:
            if y > 0 and grid[y-1][x][1][1]["S"] != t[1]["N"]:
                continue

            if x > 0 and grid[y][x-1][1][1]["E"] != t[1]["W"]:
                continue

            # Valid tile
            grid[y][x] = (t_id, t)

            _x = 0 if x == len(grid) - 1 else x + 1
            _y = y + 1 if x == len(grid) - 1 else y

            if _solve(tiles, available, grid, _x, _y):
                return True
        available.append(t_id)

    return False


def solve2(_grid):
    grid = list()

    for y, row in enumerate(_grid):
        grid_row = list()
        for x, col in enumerate(row):
            tile = list()
            for l in col[1:-1]:
                tile.append(l[1:-1])
            grid_row.append(tile)
        grid.append(grid_row)

    grid = [list(chain(*i)) for row in grid for i in zip(*row)]
    monster = [
        "                  # ",
        "#    ##    ##    ###",
        " #  #  #  #  #  #   ",
    ]

    result = "".join(chain(*grid)).count("#")
    monster_size = "".join(chain(*monster)).count("#")

    monsters = [i[0] for i in all_tiles([list(map(str, i)) for i in monster])]

    for monster in monsters:
        count = 0

        for y in range(len(grid) - len(monster) + 1):
            for x in range(len(grid[0]) - len(monster[0]) + 1):
                if match(grid, x, y, monster):
                    count += 1

        if count > 0:
            result -= count * monster_size
            break

    return result


def match(grid, x, y, monster):
    for py in range(len(monster)):
        for px in range(len(monster[0])):
            if monster[py][px] == " ":
                continue
            if grid[y+py][x+px] != "#":
                return False
    return True


def all_tiles(tile):
    # All rotations
    result = [(tile, edge_stat(tile))]
    for _ in range(3):
        transformed = list(zip(*reversed(result[-1][0])))
        result.append((transformed, edge_stat(transformed)))

    # All flipped rotations
    transformed = list(zip(*tile))
    result.append((transformed, edge_stat(transformed)))
    for _ in range(3):
        transformed = list(zip(*reversed(result[-1][0])))
        result.append((transformed, edge_stat(transformed)))

    return result


def edge_stat(tile):
    e = {
        "N": tile[0],
        "E": [j[-1] for j in tile],
        "S": tile[-1],
        "W": [j[0] for j in tile],
    }

    return {k: [i for i, e in enumerate(v) if e == "#"] for k, v in e.items()}


if __name__ == "__main__":
    main()
