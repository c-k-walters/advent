package main

import (
	"fmt"
	"math"

	"github.com/c-k-walters/advent/util"
)

func main() {
    fmt.Println("Day 1 Solutions")
    star1()
    star2()
}

func star1() {
    fmt.Println("Star 1:")
    a, b := util.FileToList("input-star1.txt")

    sum := 0
    for i := 0; i < len(a); i++ {
        dist := math.Abs((float64) (a[i] - b[i]))
        sum += (int) (dist)
    }
    fmt.Println(sum)
}

func star2() {

}
