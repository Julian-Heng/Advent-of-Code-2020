#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        nums = list(map(int, f.read().split(",")))
    print(solve(nums, 2020))
    print(solve(nums, 30000000))


def solve(nums, target):
    spoken = {}
    num = 0
    for turn in range(0, target - 1):
        if turn < len(nums) - 1:
            num = nums[turn]
            spoken[num] = turn
            continue
        elif turn == len(nums) - 1:
            num = nums[-1]

        if num not in spoken.keys():
            spoken[num] = turn
            num = 0
        else:
            tmp = num
            num = turn - spoken[num]
            spoken[tmp] = turn

    return num


if __name__ == "__main__":
    main()
