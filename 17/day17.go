package main

import (
    "bufio"
    "fmt"
    "os"
)


func main() {
    lines, err := readlines("input")

    if err != nil {
        os.Exit(1)
    }

    coords := parse(lines)
    fmt.Println(solve3D(coords))
    fmt.Println(solve4D(coords))
}


func solve3D(start [][2]int) int {
    coords := make(map[[3]int]struct{})
    for _, p := range start {
        coords[[3]int{p[0], p[1], 0}] = struct{}{}
    }

    _min := int(^uint(0) >> 1)
    _max := -_min - 1
    for n := 0; n < 3; n++ {
        for p := range coords {
            _min = min(p[n] - 1, _min)
            _max = max(p[n] + 1, _max)
        }
    }

    for t := 0; t < 6; t++ {
        _coords := make(map[[3]int]struct{})
        for z := _min; z <= _max; z++ {
            for y := _min; y <= _max; y++ {
                for x := _min; x <= _max; x++ {
                    num_active := 0
                    for _, a := range getAdjacent3D(x, y, z) {
                        _, ok := coords[a]
                        if ok {
                            num_active++
                        }
                    }

                    p := [3]int{x, y, z}
                    _, ok := coords[p]
                    if (ok && num_active == 4) || num_active == 3 {
                        _min = min(amin(p[:]) - 1, _min)
                        _max = max(amax(p[:]) + 1, _max)
                        _coords[p] = struct{}{}
                    }
                }
            }
        }

        coords = _coords
    }

    return len(coords)
}


func solve4D(start [][2]int) int {
    coords := make(map[[4]int]struct{})
    for _, p := range start {
        coords[[4]int{p[0], p[1], 0, 0}] = struct{}{}
    }

    _min := int(^uint(0) >> 1)
    _max := -_min - 1
    for n := 0; n < 4; n++ {
        for p := range coords {
            _min = min(p[n] - 1, _min)
            _max = max(p[n] + 1, _max)
        }
    }

    for t := 0; t < 6; t++ {
        _coords := make(map[[4]int]struct{})
        for w := _min; w <= _max; w++ {
            for z := _min; z <= _max; z++ {
                for y := _min; y <= _max; y++ {
                    for x := _min; x <= _max; x++ {
                        num_active := 0
                        for _, a := range getAdjacent4D(x, y, z, w) {
                            _, ok := coords[a]
                            if ok {
                                num_active++
                            }
                        }

                        p := [4]int{x, y, z, w}
                        _, ok := coords[p]
                        if (ok && num_active == 4) || num_active == 3 {
                            _min = min(amin(p[:]) - 1, _min)
                            _max = max(amax(p[:]) + 1, _max)
                            _coords[p] = struct{}{}
                        }
                    }
                }
            }
        }

        coords = _coords
    }

    return len(coords)
}


func getAdjacent3D(_x int, _y int, _z int) [][3]int {
    result := [][3]int{}
    for z := _z - 1; z <= _z + 1; z++ {
        for y := _y - 1; y <= _y + 1; y++ {
            for x := _x - 1; x <= _x + 1; x++ {
                result = append(result, [3]int{x, y, z})
            }
        }
    }
    return result
}


func getAdjacent4D(_x int, _y int, _z int, _w int) [][4]int {
    result := [][4]int{}
    for w := _w - 1; w <= _w + 1; w++ {
        for z := _z - 1; z <= _z + 1; z++ {
            for y := _y - 1; y <= _y + 1; y++ {
                for x := _x - 1; x <= _x + 1; x++ {
                    result = append(result, [4]int{x, y, z, w})
                }
            }
        }
    }
    return result
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


func parse(str []string) [][2]int {
    points := [][2]int{}
    for y, row := range str {
        for x, col := range row {
            if col == '#' {
                points = append(points, [2]int{x, y})
            }
        }
    }
    return points
}


func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}


func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}


func amin(a []int) int {
    return reduce(a, min)
}


func amax(a []int) int {
    return reduce(a, max)
}


func reduce(a []int, callback func(int, int) int) int {
    result := a[0]
    for _, i := range a[1:] {
        result = callback(result, i)
    }
    return result
}
