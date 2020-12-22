#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = [i.splitlines() for i in f.read().strip().split("\n\n")]
    p1, p2 = parse(lines)
    print(solve1(p1[:], p2[:]))
    print(solve2(p1[:], p2[:]))


def solve1(p1, p2):
    while p1 and p2:
        a = p1.pop()
        b = p2.pop()
        if a > b:
            p1.insert(0, a)
            p1.insert(0, b)
        else:
            p2.insert(0, b)
            p2.insert(0, a)

    r = p1 if p1 else p2
    return sum(i * n for i, n in enumerate(r, start=1))


def solve2(p1, p2):
    p1, p2 = _solve2(p1, p2, set())
    r = p1 if p1 else p2
    return sum(i * n for i, n in enumerate(r, start=1))


def _solve2(p1, p2, seen):
    while p1 and p2:
        k = (tuple(p1), tuple(p2))
        if k in seen:
            return p1, []
        else:
            seen.add(k)

        a = p1.pop()
        b = p2.pop()

        if len(p1) >= a and len(p2) >= b:
            _p1, _p2 = _solve2(p1[-a:], p2[-b:], set())
            p1_won = _p1 and not _p2
        else:
            p1_won = a > b

        if p1_won:
            p1.insert(0, a)
            p1.insert(0, b)
        else:
            p2.insert(0, b)
            p2.insert(0, a)

    return p1, p2


def parse(lines):
    return map(lambda d: list(map(int, d[1:][::-1])), lines)


if __name__ == "__main__":
    main()
