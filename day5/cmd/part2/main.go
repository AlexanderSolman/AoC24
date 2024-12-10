package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var iterations int

func main() {
    sum := 0
    if err := readInput(&sum); err != nil { fmt.Printf("%v\n", err) }
    fmt.Printf("Sum: %v | Iterations: %v\n", sum, iterations)
}

func checkValidPage(rules map[string]map[string]string, pages []string, index int) bool {
    if index > 0 {
        for i := index-1; i >= 0; i-- {
            iterations += 1
            if _, ok := rules[pages[index]][pages[i]]; ok {
                return false
            }
        }
        return checkValidPage(rules, pages, index-1)
    }
    return true
}

func readInput(sum *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    
    rules := make(map[string]map[string]string)

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" { continue }
        if strings.Contains(line, "|") {
            subStr := strings.Split(line, "|")
            iterations += 1
            if _, ok := rules[subStr[0]]; !ok {
                rules[subStr[0]] = make(map[string]string)
            }
            rules[subStr[0]][subStr[1]] = subStr[1]
        } else {
            subStr := strings.Split(line, ",")
            pages := []string{}
            for _, i := range subStr {
                iterations += 1
                pages = append(pages, i)
            }
            if !checkValidPage(rules, pages, len(pages)-1) {
                for i := len(pages)-1; i >= 0; i-- {
                    for j := i-1; j >= 0; j-- {
                        iterations += 1
                        if _, ok := rules[pages[i]][pages[j]]; ok {
                            temp := pages[i]
                            pages[i] = pages[j]
                            pages[j] = temp
                        }
                    }
                }
                mid, err := strconv.Atoi(pages[len(pages)/2])
                if err != nil { return fmt.Errorf("failed converting\n", err) }
                *sum += mid
            }
        }
    }
    return nil
}
