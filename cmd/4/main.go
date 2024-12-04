
package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
)

func main() {
    fmt.Println("Day 4 Solutions")
    star1()
    star2()
}

func star1() {
    fmt.Println("Star 1:")

    file, err := os.Open("cmd/4/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    puz := Puzzle{}
    puz.construct(reader)

    fmt.Println(puz.find([]rune {'X', 'M', 'A', 'S' }))
}

func star2() {
    fmt.Println("Star 2:")

    file, err := os.Open("cmd/4/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    puz := Puzzle{}
    puz.construct(reader)
    fmt.Println(puz.findCrossMas())
}

type Puzzle struct {
    data []rune
    w, h int
}

func (puz *Puzzle) at(x, y int) rune {
    if x < 0 || x >= puz.w || y < 0 || y >= puz.h {
        return -1
    }
    return puz.data[y*puz.w+x]
}

func (puz *Puzzle) construct(reader *bufio.Reader) {
    x, y := 0, 0
    for {
        r, _, err := reader.ReadRune()
        if err != nil {
            break
        }

        switch r {
        case 'X', 'M', 'A', 'S':
            puz.data = append(puz.data, r)
            x++
        case '\n':
            y++
            puz.w = x
            x = 0
        }
    }
    puz.h = y
}

func (puz *Puzzle) find(str []rune) int {
    dirs := []Direction {
        {-1, -1},
        {0, -1},
        {-1, 0},
        {1, 1},
        {1, 0},
        {0, 1},
        {1, -1},
        {-1, 1},
    }
    count := 0
    for x := 0; x < puz.w; x++ {
        for y := 0; y < puz.h; y++ {
            count += puz.findAt(x, y, str, dirs)
        }
    }
    return count
}

type Direction struct { Xp, Yp int }

func (puz *Puzzle) findAt(x, y int, str []rune, directions []Direction) int {
    if len(str) == 0 {
        return 1
    }

    if str[0] != puz.at(x, y) {
        return 0
    }

    count := 0
    for _, dir := range directions {
        count += puz.findAt(x+dir.Xp, y+dir.Yp, str[1:], []Direction{ dir })
    }

    return count
}

func (puz *Puzzle) findCrossMas() int {
    count := 0
    for x := 0; x < puz.w; x++ {
        for y := 0; y < puz.h; y++ {
            if puz.at(x, y) != 'A' { continue }
            count += puz.evaluateCross(x, y)
        }
    }
    return count
}

type Tuple[A, B any] struct {
    f A
    s B
}

func (puz *Puzzle) evaluateCross(x, y int) int {
    pairs := []Tuple[Direction, Direction] {
        {Direction{1, 1}, Direction{-1, -1}}, 
        {Direction{-1, -1}, Direction{1, 1}}, 
        {Direction{1, -1}, Direction{-1, 1}}, 
        {Direction{-1, 1}, Direction{1, -1}}, 
    }

    count := 0
    for _, pair := range pairs {
        count += puz.findAt(
            x+pair.f.Xp,
            y+pair.f.Yp,
            []rune{'M', 'A', 'S'},
            []Direction{pair.s},
        )
    }
    if count == 2 {
        return 1
    } else {
        fmt.Println(count)
        return 0
    }
}
