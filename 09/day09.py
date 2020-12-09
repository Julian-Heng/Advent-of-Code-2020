#!/usr/bin/env python3

from itertools import combinations


def main():
    with open("input", "r") as f:
        nums = list(map(int, f.read().strip().splitlines()))

    n = solve(nums)
    print(n)
    print(solve2(nums, n))


def solve(nums):
    i = 0
    j = 24
    while j + 1 < len(nums):
        n = nums[j+1]
        a = nums[i:j+1]
        if not any((a + b == n for a, b in combinations(a, 2))):
            return n
        i += 1
        j += 1


def solve2(nums, n):
    c = combinations(range(len(nums)), 2)
    a = next((nums[i:j+1] for i, j in c if sum(nums[i:j+1]) == n), 0)
    return min(a) + max(a)


if __name__ == "__main__":
    main()
