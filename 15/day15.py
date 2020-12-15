#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        nums = list(map(int, f.read().split(",")))
    print(solve(nums[:], 2020))
    print(solve(nums[:], 30000000))


def solve(nums, target):
    spoken = {}
    num = nums.pop(0)
    last_num = -1
    for turn in range(1, target):
        if len(nums) > 0:
            spoken[num] = turn
            last_num = num
            num = nums.pop(0)
            continue

        if num not in spoken.keys():
            spoken[num] = turn
            last_num = num
            num = 0
        else:
            tmp = num
            last_num = num
            num = turn - spoken[num]
            spoken[tmp] = turn

    return num


if __name__ == "__main__":
    main()