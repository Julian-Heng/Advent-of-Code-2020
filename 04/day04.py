#!/usr/bin/env python3

import re


def main():
    with open("input", "r") as f:
        lines = f.read().strip().split("\n\n")

    print(sum(solve(i, False) for i in lines))
    print(sum(solve(i, True) for i in lines))


def solve(p, extended=False):
    # Part 1
    valid = set(("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"))
    passport = dict(i.split(":", 1) for i in p.split())
    if not set(passport.keys()).issuperset(valid):
        return False

    if not extended:
        return True

    # Part 2
    incm = {"cm": (150, 193), "in": (59, 76)}
    col = ("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
    checks = {
        "byr": lambda v: check_num(v, 1920, 2002),
        "iyr": lambda v: check_num(v, 2010, 2020),
        "eyr": lambda v: check_num(v, 2020, 2030),
        "hgt": lambda v: check_num(v[:-2], *incm.get(v[-2:], [None, None])),
        "hcl": lambda v: re.match(r"^#[0-9A-Za-z]{6}$", v),
        "ecl": lambda v: v in col,
        "pid": lambda v: re.match(r"^\d{9}$", v),
        "cid": lambda v: True,
    }

    return all(checks[k](v) for k, v in passport.items())


def check_num(n, l, h):
    return not (l is None or h is None) and l <= int(n) <= h


if __name__ == "__main__":
    main()
