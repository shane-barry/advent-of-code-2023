package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")

    if (err != nil) {
        panic(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    sum := 0
    for fileScanner.Scan() {
        // valid := true
        line := strings.Split(fileScanner.Text(), ":")
        // game := line[0]
        // id, _ := strconv.Atoi(strings.Split(strings.TrimSpace(game), " ")[1])
        sets := strings.Split(strings.TrimSpace(line[1]), ";")
        maxRed, maxGreen, maxBlue := 0, 0, 0
        for i := 0; i < len(sets); i++ {
            cubes := strings.Split(strings.TrimSpace(sets[i]), ",")
            for j := 0; j < len(cubes); j++ {
                count := strings.Split(strings.TrimSpace(cubes[j]), " ")[0]
                color := strings.Split(strings.TrimSpace(cubes[j]), " ")[1]
                c, _ := strconv.Atoi(count)
                // if !isValidCube(c, color) {
                //     valid = false
                // }
                switch color {
                case "blue":
                    maxBlue = max(maxBlue, c)
                case "red":
                    maxRed = max(maxRed, c)
                case "green":
                    maxGreen = max(maxGreen, c)
                }
            }
        }
        // if valid {
        //     sum += id
        // }
        sum += (maxRed * maxBlue * maxGreen)
    }

    fmt.Printf("%d\n", sum)
}

func isValidCube(count int, color string) bool {
    switch color {
    case "blue":
        return count <= 14
    case "red":
        return count <= 12
    case "green":
        return count <= 13
    }
    return false
}

func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b
}
