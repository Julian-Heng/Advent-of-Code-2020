package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type pixel int

const (
	pixelOff pixel = 0
	pixelOn  pixel = 1
)

type tileData struct {
	id    int
	data  [][]pixel
	north []int
	east  []int
	south []int
	west  []int
}

type tile struct {
	id   int
	data []tileData
}

func main() {
	strFile, err := read("input")

	if err != nil {
		os.Exit(1)
	}

	tiles := parse(strFile)

	grid, result := solve1(tiles)
	fmt.Println(result)
	fmt.Println(solve2(grid))
}

func solve1(tiles map[int]tile) ([][]tileData, int) {
	n := int(math.Sqrt(float64(len(tiles))))
	grid := make([][]tileData, n)
	for i := range grid {
		grid[i] = make([]tileData, n)
	}

	available := make(map[int]struct{})
	for _, t := range tiles {
		available[t.id] = struct{}{}
	}

	_solve1(tiles, available, grid, 0, 0)

	result := grid[0][0].id
	result *= grid[0][len(grid[0])-1].id
	result *= grid[len(grid)-1][0].id
	result *= grid[len(grid)-1][len(grid)-1].id

	return grid, result
}

func _solve1(
	tiles map[int]tile,
	available map[int]struct{},
	grid [][]tileData,
	x int,
	y int) bool {

	if len(available) == 0 {
		return true
	}

	if y >= len(grid) || x >= len(grid[0]) {
		return false
	}

	tileids := []int{}
	for tileid := range available {
		tileids = append(tileids, tileid)
	}

	sort.Ints(tileids)

	for _, tileid := range tileids {
		tile := tiles[tileid]
		delete(available, tileid)

		for _, t := range tile.data {
			if y > 0 && !sliceEqual(grid[y-1][x].south, t.north) {
				continue
			}

			if x > 0 && !sliceEqual(grid[y][x-1].east, t.west) {
				continue
			}

			grid[y][x] = t

			_x := x + 1
			_y := y

			if x == len(grid)-1 {
				_x = 0
				_y = y + 1
			}

			if _solve1(tiles, available, grid, _x, _y) {
				return true
			}
		}

		available[tileid] = struct{}{}
	}

	return false
}

func solve2(_grid [][]tileData) int {
	tmpGrid := [][][][]pixel{}
	grid := [][]pixel{}

	for _, row := range _grid {
		tmpRow := [][][]pixel{}
		for _, col := range row {
			tmpRow = append(tmpRow, removeBorder(col.data))
		}
		tmpGrid = append(tmpGrid, tmpRow)
	}

	for y := 0; y < len(tmpGrid); y++ {
		for x := 0; x < len(tmpGrid[y][0]); x++ {
			grid = append(grid, []pixel{})
		}
	}

	n := len(tmpGrid[0][0])
	for y := range grid {
		for x := 0; x < (n * len(tmpGrid[0])); x++ {
			grid[y] = append(grid[y], tmpGrid[y/n][x/n][y%n][x%n])
		}
	}

	monsters, err := makeTiledata([]string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}, false)

	if err != nil {
		return 0
	}

	result := pixelCount(grid, pixelOn)
	monsterSize := pixelCount(monsters[0].data, pixelOn)

	for _, monster := range monsters {
		count := 0

		for y := 0; y < (len(grid) - len(monster.data) + 1); y++ {
			for x := 0; x < (len(grid) - len(monster.data[0]) + 1); x++ {
				if match(grid, x, y, monster.data) {
					count++
				}
			}
		}

		if count > 0 {
			result -= count * monsterSize
			break
		}
	}

	return result
}

func match(grid [][]pixel, x int, y int, monster [][]pixel) bool {
	for py := 0; py < len(monster); py++ {
		for px := 0; px < len(monster[py]); px++ {
			if monster[py][px] == pixelOff {
				continue
			}
			if grid[y+py][x+px] != pixelOn {
				return false
			}
		}
	}
	return true
}

