package main

import (
    "advent/util"
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
)

func main() {
    fmt.Println("Advent of Code 2023")
    fmt.Println("Day 1")
    star1()
    star2()
}

func star1() {
    fmt.Println("Star 1:")
    util.FileReader("cmd/2023/1/input.txt", getSandwhichDigits)
}

func getSandwhichDigits(reader *bufio.Reader) {
    sum := 0
    fDig := -1
    lDig := -1

    for {
        r, _, err := reader.ReadRune()
        if err != nil {
            break
        }

        switch r {
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            lDig = int (r - '0')
            if fDig == -1 {
                fDig = int(r - '0')
            }
        case '\n':
            sum += (fDig * 10) + lDig
            fDig = -1
            lDig = -1
        }
    }

    fmt.Println("Sum of digits is: ", sum)
}



func star2() {
    fmt.Println("Star 2:")
    digitMap := map[string]int {
        "1": 1,
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    file, err := os.Open("cmd/2023/1/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    r, _ := regexp.Compile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
    r_rev, _ := regexp.Compile("enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|[0-9]")
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        digits := r.FindAllString(line, 1)
        rev_digits := r_rev.FindAllString(reverse(line), 1)

        fDig := digitMap[digits[0]]
        lDig := digitMap[reverse(rev_digits[0])] 
        sum += (fDig * 10) + lDig
    }
    fmt.Println("Sum of digits", sum)
}

func reverse(str string) string {
    runes := []rune(str)
    for i, j := 0, len(str)-1; i<j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
