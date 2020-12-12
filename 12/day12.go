package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)


type instruction int
const (
    INSTRUCTION_UNKNOWN     instruction = -1
    INSTRUCTION_DIRECTION_N instruction = 0
    INSTRUCTION_DIRECTION_E instruction = 1
    INSTRUCTION_DIRECTION_S instruction = 2
    INSTRUCTION_DIRECTION_W instruction = 3
    INSTRUCTION_ROTATE_L    instruction = 4
    INSTRUCTION_ROTATE_R    instruction = 5
    INSTRUCTION_FORWARD     instruction = 6
)


type navigation struct {
    action instruction
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
    currentDir := INSTRUCTION_DIRECTION_E

    for _, i := range navs {
        action := i.action
        magnitude := i.magnitude

        switch action {
        case INSTRUCTION_DIRECTION_N: N += magnitude
        case INSTRUCTION_DIRECTION_E: E += magnitude
        case INSTRUCTION_DIRECTION_S: S += magnitude
        case INSTRUCTION_DIRECTION_W: W += magnitude
        case INSTRUCTION_ROTATE_L:
            for i := 0; i < (magnitude / 90); i++ {
                switch currentDir {
                case INSTRUCTION_DIRECTION_N:
                    currentDir = INSTRUCTION_DIRECTION_W
                case INSTRUCTION_DIRECTION_E:
                    currentDir = INSTRUCTION_DIRECTION_N
                case INSTRUCTION_DIRECTION_S:
                    currentDir = INSTRUCTION_DIRECTION_E
                case INSTRUCTION_DIRECTION_W:
                    currentDir = INSTRUCTION_DIRECTION_S
                }
            }
        case INSTRUCTION_ROTATE_R:
            for i := 0; i < (magnitude / 90); i++ {
                switch currentDir {
                case INSTRUCTION_DIRECTION_N:
                    currentDir = INSTRUCTION_DIRECTION_E
                case INSTRUCTION_DIRECTION_E:
                    currentDir = INSTRUCTION_DIRECTION_S
                case INSTRUCTION_DIRECTION_S:
                    currentDir = INSTRUCTION_DIRECTION_W
                case INSTRUCTION_DIRECTION_W:
                    currentDir = INSTRUCTION_DIRECTION_N
                }
            }
        case INSTRUCTION_FORWARD:
            switch currentDir {
            case INSTRUCTION_DIRECTION_N: N += magnitude
            case INSTRUCTION_DIRECTION_E: E += magnitude
            case INSTRUCTION_DIRECTION_S: S += magnitude
            case INSTRUCTION_DIRECTION_W: W += magnitude
            }
        }
    }

    return abs(N - S) + abs(E - W)
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
        case INSTRUCTION_DIRECTION_N: WN += magnitude
        case INSTRUCTION_DIRECTION_E: WE += magnitude
        case INSTRUCTION_DIRECTION_S: WS += magnitude
        case INSTRUCTION_DIRECTION_W: WW += magnitude
        case INSTRUCTION_ROTATE_L:
            for i := 0; i < (magnitude / 90); i++ {
                tmp := WN
                WN = WE
                WE = WS
                WS = WW
                WW = tmp
            }
        case INSTRUCTION_ROTATE_R:
            for i := 0; i < (magnitude / 90); i++ {
                tmp := WN
                WN = WW
                WW = WS
                WS = WE
                WE = tmp
            }
        case INSTRUCTION_FORWARD:
            N += WN * magnitude
            E += WE * magnitude
            S += WS * magnitude
            W += WW * magnitude
        }
    }

    return abs(N - S) + abs(E - W)
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

        action := INSTRUCTION_UNKNOWN
        magnitude, err := strconv.Atoi(line[1:])

        if err != nil {
            return navs, err
        }

        switch line[0] {
        case 'N': action = INSTRUCTION_DIRECTION_N
        case 'E': action = INSTRUCTION_DIRECTION_E
        case 'S': action = INSTRUCTION_DIRECTION_S
        case 'W': action = INSTRUCTION_DIRECTION_W
        case 'L': action = INSTRUCTION_ROTATE_L
        case 'R': action = INSTRUCTION_ROTATE_R
        case 'F': action = INSTRUCTION_FORWARD
        }

        navs = append(navs, navigation{action, magnitude})
    }

    return navs, scanner.Err()
}


func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}
