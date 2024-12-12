package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const input = "cmd/2024/11/input.txt"
const example = "cmd/2024/11/example.txt"

func main() {
    file, err := os.Open(input)
    if err != nil { return }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()
    exp := regexp.MustCompile("[0-9]+")
    strings := exp.FindAllString(line, -1)
    fmt.Println(strings)

    stoneCount := 0
    seen := make(map[answer]int)
    for _, str := range strings {
        val, _ := strconv.Atoi(str)
        stoneCount += countStones(val, 75, &seen)
    }

    fmt.Println("Stones: ", stoneCount)
}

type answer struct {
    stone, blinks int
}

// need to keep track of solutions we have seen
// map
// how to we solve how many stones  
func countStones(stone int, blinks int, seen *map[answer]int) int {
    if blinks == 0 { 
        return 1
    }
    stored, ok := (*seen)[answer{stone, blinks}]
    if ok { return stored }
    var stones int
    switch {
    case stone == 0:
        stones = countStones(1, blinks-1, seen)
    case evenDigits(stone):
        length := len(strconv.Itoa(stone))
        divisor := math.Pow10(length/2)
        stones = countStones(stone / int(divisor), blinks - 1, seen) + 
            countStones(stone % int(divisor), blinks - 1, seen)
    default:
        stones = countStones(stone*2024, blinks-1, seen)
    }
    (*seen)[answer{stone, blinks}] = stones
    return stones 
}

func evenDigits(num int) bool {
    str := strconv.Itoa(num)
    return len(str) % 2 == 0
}
