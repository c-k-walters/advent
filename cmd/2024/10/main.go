package main

import (
	"bufio"
	"fmt"
	"os"
)

const input = "cmd/2024/10/input.txt"
const example = "cmd/2024/10/example.txt"

func main() {
    file, err := os.Open(input)
    if err != nil { return }
    reader := bufio.NewReader(file)

    puzzle := topoMap{}
    puzzle.construct(reader)
    sum1, sum2 := puzzle.findTrailheads()
    fmt.Println("trailtails total: ", sum1)
    fmt.Println("Number of distinct paths: ", sum2)
}

type topoMap struct {
    w, h int
    terrain []int
}

func (t *topoMap) findTrailheads() (int, int) {
    sum := 0
    modifiedSum := 0
    x, y := 0, 0
    for i := 0; i < len(t.terrain); i++ {
        if i != 0 { x, y = i % t.w, i / t.w }
        if t.at(x, y) != 0 { continue }
        tails := make(map[int]bool)
        modifiedSum += t.exploreTrailAt(x, y, &tails, 0)
        sum += len(tails)
    }
    return sum, modifiedSum
}

func (t *topoMap) exploreTrailAt(x, y int, tails *map[int]bool, curr int) int {
    // breadth first search
    // don't need to keep track of where we have gone since we can only move up
    // check l, r, u, d
    at := t.at(x, y)
    if at != curr { return 0 }
    if at == 9 { (*tails)[y*t.w + x] = true; return 1 }

    return t.exploreTrailAt(x + 1, y, tails, curr+1) +
    t.exploreTrailAt(x, y + 1, tails, curr+1) +
    t.exploreTrailAt(x - 1, y, tails, curr+1) +
    t.exploreTrailAt(x, y - 1, tails, curr+1)
}

func (t *topoMap) at(x, y int) int {
    // A little clever thing here. In this puzzle you can only increase your
    // elevation by 1 at a time. Therefore no tile could ever traverse to a 
    // tile of elevation 11, and we may ignore bounds checks in our algorithm
    // later on.
    if x < 0 || x >= t.w || y < 0 || y >= t.h { return 11 }
    return t.terrain[y * t.w + x]
}

func (t *topoMap) construct(reader *bufio.Reader) {
    t.terrain = make([]int, 0)
    x, y := 0, 0
    for {
        r, _, err := reader.ReadRune()
        if err != nil { break }
        switch r {
        case '\n':
            t.w = x
            y++
            x = 0
        default:
            t.terrain = append(t.terrain, int(r-'0'))
            x++
        }
    }
    t.h = y
}
