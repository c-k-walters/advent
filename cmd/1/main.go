package main

import (
	"fmt"

	"github.com/c-k-walters/advent"
	"github.com/c-k-walters/advent/util"
)

func main() {
    fmt.Println("Day 1 Solutions")
}

func star1() {
    fmt.Println("Star 1:")
    a, b := util.fileToLists("input-star1.txt")
}

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

func star2() {

}
