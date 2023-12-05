package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    points := 0
    for fileScanner.Scan() {
        winCount := 0
        line := strings.Split(fileScanner.Text(), ":")
        nums := strings.Split(line[1], "|")
        for _, winner := range strings.Fields(nums[0]) {
            for _, num := range strings.Fields(nums[1]) {
                if winner == num {
                    winCount++
                }
            }
        }
        if (winCount > 0) {
            points += 1 << (winCount - 1)
        }
    }
    fmt.Println(points)
}
