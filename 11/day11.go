package main

import (
	"bufio"
	"fmt"
	"os"
)

type seat int

// seat states
const (
	SeatEmpty    seat = 0
	SeatOccupied seat = 1
	SeatFloor    seat = 2
)

func main() {
	lines, err := readlines("input")

	if err != nil {
		os.Exit(1)
	}

	Seats, width, height := parse(lines)
	fmt.Println(solve(Seats, width, height, adjacentAt, 4))
	fmt.Println(solve(Seats, width, height, adjacentAll, 5))
}

func solve(
	Seats []seat,
	width int,
	height int,
	adjCallback func([]seat, int, int, int, int, func(seat) bool),
	limit int,
) int {
	after := make([]seat, len(Seats))
	before := make([]seat, len(Seats))
	copy(after, Seats)

	for !sameSeats(before, after) {
		copy(before, after)
		for i := range after {
			switch after[i] {
			case SeatEmpty:
				x, y := idxToCoords(i, width)
				check := true
				adjCallback(before, x, y, width, height, func(s seat) bool {
					if s == SeatOccupied {
						check = false
						return false
					}
					return true
				})

				if check {
					after[i] = SeatOccupied
				}

			case SeatOccupied:
				x, y := idxToCoords(i, width)
				count := 0
				adjCallback(before, x, y, width, height, func(s seat) bool {
					if s == SeatOccupied {
						count++
						if count >= limit {
							after[i] = SeatEmpty
							return false
						}
					}
					return true
				})
			}
		}
	}

	return countElements(after, SeatOccupied)
}

func adjacentAt(
	a []seat,
	_x int,
	_y int,
	w int,
	h int,
	callback func(seat) bool,
) {
	xmin := max(0, _x-1)
	xmax := min(w, _x+2)
	ymin := max(0, _y-1)
	ymax := min(h, _y+2)

	for y := ymin; y < ymax; y++ {
		for x := xmin; x < xmax; x++ {
			if x == _x && y == _y {
				continue
			}

			if !callback(a[coordsToIdx(x, y, w)]) {
				return
			}
		}
	}
}

func adjacentAll(
	a []seat,
	_x int,
	_y int,
	w int,
	h int,
	callback func(seat) bool,
) {
	directions := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	targets := []seat{SeatEmpty, SeatOccupied}

	for _, d := range directions {
		s := raycast(a, w, h, _x, _y, d[0], d[1], targets)
		if !callback(s) {
			return
		}
	}
}

func raycast(
	a []seat,
	w int,
	h int,
	x int,
	y int,
	dx int,
	dy int,
	t []seat,
) seat {
	x += dx
	y += dy
	c := SeatFloor
	for inBounds(x, y, w, h) && !contains(t, c) {
		c = a[coordsToIdx(x, y, w)]
		x += dx
		y += dy
	}

	return c
}

func readlines(path string) ([]string, error) {
	lines := []string{}
	file, err := os.Open(path)

	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func parse(lines []string) ([]seat, int, int) {
	Seats := []seat{}
	height := len(lines)
	width := len(lines[0])

	for _, r := range lines {
		for _, c := range r {
			seat := SeatFloor
			switch c {
			case 'L':
				seat = SeatEmpty

			case '#':
				seat = SeatOccupied

			case '.':
				seat = SeatFloor
			}
			Seats = append(Seats, seat)
		}
	}
	return Seats, width, height
}

func sameSeats(a []seat, b []seat) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func countElements(a []seat, target seat) int {
	result := 0
	for _, v := range a {
		if v == target {
			result++
		}
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func idxToCoords(i int, w int) (int, int) {
	return i % w, i / w
}

func coordsToIdx(x int, y int, w int) int {
	return x + (y * w)
}

func inBounds(x int, y int, w int, h int) bool {
	return 0 <= x && x < w && 0 <= y && y < h
}

func contains(a []seat, t seat) bool {
	for _, v := range a {
		if v == t {
			return true
		}
	}
	return false
}
