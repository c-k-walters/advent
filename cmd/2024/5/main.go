package main

import (
	"bufio"
	"log"
    "fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
    star1()
    star2()
}

// how does this work
// a rule is structures as dd|ee
// where dd must be before ee
// a update can look like:
// 75, 47, 61, 53, 29
// so iterating through -- 
// we pass 75. nothing before
// we pass 47, 75 before. check that rules[47] does not contain 75
// 61 -> check that rules[61] does not contain [75, 47]
func star1() {
    rules := getRules()
    updates := getUpdates()
    fmt.Println("got rules and updates")
    sum := 0
    for _, update := range updates {
        valid := true
        activeRuleValues := make([]int, 0)

        for _, v := range update {
            // chcek that rules[v] does not contain any in active rules
            xtion := intersection(rules[v], activeRuleValues)
            if len(xtion) != 0 { 
                valid = false
                continue
            }
            activeRuleValues = append(activeRuleValues, v)
        }

        if valid {
            sum += update[len(update) / 2]
        }
    }
    fmt.Println("sum of updates: ", sum)
}

func star2() {
    rules := getRules()
    updates := getUpdates()
    sum := 0

    for _, update := range updates {
        // check if valid
        swap := false

        activeRuleValues := make([]int, 0)
        for i := 0; i < len(update); {
            xtion := intersection(rules[update[i]], activeRuleValues)
            if len(xtion) != 0 {
                swap = true
                update[i], update[i-1] = update[i-1], update[i]
                i = 0
                activeRuleValues = []int{}
            } else {
                activeRuleValues = append(activeRuleValues, update[i])
                i++
            }
        }
        // sum if was invalid originally
        if swap {
            sum += update[len(update) / 2]
        }
    }

    fmt.Println("sum of corrected updates: ", sum)
}

func getRules() map[int][]int {
    file, err := os.Open("cmd/2024/5/rules.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    rules := make(map[int][]int)
    for scanner.Scan() {
        line := scanner.Text()
        reg := regexp.MustCompile("[0-9]+")
        parts := reg.FindAllString(line, 2)

        l, _ := strconv.Atoi(parts[0])
        r, _ := strconv.Atoi(parts[1])

        elem, ok := rules[l]
        if ok {
            rules[l] = append(elem, r)
        } else {
            rules[l] = []int{r}
        }
    }
    return rules
}

func getUpdates() [][]int {
    
    file, err := os.Open("cmd/2024/5/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    updates := make([][]int, 0, 0)
    for scanner.Scan() {
        line := scanner.Text()
        reg := regexp.MustCompile("[0-9]+")

        found := reg.FindAllString(line, -1)
        update := make([]int, len(found))

        for ind, elem := range found {
            update[ind], _ = strconv.Atoi(elem)
        }

        updates = append(updates, update)
    }
    return updates
}

func intersection(a, b []int) []int {
    cross := make(map[int]int)
    xtion := make([]int, 0)

    for _, va := range a {
        cross[va] = 1
    }

    for _, vb := range b {
        _, ok := cross[vb]
        if ok {
            xtion = append(xtion, vb)
        }
    }
    return xtion
}
