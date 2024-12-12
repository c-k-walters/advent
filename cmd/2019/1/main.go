package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    file, err := os.Open("cmd/2019/1/input.txt")
    if err != nil { return }

    scanner := bufio.NewScanner(file)

    fuel := 0
    extraFuel := 0
    for scanner.Scan() {
        line := scanner.Text()

        mass, _ := strconv.Atoi(line)
        currentFuel := mass / 3 - 2
        fuel += currentFuel
        extraFuel += remainingFuel(currentFuel)
    }

    fmt.Println("Fuel used in part 1: ", fuel)
    fmt.Println("Fuel after accounting for recur: ", fuel + extraFuel)
}

func remainingFuel(total int) int {
    extra := total / 3 - 2
    if extra < 0 { return 0 }
    
    return extra + remainingFuel(extra)
}
