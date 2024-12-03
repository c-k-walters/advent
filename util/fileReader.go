package util 

import (
    "log"
    "os"
    "strconv"
    "bufio"
    "strings"
)

func FileToList(fileName string) ([]int, []int) {
    file, err := os.Open(fileName)

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
        list1 = append(list1, pair0)

        pair1, err := strconv.Atoi(pair[1])
        if err != nil {
            log.Fatal(err)
        }
        list2 = append(list2, pair1)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return list1, list2
}
