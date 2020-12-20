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


type Pixel int
const (
    PIXEL_OFF Pixel = 0
    PIXEL_ON  Pixel = 1
)


type TileData struct {
    Id    int
    Data  [][]Pixel
    North []int
    East  []int
    South []int
    West  []int
}


type Tile struct {
    Id   int
    Data []TileData
}


func main() {
    str_file, err := read("input")

    if err != nil {
        os.Exit(1)
    }

    tiles := parse(str_file)

    grid, result := solve1(tiles)
    fmt.Println(result)
    fmt.Println(solve2(grid))
}


func solve1(tiles map[int]Tile) ([][]TileData, int) {
    n := int(math.Sqrt(float64(len(tiles))))
    grid := make([][]TileData, n)
    for i := range grid {
        grid[i] = make([]TileData, n)
    }

    available := make(map[int]struct{})
    for _, t := range tiles {
        available[t.Id] = struct{}{}
    }

    _solve1(tiles, available, grid, 0, 0)

    result := grid[0][0].Id
    result *= grid[0][len(grid[0])-1].Id
    result *= grid[len(grid)-1][0].Id
    result *= grid[len(grid)-1][len(grid)-1].Id

    return grid, result
}


func _solve1(
    tiles map[int]Tile,
    available map[int]struct{},
    grid [][]TileData,
    x int,
    y int) bool {

    if len(available) == 0 {
        return true
    }

    if y >= len(grid) || x >= len(grid[0]) {
        return false
    }

    t_ids := []int{}
    for t_id := range available {
        t_ids = append(t_ids, t_id)
    }

    sort.Ints(t_ids)

    for _, t_id := range t_ids {
        tile := tiles[t_id]
        delete(available, t_id)

        for _, t := range tile.Data {
            if y > 0 && ! sliceEqual(grid[y-1][x].South, t.North) {
                continue
            }

            if x > 0 && ! sliceEqual(grid[y][x-1].East, t.West) {
                continue
            }

            grid[y][x] = t

            _x := x + 1
            _y := y

            if x == len(grid) - 1 {
                _x = 0
                _y = y + 1
            }

            if _solve1(tiles, available, grid, _x, _y) {
                return true
            }
        }

        available[t_id] = struct{}{}
    }

    return false
}


func solve2(_grid [][]TileData) int {
    tmp_grid := [][][][]Pixel{}
    grid := [][]Pixel{}

    for _, row := range _grid {
        tmp_row := [][][]Pixel{}
        for _, col := range row {
            tmp_row = append(tmp_row, remove_border(col.Data))
        }
        tmp_grid = append(tmp_grid, tmp_row)
    }

    for y := 0; y < len(tmp_grid); y++ {
        for x := 0; x < len(tmp_grid[y][0]); x++ {
            grid = append(grid, []Pixel{})
        }
    }

    n := len(tmp_grid[0][0])
    for y := range grid {
        for x := 0; x < (n * len(tmp_grid[0])); x++ {
            grid[y] = append(grid[y], tmp_grid[y / n][x / n][y % n][x % n])
        }
    }

    monsters, err := makeTileData([]string{
        "                  # ",
        "#    ##    ##    ###",
        " #  #  #  #  #  #   ",
    }, false)

    if err != nil {
        return 0
    }

    result := pixel_count(grid, PIXEL_ON)
    monster_size := pixel_count(monsters[0].Data, PIXEL_ON)

    for _, monster := range monsters {
        count := 0

        for y := 0; y < (len(grid) - len(monster.Data) + 1); y++ {
            for x := 0; x < (len(grid) - len(monster.Data[0]) + 1); x++ {
                if match(grid, x, y, monster.Data) {
                    count++
                }
            }
        }

        if count > 0 {
            result -= count * monster_size
            break
        }
    }

    return result
}


func match(grid [][]Pixel, x int, y int, monster [][]Pixel) bool {
    for py := 0; py < len(monster); py++ {
        for px := 0; px < len(monster[py]); px++ {
            if monster[py][px] == PIXEL_OFF {
                continue
            }
            if grid[y+py][x+px] != PIXEL_ON {
                return false
            }
        }
    }
    return true
}


