package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    fmt.Println("Day 3 Solutions")
    star1()
    star2()
}

type State struct {
    // rough state machine
    // 0, 1, 2 - mul keyword
    // 3 - left (
    // 4 - left op
    // 5 - comma
    // 6 - right op
    // 7 - right )
    // 8 - d
    // 9 - o
    // 10 - n
    // 11 - '
    // 12 - t
    // 13 - right ) for do
    // 14 - right ) for don't
    mode int
    useDisable bool
    disableMul bool

    left MulOperand
    right MulOperand

    sum int
}

type MulOperand struct {
    value int
    digits int
}

func (state *State) consume(r rune) {
    switch r {
        case 'm':
            state.resetBuf()
            state.mode++
        case 'u':
            if state.mode == 1 { state.mode++ } else { state.resetBuf() }
        case 'l':
            if state.mode == 2 { 
                state.mode++
            } else { state.resetBuf() }
        case '(':
            switch state.mode {
            case 3:
                state.mode++
            case 9:
                state.mode = 13
            case 12:
                state.mode = 14
            default:
                state.resetBuf()
            }
        case ',':
            if state.mode == 4 && state.left.digits > 0 { state.mode = 6 } else { state.resetBuf() }
        case ')':
            switch state.mode {
            case 6:
                if state.right.digits > 0 {
                    state.processBuf()
                } else {
                    state.resetBuf()
                }
            case 13:
                state.disableMul = false
            case 14:
                state.disableMul = true
            default:
                state.resetBuf()
            }
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            switch {
            case state.mode == 4:
                state.leftDigit(r)
            case state.mode == 6:
                state.rightDigit(r)
            default:
                state.resetBuf()
            }
        case 'd':
            state.resetBuf()
            state.mode = 8
        case 'o':
            if state.mode == 8 { state.mode = 9 } else { state.resetBuf() }
        case 'n':
            if state.mode == 9 { state.mode = 10 } else { state.resetBuf() }
        case '\'':
            if state.mode == 10 { state.mode = 11 } else { state.resetBuf() }
        case 't':
            if state.mode == 11 { state.mode = 12 } else { state.resetBuf() }
        default:
            state.resetBuf()
    }
}

func (state *State) resetBuf() {
    state.mode = 0
    state.left = MulOperand{}
    state.right = MulOperand{}
}

func (state *State) processBuf() {
    if state.useDisable && state.disableMul {
        state.resetBuf()
        return
    }
    state.sum += state.left.value * state.right.value
    state.resetBuf()
}

func (state *State) rightDigit(r rune) {
    if state.right.digits >= 3 {
        state.resetBuf()
    } else {
        state.right.appendRune(r)
    }
}

func (state *State) leftDigit(r rune) {
    if state.left.digits >= 3 {
        state.resetBuf()
    } else {
        state.left.appendRune(r)
    }
}

func (op *MulOperand) appendRune(r rune) {
    op.value *= 10
    op.value += int (r - '0')
    op.digits++
}

func star1() {
    fmt.Println("Star 1:")

    file, err := os.Open("cmd/3/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var state State

    for {
        r, _, err := reader.ReadRune()
        if err != nil {
            break
        }
        state.consume(r)
    }

    fmt.Println(state.sum)
}

func star2() {
    fmt.Println("Star 1:")

    file, err := os.Open("cmd/3/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var state State
    state.useDisable = true

    for {
        r, _, err := reader.ReadRune()
        if err != nil {
            break
        }
        state.consume(r)
    }

    fmt.Println(state.sum)
}
