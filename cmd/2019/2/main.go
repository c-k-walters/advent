package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("cmd/2019/2/input.txt")
    if err != nil { return }

    scanner := bufio.NewScanner(file)
    _ = scanner.Scan()
    program := scanner.Text()
    split := strings.Split(program, ",")
    intcodeProgram := IntcodeProgram{program: make([]int, len(split))}
    for i, s := range split { intcodeProgram.program[i], _ = strconv.Atoi(s) }
    intcodeProgram.program[1] = 12
    intcodeProgram.program[2] = 2
    
    intcodeProgram.run()

    fmt.Println("position 0 reads: ", intcodeProgram.program[0])
}


// really should add checking for unexpected 99 exit code
// will add later
func (p *IntcodeProgram) run() {
    p.bounds = len(p.program)
    for {
        p.opcode = p.program[p.ptr]
        p.ptr++
        switch p.opcode {
        case 99:
            return
        case 1: // add 
            p.add()
        case 2:
            p.multiply()
        }
    }
}

func (p *IntcodeProgram) multiply() {
    p.f, p.s, p.t = p.program[p.ptr], p.program[p.ptr+1], p.program[p.ptr+2]
    p.program[p.t] = p.program[p.f] * p.program[p.s]
    fmt.Println("multiply: ", p.program[p.t])
    p.ptr = p.ptr + 3
}

func (p *IntcodeProgram) add() {
    p.f, p.s, p.t = p.program[p.ptr], p.program[p.ptr+1], p.program[p.ptr+2]
    p.program[p.t] = p.program[p.f] + p.program[p.s]
    fmt.Println("add: ", p.program[p.t])
    p.ptr = p.ptr + 3
}

type IntcodeProgram struct {
    program []int
    ptr, opcode, f, s, t int
    bounds int
}
