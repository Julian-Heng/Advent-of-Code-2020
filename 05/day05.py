#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()
    seats = [solve(l) for l in lines]

    # Part 1
    seats.sort()
    print(seats[-1])

    # Part 2
    first = seats[0]
    for second in seats[1:]:
        if second - first != 1:
            break
        first = second
    print(first + 1)


def solve(s):
    fb, lr = s[:7], s[-3:]

    fb = [i == "B" for i in s[:7]]
    lr = [i == "R" for i in s[-3:]]

    row = traverse(fb, 0, 127)
    col = traverse(lr, 0, 7)
    return (row * 8) + col


def traverse(p, l, h):
    if not p:
        return l

    offset = ((h - l) // 2) + 1
    l2 = l + (offset if p[0] else 0)
    h2 = h - (offset if not p[0] else 0)
    return traverse(p[1:], l2, h2)


if __name__ == "__main__":
    main()
