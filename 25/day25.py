#!/usr/bin/env python3


def main():
    with open("input", "r") as f:
        card, door = map(int, f.read().strip().splitlines())
    print(solve(card, door))


def solve(card, door):
    card_loop = get_loop_size(card)
    door_loop = get_loop_size(door)
    card_encrypt = get_encryption_key(card, door_loop)
    door_encrypt = get_encryption_key(door, card_loop)
    return card_encrypt if card_encrypt == door_encrypt else -1


def get_loop_size(n):
    v = 1
    s = 7
    loop = 0
    while v != n:
        v = (v * s) % 20201227
        loop += 1
    return loop


def get_encryption_key(s, l):
    v = 1
    for i in range(l):
        v = (v * s) % 20201227
    return v


if __name__ == "__main__":
    main()
