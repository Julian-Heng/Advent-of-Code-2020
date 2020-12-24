package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type instruction int

const (
	instructionUnknown    instruction = -1
	instructionDirectionN instruction = 0
	instructionDirectionE instruction = 1
	instructionDirectionS instruction = 2
	instructionDirectionW instruction = 3
	instructionRotateL    instruction = 4
	instructionRotateR    instruction = 5
	instructionForward    instruction = 6
)

type navigation struct {
	action    instruction
	magnitude int
}

func main() {
	navs, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve1(navs))
	fmt.Println(solve2(navs))
}

func solve1(navs []navigation) int {
	N := 0
	E := 0
	S := 0
	W := 0
	currentDir := instructionDirectionE

	for _, i := range navs {
		action := i.action
		magnitude := i.magnitude

		switch action {
		case instructionDirectionN:
			N += magnitude

		case instructionDirectionE:
			E += magnitude

		case instructionDirectionS:
			S += magnitude

		case instructionDirectionW:
			W += magnitude

		case instructionRotateL:
			for i := 0; i < (magnitude / 90); i++ {
				switch currentDir {
				case instructionDirectionN:
					currentDir = instructionDirectionW

				case instructionDirectionE:
					currentDir = instructionDirectionN

				case instructionDirectionS:
					currentDir = instructionDirectionE

				case instructionDirectionW:
					currentDir = instructionDirectionS
				}
			}

		case instructionRotateR:
			for i := 0; i < (magnitude / 90); i++ {
				switch currentDir {
				case instructionDirectionN:
					currentDir = instructionDirectionE

				case instructionDirectionE:
					currentDir = instructionDirectionS

				case instructionDirectionS:
					currentDir = instructionDirectionW

				case instructionDirectionW:
					currentDir = instructionDirectionN
				}
			}

		case instructionForward:
			switch currentDir {
			case instructionDirectionN:
				N += magnitude

			case instructionDirectionE:
				E += magnitude

			case instructionDirectionS:
				S += magnitude

			case instructionDirectionW:
				W += magnitude
			}
		}
	}

	return abs(N-S) + abs(E-W)
}

func solve2(navs []navigation) int {
	WN := 1
	WE := 10
	WS := 0
	WW := 0

	N := 0
	E := 0
	S := 0
	W := 0

	for _, i := range navs {
		action := i.action
		magnitude := i.magnitude

		switch action {
		case instructionDirectionN:
			WN += magnitude

		case instructionDirectionE:
			WE += magnitude

		case instructionDirectionS:
			WS += magnitude

		case instructionDirectionW:
			WW += magnitude

		case instructionRotateL:
			for i := 0; i < (magnitude / 90); i++ {
				tmp := WN
				WN = WE
				WE = WS
				WS = WW
				WW = tmp
			}

		case instructionRotateR:
			for i := 0; i < (magnitude / 90); i++ {
				tmp := WN
				WN = WW
				WW = WS
				WS = WE
				WE = tmp
			}

		case instructionForward:
			N += WN * magnitude
			E += WE * magnitude
			S += WS * magnitude
			W += WW * magnitude
		}
	}

	return abs(N-S) + abs(E-W)
}

func parse(path string) ([]navigation, error) {
	navs := []navigation{}
	file, err := os.Open(path)

	if err != nil {
		return navs, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		action := instructionUnknown
		magnitude, err := strconv.Atoi(line[1:])

		if err != nil {
			return navs, err
		}

		switch line[0] {
		case 'N':
			action = instructionDirectionN

		case 'E':
			action = instructionDirectionE

		case 'S':
			action = instructionDirectionS

		case 'W':
			action = instructionDirectionW

		case 'L':
			action = instructionRotateL

		case 'R':
			action = instructionRotateR

		case 'F':
			action = instructionForward
		}

		navs = append(navs, navigation{action, magnitude})
	}

	return navs, scanner.Err()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
