package main

import (
    "bufio"
    "fmt"
    "os"
)


type seat int
const (
    SEAT_EMPTY      seat = 0
    SEAT_OCCUPIED   seat = 1
    SEAT_FLOOR      seat = 2
)


func main() {
    lines, err := readlines("input")

    if err != nil {
        os.Exit(1)
    }

    seats, width, height := parse(lines)
    fmt.Println(solve(seats, width, height, adjacent_at, 4))
    fmt.Println(solve(seats, width, height, adjacent_all, 5))
}


func solve(
    seats []seat,
    width int,
    height int,
    adj_callback func([]seat, int, int, int, int, func(seat) bool),
    limit int) int {

    after := make([]seat, len(seats))
    before := make([]seat, len(seats))
    copy(after, seats)

    for ! sameSeats(before, after) {
        copy(before, after)
        for i := range after {
            switch after[i] {
            case SEAT_EMPTY:
                x, y := idxToCoords(i, width)
                check := true
                adj_callback(before, x, y, width, height, func(s seat) bool {
                    if s == SEAT_OCCUPIED {
                        check = false
                        return false
                    }
                    return true
                })

                if check {
                    after[i] = SEAT_OCCUPIED
                }

            case SEAT_OCCUPIED:
                x, y := idxToCoords(i, width)
                count := 0
                adj_callback(before, x, y, width, height, func(s seat) bool {
                    if s == SEAT_OCCUPIED {
                        count++
                    }
                    return true
                })

                if count >= limit {
                    after[i] = SEAT_EMPTY
                }
            }
        }
    }

    return countElements(after, SEAT_OCCUPIED)
}


func adjacent_at(
    a []seat,
    _x int,
    _y int,
    w int,
    h int,
    callback func(seat) bool) {

    xmin := max(0, _x - 1)
    xmax := min(w, _x + 2)
    ymin := max(0, _y - 1)
    ymax := min(h, _y + 2)

    for y := ymin; y < ymax; y++ {
        for x := xmin; x < xmax; x++ {
            if x == _x && y == _y {
                continue
            }

            if ! callback(a[coordsToIdx(x, y, w)]) {
                return
            }
        }
    }
}


func adjacent_all(
    a []seat,
    _x int,
    _y int,
    w int,
    h int,
    callback func(seat) bool) {

    directions := [][]int{
        {-1, -1}, {0, -1}, {1, -1},
        {-1,  0},          {1,  0},
        {-1,  1}, {0,  1}, {1,  1},
    }

    targets := []seat{SEAT_EMPTY, SEAT_OCCUPIED}

    for _, d := range directions {
        s := raycast(a, w, h, _x, _y, d[0], d[1], targets)
        if ! callback(s) {
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
    t []seat) seat {

    x += dx
    y += dy
    c := SEAT_FLOOR
    for inBounds(x, y, w, h) && ! contains(t, c) {
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
    seats := []seat{}
    height := len(lines)
    width := len(lines[0])

    for _, r := range lines {
        for _, c := range r {
            seat := SEAT_FLOOR
            switch c {
            case 'L':
                seat = SEAT_EMPTY

            case '#':
                seat = SEAT_OCCUPIED

            case '.':
                seat = SEAT_FLOOR
            }
            seats = append(seats, seat)
        }
    }
    return seats, width, height
}


func sameSeats(a []seat, b[]seat) bool {
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