func parse(s string) map[int]tile {
	tiles := map[int]tile{}
	nl := getNewline(s)

	for _, t := range strings.SplitN(strings.TrimSpace(s), nl+nl, -1) {
		tile, err := makeTile(t)
		if err == nil {
			tiles[tile.id] = tile
		}
	}

	return tiles
}

func makeTile(t string) (tile, error) {
	reg := regexp.MustCompile(`(\d+)`)
	nl := getNewline(t)
	split := strings.SplitN(strings.TrimSpace(t), nl, -1)

	id, err := strconv.Atoi(reg.FindString(split[0]))

	if err != nil {
		return tile{}, err
	}

	data, err := makeTiledata(split[1:], true)

	if err != nil {
		return tile{}, errors.New("data invalid")
	}

	for i := range data {
		data[i].id = id
	}

	return tile{id, data}, nil
}

func makeTiledata(lines []string, borderCheck bool) ([]tileData, error) {
	data := []tileData{}

	main := tileData{}
	for _, r := range lines {
		row := []pixel{}
		for _, c := range r {
			switch c {
			case '#':
				row = append(row, pixelOn)

			default:
				row = append(row, pixelOff)
			}
		}
		main.data = append(main.data, row)
	}

	data = append(data, main)

	// All rotations
	last := main.data
	next := tileData{data: rotate(last)}
	data = append(data, next)

	last = next.data
	next = tileData{data: rotate(last)}
	data = append(data, next)

	last = next.data
	next = tileData{data: rotate(last)}
	data = append(data, next)

	// All flipped rotations
	last = flip(main.data)
	next = tileData{data: last}
	data = append(data, next)

	last = next.data
	next = tileData{data: rotate(last)}
	data = append(data, next)

	last = next.data
	next = tileData{data: rotate(last)}
	data = append(data, next)

	last = next.data
	next = tileData{data: rotate(last)}
	data = append(data, next)

	if borderCheck {
		for i := range data {
			calculateBorder(&data[i])
		}
	}

	return data, nil
}

func calculateBorder(tile *tileData) {
	north := tile.data[0]
	east := []pixel{}
	south := tile.data[len(tile.data)-1]
	west := []pixel{}

	for _, row := range tile.data {
		west = append(west, row[0])
		east = append(east, row[len(row)-1])
	}

	tile.north = indexOfPixels(north, pixelOn)
	tile.east = indexOfPixels(east, pixelOn)
	tile.south = indexOfPixels(south, pixelOn)
	tile.west = indexOfPixels(west, pixelOn)
}

func indexOfPixels(a []pixel, c pixel) []int {
	results := []int{}
	for i, p := range a {
		if p == c {
			results = append(results, i)
		}
	}
	return results
}

func rotate(px [][]pixel) [][]pixel {
	result := [][]pixel{}
	for j := 0; j < len(px[0]); j++ {
		result = append(result, []pixel{})
	}

	for i := len(px) - 1; i >= 0; i-- {
		for j := 0; j < len(px[i]); j++ {
			result[j] = append(result[j], px[i][j])
		}
	}
	return result
}

func flip(px [][]pixel) [][]pixel {
	result := [][]pixel{}
	for j := 0; j < len(px[0]); j++ {
		result = append(result, []pixel{})
	}

	for i := 0; i < len(px); i++ {
		for j := 0; j < len(px[i]); j++ {
			result[j] = append(result[j], px[i][j])
		}
	}
	return result
}

func removeBorder(px [][]pixel) [][]pixel {
	result := [][]pixel{}
	for i := 1; i < len(px)-1; i++ {
		row := []pixel{}
		for j := 1; j < len(px[i])-1; j++ {
			row = append(row, px[i][j])
		}
		result = append(result, row)
	}
	return result
}

func pixelCount(px [][]pixel, c pixel) int {
	result := 0
	for _, row := range px {
		for _, col := range row {
			if col == c {
				result++
			}
		}
	}
	return result
}

func read(path string) (string, error) {
	bfile, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	file := string(bfile)
	return file, err
}

func getNewline(s string) string {
	for _, v := range s {
		if v == '\r' {
			return "\r\n"
		}
	}
	return "\n"
}

func sliceEqual(a []int, b []int) bool {
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
