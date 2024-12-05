package main

import (
    "bufio"
    "strings"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    fmt.Println("Day 2 Solutions")
    star1()
    star2()
}

func star1() {
    fmt.Println("Star 1:")

    file, err := os.Open("cmd/2/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    safeCount := 0
    for scanner.Scan() {
        line := scanner.Text()
        report := strings.Fields(line)
        safe, _ := isSafeReport(report, -1)
        if safe {
            safeCount++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(safeCount)
}

func isSafeReport(report []string, skipInd int) (bool, int) {
    if len(report) < 2 {
        return true, -1
    }

    a, err0 := strconv.Atoi(report[0])
    b, err1 := strconv.Atoi(report[1])
    if len(report) == 2{
        if err0 != nil || err1 != nil {
            log.Fatal(err0)
            log.Fatal(err1)
        }
        return outOfBounds(a < b, a, b), -1
    }

    inc := false
    dec := false
    prev := -1

    for i := 0; i < len(report); i++ {

        curr, err := strconv.Atoi(report[i])
        if err != nil {
            log.Fatal(err)
        }

        if i == skipInd {
           continue 
        }

        if prev == -1 {
            prev = curr
            continue
        }

        if !inc && !dec {
            inc = prev < curr
            dec = prev > curr
        }

        if outOfBounds((inc || !dec), prev, curr) {
            return false, i
        } else {
            prev = curr
        }
    }

    return true, -1
}


func outOfBounds(increasing bool, a, b int) bool {
    if a == b {
        return true
    }

    if increasing {
        return b > a+3 || b < a
    } else {
        return a > b+3 || a < b
    }
}

func star2() {
    fmt.Println("Star 2:")

    file, err := os.Open("cmd/2/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    safeCount := 0
    for scanner.Scan() {
        line := scanner.Text()
        report := strings.Fields(line)
        if isSafeWithRemove(report) {
            safeCount++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(safeCount)
}


func isSafeWithRemove(report []string) bool {
    safe, failedInd := isSafeReport(report, -1)

    if safe {
        return true
    }

    lowBound := failedInd -2
    if lowBound < 0 {
        lowBound = 0
    }

    uppBound := failedInd + 2
    if uppBound > len(report) {
        uppBound = len(report)
    }

    for i := lowBound; i < uppBound; i++ {
        safe, _ := isSafeReport(report, i)
        if safe {
            return true
        }
    }
    return false
    
}
