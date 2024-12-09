package main

import (
    "os"
    "bufio"
    "fmt"
)

const example = "cmd/2024/9/example.txt"
const input = "cmd/2024/9/input.txt"

func main() {
    file, err := os.Open(input)
    if err != nil { return }

    reader := bufio.NewReader(file)
    memory := constructMemory(reader)
    defragPreserveFiles(&memory)
    
    fmt.Println("Memory checksum: ", checksum(memory))
}

func defragPreserveFiles(mem *LinkedList2[MemoryBlock]) {
    // for file in decreasing order
        // for block of free space starting from the left
    // initial conditions
    id := mem.tail.val.id
    // 8, 9 worked so far
    // 5, 6 didn't work

    for id > 0 {
        fileHead, fileTail, size := getFile(mem, id)
        freeHead, freeTail := getFreeSpace(mem, size, fileHead)
        
        if fileHead != nil && fileTail != nil && freeHead != nil && freeTail != nil {   
            swapLinks(fileHead, fileTail, freeHead, freeTail)
        }
        id--
    }
}

// Swap two runs of links, assuming the inside of each run has no nil values
func swapLinks[T any](lh, lt, rh, rt *Node[T]) {
    for l, r := lh, rh; ;l, r = l.next, r.next {
        l.val, r.val = r.val, l.val
        if l == lt && r == rt { return }
    }
}

func getFreeSpace(mem *LinkedList2[MemoryBlock], size int, lTail *Node[MemoryBlock]) (head, tail *Node[MemoryBlock]) {
    currSize := 0
    for l := mem.head; l != nil; l = l.next {
        if currSize == size { return }
        if l == lTail { head = nil; tail = nil; return }
        if !l.val.isFree { currSize = 0; head = nil; tail = nil; continue }
        
        // found free space
        if head == nil { head = l }
        tail = l
        currSize++
    }
    if currSize < size { head = nil; tail = nil }
    return 
}

func getFile(mem *LinkedList2[MemoryBlock], id int) (head, tail *Node[MemoryBlock], size int) {
    for l := mem.head; l != nil; l = l.next {
        if l.val.id == id && !l.val.isFree {
            if head == nil { head = l }
            tail = l
            size++
        }
    }
    return 
}

func defrag(mem *LinkedList2[MemoryBlock]) {
    for l, r := mem.head, mem.tail; l != r; {
        for !l.val.isFree {
            l = l.next
            if l == r { return }
        }
        for r.val.isFree {
            r = r.prev
            if l == r { return }
        }
        // swap l with r
        l.val, r.val = r.val, l.val
    }
}

func checksum(mem LinkedList2[MemoryBlock]) int {
    check := 0
    loc := 0
    for l := mem.head; l != nil; l = l.next {
        if !l.val.isFree { check += loc * l.val.id }
        loc++
    }
    return check
}

type MemoryBlock struct {
    id int
    isFree bool
}

type LinkedList2[T any] struct {
    head, tail *Node[T]
}

type Node[T any] struct {
    next, prev *Node[T]
    val T
}

func constructMemory(reader *bufio.Reader) (mem LinkedList2[MemoryBlock]) {
    freeBlock := false
    id := 0
    for { 
        r, _, err := reader.ReadRune()
        if err != nil || r == '\n' { return }

        val := int(r - '0')
        for i := 0; i < val; i++ {
            mem.Push(MemoryBlock{ id: id, isFree: freeBlock })
        }
        freeBlock = !freeBlock
        if !freeBlock { id++ }
    }
}

func printMemory(mem LinkedList2[MemoryBlock]) {
    fmt.Println("Memory:")
    for curr := mem.head; curr != nil; curr = curr.next {
        if curr.val.isFree { fmt.Print(".") } else {
            fmt.Print(curr.val.id)
        }
    }
    fmt.Print("\n")
}

func (lst *LinkedList2[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &Node[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &Node[T]{val: v}
        lst.tail.next.prev = lst.tail
        lst.tail = lst.tail.next
    }
}
