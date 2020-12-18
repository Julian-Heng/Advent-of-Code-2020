#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        lines = [i.replace(" ", "") for i in f.read().strip().splitlines()]

    print(sum(solve(l, version=1) for l in lines))
    print(sum(solve(l, version=2) for l in lines))


def solve(eq, version=1):
    tokens = tokenize(eq)
    result = None

    actions = {
        "+": lambda a, b: a + b,
        "-": lambda a, b: a - b,
        "*": lambda a, b: a * b,
    }

    for i, t in enumerate(tokens):
        n = 0
        if any(i.isnumeric() for i in t):
            if t.startswith("(") and t.endswith(")"):
                n = solve(t[1:-1], version=version)
            else:
                n = int(t)

            if result is None:
                result = n
            else:
                result = action(result, n)
        else:
            if version == 2 and t == "*":
                result *= solve("".join(tokens[i+1:]), version=version)
                return result
            else:
                action = actions[t]

    return result


def tokenize(eq):
    tokens = list()
    token = ""
    state = 1
    level = 0

    for i, t in enumerate(eq):
        if t == "(":
            level += 1
            if state == 1:
                state = 0

                if token:
                    tokens.append(token)
                token = ""
            token += t
        elif t == ")":
            level -= 1
            token += t
            if level == 0:
                state = 1
                if token:
                    tokens.append(token)
                token = ""
        elif t in ("+", "-", "*") and state != 0:
            if token:
                tokens.append(token)
            token = ""
            tokens.append(t)
        else:
            token += t

    if token:
        tokens.append(token)

    return tokens


if __name__ == "__main__":
    main()
