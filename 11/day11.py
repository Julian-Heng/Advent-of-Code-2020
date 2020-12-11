#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = [list(map(str, i)) for i in f.read().strip().splitlines()]

    print(solve([i[:] for i in lines], adjacent_at, 4))
    print(solve([i[:] for i in lines], adjacent_all, 5))


def solve(lines, adj_callback, limit):
    after = lines
    before = None

    while before != after:
        before = [i[:] for i in after]
        for row in range(len(after)):
            for col in range(len(after[row])):
                seat = after[row][col]
                if seat == ".":
                    continue

                adj = adj_callback(before, row, col)
                if seat == "L" and not any("#" in i for i in adj):
                    after[row][col] = "#"
                if seat == "#" and grid_count(adj, "#") > limit:
                    after[row][col] = "L"

    return grid_count(after, "#")


def adjacent_at(grid, row, col):
    rmin = row - 1
    rmax = row + 1
    cmin = col - 1
    cmax = col + 1

    result = []
    for i in range(rmin, rmax + 1):
        r = []
        for j in range(cmin, cmax + 1):
            r.append(grid[i][j] if within(grid, i, j) else " ")
        result.append(r)

    return result


def adjacent_all(grid, row, col):
    d = [[(j, i) for j in range(-1, 2)] for i in range(-1, 2)]
    t = ("L", "#")
    return map(lambda i: map(lambda j: raycast(grid, col, row, j, t), i), d)


def raycast(grid, x, y, direction, target):
    if direction == (0, 0):
        return grid[y][x]

    cell = " "
    x, y = map(lambda a, b: a + b, (x, y), direction)
    while within(grid, y, x) and cell not in target:
        cell = grid[y][x]
        x, y = map(lambda a, b: a + b, (x, y), direction)
    return cell


def within(a, i, j):
    return 0 <= i < len(a) and 0 <= j < len(a[i])


def grid_count(grid, target):
    return sum(sum(j == target for j in i) for i in grid)


if __name__ == "__main__":
    main()
