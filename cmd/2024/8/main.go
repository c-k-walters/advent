package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)

func main() {
    file, err := os.Open("cmd/2024/8/input.txt")
    if err != nil { log.Fatal(err) }

    reader := bufio.NewReader(file)
     
    grid := Grid{}
    grid.construct(reader)
    
    affected := make(map[Tile]bool)
    for _, tiles := range grid.antennas {
        affectedTiles := getAffectedTiles(tiles, grid.w, grid.h)
        for tile := range affectedTiles {
            affected[tile] = true
        }
    }
    fmt.Println("Affected Tiles: ", len(affected))
}

func getAffectedTiles(antennas []Tile, w, h int) map[Tile]bool {
    affectedTiles := make(map[Tile]bool)
    
    for i := 0; i < len(antennas) -1; i++ {
        for j := i+1; j < len(antennas); j++ {
            l, r := antennas[i], antennas[j]
            slope := Tile{ r.x-l.x, r.y-l.y }
            affectedTiles[l], affectedTiles[r] = true, true
            for x, y := l.x, l.y; x >= 0 && x < w && y >= 0 && y < h; x, y = x-slope.x, y-slope.y {
                affectedTiles[Tile{x, y}] = true
            }
            for x, y := r.x, r.y; x >= 0 && x < w && y >= 0 && y < h; x, y = x+slope.x, y+slope.y {
                affectedTiles[Tile{x, y}] = true
            }
        }
    }

    return affectedTiles
}

type Grid struct {
    w, h int
    antennas map[rune][]Tile
}

type Tile struct {
    x, y int
}

func (g *Grid) construct(reader *bufio.Reader) {
    x, y := 0, 0
    g.antennas = make(map[rune][]Tile)
    for {
        r, _, err := reader.ReadRune()
        if err != nil { break }
        
        switch r {
        case '\n':
            g.w = x
            y++
            x = 0
        case '.':
            x++
        default:
            g.antennas[r] = append(g.antennas[r], Tile{x, y})
            x++
        }
    }
    g.h = y
}
