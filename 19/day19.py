#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = f.read().strip().split("\n\n")
    rules1, rules2, lines = parse(lines)
    print(solve(rules1, lines))
    print(solve(rules2, lines))


def parse(lines):
    _rules, lines = map(str.splitlines, lines)
    _rules = [i.replace("\"", "").split(": ", 2) for i in _rules]
    rules1 = dict()
    rules2 = dict()
    for k, v in _rules:
        rules1[k] = [i.split() for i in v.split(" | ")]

    rules2 = rules1.copy()
    rules2["8"] = [["42"], ["42", "8"]]
    rules2["11"] = [["42", "31"], ["42", "11", "31"]]

    return rules1, rules2, lines


def solve(rules, lines):
    def _solve(s):
        ok, n = match(rules, "0", s)
        return ok and n == len(s)
    return sum(map(_solve, lines))


def match(rules, rule, s):
    if len(s) == 0:
        return True, 0

    if rule in ("a", "b"):
        return s[0] == rule, int(s[0] == rule)

    valid = False
    for r in rules[rule]:
        rule_valid = True
        offset = 0
        for sr in r:
            result, n = match(rules, sr, s[offset:])
            offset += n
            if result and not s[offset:]:
                if sr != rule and sr != r[-1]:
                    result = False

            if not result:
                rule_valid = False
                break

        if rule_valid:
            valid = True
            break

    return valid, offset


if __name__ == "__main__":
    main()
