package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type instruction int

const (
	e       instruction = 0
	se      instruction = 1
	sw      instruction = 2
	w       instruction = 3
	nw      instruction = 4
	ne      instruction = 5
	unknown instruction = -1
)

type tile struct {
	col int
	row int
}

type tileSet map[tile]struct{}

func main() {
	instructions, err := parse("input")

	if err != nil {
		os.Exit(1)
	}

	black := solve1(instructions)
	fmt.Println(len(black))
	fmt.Println(solve2(black))
}

func solve1(instructions [][]instruction) tileSet {
	tiles := tileSet{}

	for _, instruction := range instructions {
		c := 0
		r := 0
		for _, ins := range instruction {
			dc, dr := getDelta(ins, c, r)
			c += dc
			r += dr
		}

		t := tile{col: c, row: r}
		if _, ok := tiles[t]; ok {
			delete(tiles, t)
		} else {
			tiles[t] = struct{}{}
		}
	}

	return tiles
}

func solve2(tiles tileSet) int {
	_min := int(^uint(0) >> 1)
	_max := -_min - 1
	for t := range tiles {
		_min = min(min(t.col, t.row)-1, _min)
		_max = max(max(t.col, t.row)+1, _max)
	}

	for i := 0; i < 100; i++ {
		_tiles := tileSet{}
		for r := _min; r <= _max; r++ {
			for c := _min; c <= _max; c++ {
				numActive := 0
				for _, a := range getAdjacent(c, r) {
					if _, ok := tiles[a]; ok {
						numActive++
					}
				}

				t := tile{col: c, row: r}
				_, ok := tiles[t]
				if (ok && numActive == 1) || numActive == 2 {
					_min = min(min(t.col, t.row)-1, _min)
					_max = max(max(t.col, t.row)+1, _max)
					_tiles[t] = struct{}{}
				}
			}
		}

		tiles = _tiles
	}

	return len(tiles)
}

func parse(path string) ([][]instruction, error) {
	instructions := [][]instruction{}
	file, err := os.Open(path)

	if err != nil {
		return instructions, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(`((?:s|n)?(?:e|w))`)
	for scanner.Scan() {
		row := []instruction{}
		for _, i := range reg.FindAllString(scanner.Text(), -1) {
			ins := unknown
			switch i {
			case "e":
				ins = e

			case "se":
				ins = se

			case "sw":
				ins = sw

			case "w":
				ins = w

			case "nw":
				ins = nw

			case "ne":
				ins = ne
			}
			row = append(row, ins)
		}
		instructions = append(instructions, row)
	}

	return instructions, scanner.Err()
}

func getAdjacent(c int, r int) []tile {
	adj := []tile{}
	instructions := []instruction{e, se, sw, w, nw, ne}
	for _, ins := range instructions {
		dc, dr := getDelta(ins, c, r)
		adj = append(adj, tile{col: c + dc, row: r + dr})
	}
	return adj
}

func getDelta(ins instruction, c int, r int) (int, int) {
	dc := 0
	dr := 0

	switch ins {
	case e:
		dc = 1

	case se:
		dr = 1
		if (r % 2) == 0 {
			dc = 0
		} else {
			dc = 1
		}

	case sw:
		dr = 1
		if (r % 2) == 0 {
			dc = -1
		} else {
			dc = 0
		}

	case w:
		dc = -1

	case nw:
		dr = -1
		if (r % 2) == 0 {
			dc = -1
		} else {
			dc = 0
		}

	case ne:
		dr = -1
		if (r % 2) == 0 {
			dc = 0
		} else {
			dc = 1
		}
	}

	return dc, dr
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
