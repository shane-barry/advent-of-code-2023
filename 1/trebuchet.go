package main

import (
    "bufio"
	"fmt"
    "os"
    "slices"
    "strconv"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
    file, err := os.Open("input.txt")

    if (err != nil) {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    sum := 0
    for fileScanner.Scan() {
        v := fileScanner.Text()
        first := false
        f := 0
        l := 0
        for i := 0; i < len(v); i++ {
            if v[i] >= '0' && v[i] <= '9' {
                n, _ := strconv.Atoi(string(v[i]))
                if first {
                    l = n
                } else {
                    f = n
                    l = n
                    first = true
                }
            } else {
                j := i + 1
                for j < len(v) {
                    w := v[i:j+1]
                    if (slices.Contains(numbers, w)) {
                        r := getNumber(w)
                        if first {
                            l = r
                        } else {
                            f = r
                            l = r
                            first = true
                        }
                    }
                    j++
                }
            }
        }
        sum += ((f * 10) + l)
    }

    fmt.Printf("%d\n", sum)
}

func getNumber(w string) int {
    var result int
    switch w {
        case "one":
        result = 1
        case "two":
        result = 2
        case "three":
        result = 3
        case "four":
        result = 4
        case "five":
        result = 5
        case "six":
        result = 6
        case "seven":
        result = 7
        case "eight":
        result = 8
        case "nine":
        result = 9
    }

    return result
}
