#!/usr/bin/env python3

import re


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()

    instructions = parse(lines)
    print(solve(instructions, version=1))
    print(solve(instructions, version=2))


def parse(lines):
    instructions = []
    for l in lines:
        if l.startswith("mask"):
            instructions.append(l.split(" = ", 2)[-1])
        elif l.startswith("mem"):
            instructions.append(tuple(map(int, re.findall(r"\d+", l))))

    return instructions


def solve(instructions, version=1):
    mem = {}
    mask = ("X" if version == 1 else "0") * 36

    for i in instructions:
        if "X" in i:
            mask = i
        else:
            if version == 1:
                mem[i[0]-1] = apply_mask(mask, i[1])
            elif version == 2:
                for j in calculate_floating(mask, i[0]):
                    mem[j-1] = i[1]

    return sum(mem.values())


def apply_mask(mask, val):
    result = 0
    for i, j in zip(mask, list(bits(val, pad=len(mask)))[::-1]):
        result = (result << 1) ^ (j if i == "X" else int(i))
    return result


def calculate_floating(mask, mem):
    results = []
    for i, j in zip(mask, list(bits(mem, pad=len(mask)))[::-1]):
        if i == "X":
            if not results:
                results.append(0)
                results.append(1)
            else:
                _results = []
                for r in results:
                    _results.append((r << 1) ^ 0)
                    _results.append((r << 1) ^ 1)
                results = _results
        else:
            m = int(i)
            a = m if m == 1 else j
            if not results:
                results.append(a)
            else:
                for r in range(len(results)):
                    v = results[r]
                    results[r] = (v << 1) ^ a

    return results


def bits(n, pad=None):
    if pad is None:
        pad = n.bit_length()

    count = 1
    while n or count <= pad:
        yield n & 1
        n >>= 1
        count += 1


if __name__ == "__main__":
    main()
