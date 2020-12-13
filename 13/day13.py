#!/usr/bin/env python3

try:
    from math import prod
except ImportError:
    from functools import reduce
    def prod(iter):
        return reduce(lambda x, y: x * y, iter)


def main():
    with open("input", "r") as f:
        lines = f.read().strip().splitlines()
    target = int(lines[0])
    buses = lines[1].split(",")
    buses = list(map(lambda x: int(x) if x.isnumeric() else None, buses))
    print(solve1(target, buses))
    print(solve2(buses))


def solve1(target, buses):
    times = map(lambda i: (i * round(target / i), i), (i for i in buses if i))
    times = ((i, j) for i, j in times if i > target)
    best_time, bus_id = min(times)
    return (best_time - target) * bus_id


def solve2(buses, start=0):
    return crt(*zip(*((i, i - n) for n, i in enumerate(buses) if i)))


def crt(n, a):
    _sum = 0
    _prod = prod(n)

    for n_i, a_i in zip(n, a):
        p = _prod // n_i
        _sum += a_i * mod_inv(p, n_i) * p

    return _sum % _prod


def mod_inv(a, b):
    b0 = b
    x0 = 0
    x1 = 1

    if b == 1:
        return 1

    while a > 1:
        q = a // b

        tmp = a
        a = b
        b = tmp %b

        tmp = x0
        x0 = x1 - q * x0
        x1 = tmp

    if x1 < 0:
        x1 += b0

    return x1


def solve2_old(buses, start=0):
    offsets = list((i, n) for n, i in enumerate(buses) if i)
    t = start
    while not all((t + i) % b == 0 for b, i in offsets[1:]):
        print(t)
        t += offsets[0][0]

    return t


if __name__ == "__main__":
    main()
