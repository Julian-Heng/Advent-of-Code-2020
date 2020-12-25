package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	card, door, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve(card, door))
}

func solve(card int, door int) int {
	cardLoop := getLoopSize(card)
	doorLoop := getLoopSize(door)
	cardEncrypt := getEncryptionKey(card, doorLoop)
	doorEncrypt := getEncryptionKey(door, cardLoop)
	if cardEncrypt == doorEncrypt {
		return cardEncrypt
	}
	return -1
}

func getLoopSize(n int) int {
	v := 1
	s := 7
	loop := 0
	for loop = 0; v != n; loop++ {
		v = (v * s) % 20201227
	}
	return loop
}

func getEncryptionKey(s int, l int) int {
	v := 1
	for i := 0; i < l; i++ {
		v = (v * s) % 20201227
	}
	return v
}

func parse(path string) (int, int, error) {
	card := -1
	door := -1
	file, err := os.Open(path)

	if err != nil {
		return card, door, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return -1, -1, err
		}

		if card == -1 {
			card = n
		} else {
			door = n
		}
	}

	return card, door, err
}