func parse(s string) map[int]Tile {
    tiles := map[int]Tile{}
    nl := getNewline(s)

    for _, t := range strings.SplitN(strings.TrimSpace(s), nl + nl, -1) {
        tile, err := makeTile(t)
        if err == nil {
            tiles[tile.Id] = tile
        }
    }

    return tiles
}


func makeTile(t string) (Tile, error) {
    reg := regexp.MustCompile(`(\d+)`)
    nl := getNewline(t)
    split := strings.SplitN(strings.TrimSpace(t), nl, -1)

    id, err := strconv.Atoi(reg.FindString(split[0]))

    if err != nil {
        return Tile{}, err
    }

    data, err := makeTileData(split[1:], true)

    if err != nil {
        return Tile{}, errors.New("Data invalid")
    }

    for i, _ := range data {
        data[i].Id = id
    }

    return Tile{id, data}, nil
}


func makeTileData(lines []string, border_check bool) ([]TileData, error) {
    data := []TileData{}

    main := TileData{}
    for _, r := range lines {
        row := []Pixel{}
        for _, c := range r {
            switch c {
            case '#': row = append(row, PIXEL_ON)
            default: row = append(row, PIXEL_OFF)
            }
        }
        main.Data = append(main.Data, row)
    }

    data = append(data, main)

    // All rotations
    last := main.Data
    next := TileData{Data: rotate(last)}
    data = append(data, next)

    last = next.Data
    next = TileData{Data: rotate(last)}
    data = append(data, next)

    last = next.Data
    next = TileData{Data: rotate(last)}
    data = append(data, next)

    // All flipped rotations
    last = flip(main.Data)
    next = TileData{Data: last}
    data = append(data, next)

    last = next.Data
    next = TileData{Data: rotate(last)}
    data = append(data, next)

    last = next.Data
    next = TileData{Data: rotate(last)}
    data = append(data, next)

    last = next.Data
    next = TileData{Data: rotate(last)}
    data = append(data, next)

    if border_check {
        for i, _ := range data {
            calculate_border(&data[i])
        }
    }

    return data, nil
}


func calculate_border(tile *TileData) {
    north := tile.Data[0]
    east := []Pixel{}
    south := tile.Data[len(tile.Data)-1]
    west := []Pixel{}

    for _, row := range tile.Data {
        west = append(west, row[0])
        east = append(east, row[len(row)-1])
    }

    tile.North = indexOfPixels(north, PIXEL_ON)
    tile.East = indexOfPixels(east, PIXEL_ON)
    tile.South = indexOfPixels(south, PIXEL_ON)
    tile.West = indexOfPixels(west, PIXEL_ON)
}


func indexOfPixels(a []Pixel, c Pixel) []int {
    results := []int{}
    for i, p := range a {
        if p == c {
            results = append(results, i)
        }
    }
    return results
}


func rotate(px [][]Pixel) [][]Pixel {
    result := [][]Pixel{}
    for j := 0; j < len(px[0]); j++ {
        result = append(result, []Pixel{})
    }

    for i := len(px) - 1; i >= 0; i-- {
        for j := 0; j < len(px[i]); j++ {
            result[j] = append(result[j], px[i][j])
        }
    }
    return result
}


func flip(px [][]Pixel) [][]Pixel {
    result := [][]Pixel{}
    for j := 0; j < len(px[0]); j++ {
        result = append(result, []Pixel{})
    }

    for i := 0; i < len(px); i++ {
        for j := 0; j < len(px[i]); j++ {
            result[j] = append(result[j], px[i][j])
        }
    }
    return result
}


func remove_border(px [][]Pixel) [][]Pixel {
    result := [][]Pixel{}
    for i := 1; i < len(px) - 1; i++ {
        row := []Pixel{}
        for j := 1; j < len(px[i]) - 1; j++ {
            row = append(row, px[i][j])
        }
        result = append(result, row)
    }
    return result
}


func pixel_count(px [][]Pixel, c Pixel) int {
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


func sliceEqual(a []int, b[]int) bool {
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
