package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)


func main() {
    nums, err := parse("input")

    if err != nil {
        os.Exit(1)
    }

    fmt.Println(solve(nums, 2020))
    fmt.Println(solve(nums, 30000000))
}


func solve(nums []int, limit int) int {
    spoken := map[int]int{}
    num := 0

    for i := 0; i < limit - 1; i++ {
        if i < (len(nums) - 1) {
            num = nums[i]
            spoken[num] = i
            continue
        } else if i == (len(nums) - 1) {
            num = nums[len(nums)-1]
        }

        n, ok := spoken[num]
        if ok {
            tmp := num
            num = i - n
            spoken[tmp] = i
        } else {
            spoken[num] = i
            num = 0
        }
    }

    return num
}


func parse(path string) ([]int, error) {
    nums := []int{}
    file, err := os.Open(path)

    if err != nil {
        return nums, err
    }

    defer file.Close()

    lines, err := csv.NewReader(file).ReadAll()

    if err != nil {
        return nums, err
    }

    for _, l := range lines {
        for _, s := range l {
            n, err := strconv.Atoi(s)

            if err != nil {
                return nums, err
            }

            nums = append(nums, n)
        }
    }

    return nums, err
}
