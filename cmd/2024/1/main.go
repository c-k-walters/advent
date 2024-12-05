package main

import (
	"fmt"
	"math"
    "advent/util"
    "sort"
)

func main() {
    fmt.Println("Day 1 Solutions")
    star1()
    star2()
}

func star1() {
    fmt.Println("Star 1:")
    a, b := util.FileToList("cmd/1/input-star1.txt")

    sort.Ints(a)
    sort.Ints(b)

    sum := 0
    for i := 0; i < len(a); i++ {
        dist := math.Abs((float64) (a[i] - b[i]))
        sum += (int) (dist)
    }
    fmt.Println(sum)
}

func star2() {
    fmt.Println("Star 2:")
    a, b := util.FileToList("cmd/1/input-star1.txt")

    freqmap := util.ToFrequencyMap(b)
    similarityScore := 0

    for i := 0; i < len(a); i++ {
        elem, ok := freqmap[a[i]]
        if ok {
            similarityScore += elem * a[i]
        }
    }

    fmt.Println(similarityScore)
}
