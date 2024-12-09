package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
    "math"
)

func main() {
    file, err := os.Open("cmd/2024/7/input.txt")
    if err != nil { log.Fatal(err) }

    scanner := bufio.NewScanner(file)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        r := regexp.MustCompile("[0-9]+")
        nums := r.FindAllString(line, -1)

        goal, _ := strconv.Atoi(nums[0])
        values := make([]int, len(nums[1:]))
        for i, v := range nums[1:] {
            values[i], _ = strconv.Atoi(v)
        }

        if isValidEq(goal, values, 0) { sum += goal }
    }

    fmt.Println("Sum of valid operations: ", sum)
}

func isValidEq(goal int, values []int, total int) bool {
    if len(values) == 0 { return goal == total }
    multTotal := total * values[0]
    if total == 0 { multTotal = values[0] }

    pow := math.Log10(float64(values[0]))
    concatMul := int(math.Pow10(int(pow) + 1))

    return isValidEq(goal, values[1:], multTotal) ||
    isValidEq(goal, values[1:], total + values[0]) ||
        isValidEq(goal, values[1:], total * concatMul + values[0] )
}
