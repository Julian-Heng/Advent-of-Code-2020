#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        l = f.read().strip().split("\n\n")
        l = [[set(j) for j in i.splitlines()] for i in l]
    print(sum(len(set.union(*i)) for i in l))
    print(sum(len(set.intersection(*i)) for i in l))


if __name__ == "__main__":
    main()
