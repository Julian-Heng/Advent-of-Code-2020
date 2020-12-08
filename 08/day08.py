#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()
    print(solve1(lines)[0])
    print(solve2(lines))


def solve1(lines):
    visited = [False] * len(lines)
    acc = 0
    ip = 0

    while True:
        if ip >= len(lines):
            return acc, True

        if visited[ip]:
            return acc, False

        visited[ip] = True
        op, val = lines[ip].split(" ", 2)
        val = int(val)

        if op == "acc":
            acc += val
            ip += 1
        elif op == "jmp":
            ip += val
        elif op == "nop":
            ip += 1

    return acc


def solve2(lines):
    nopjmp = [n for n, i in enumerate(lines) if "nop" in i or "jmp" in i]
    for i in nopjmp:
        line = lines[i]
        replace = ["nop", "jmp"] if "nop" in line else ["jmp", "nop"]
        lines[i] = line.replace(*replace)
        acc, ret = solve1(lines)
        if ret:
            return acc
        else:
            lines[i] = line


if __name__ == "__main__":
    main()
