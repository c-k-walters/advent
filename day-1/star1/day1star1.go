package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calcDistance(list1, list2 []int) (int, error) {
    sort.Ints(list1)
    sort.Ints(list2)

    if len(list1) != len(list2) {
        return -1, errors.New("lengths of the two input list did not match")
    }

    sum := 0
    for i := 0; i < len(list1); i++ {
        sum += (int) (math.Abs((float64) (list1[i] - list2[i])))
    }

    return sum, nil
}

func main() {
    file, err := os.Open("day1star1input.txt")

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    list1 := make([]int, 0, 1000)
    list2 := make([]int, 0, 1000)

    for scanner.Scan() {
        line := scanner.Text()
        pair := strings.Fields(line)
        pair0, err := strconv.Atoi(pair[0])
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("-------")
        fmt.Println(pair0)
        list1 = append(list1, pair0)

        pair1, err := strconv.Atoi(pair[1])
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(pair1)
        list2 = append(list2, pair1)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    // now we have our lists to process.

    answer, err := calcDistance(list1, list2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(answer)
}
