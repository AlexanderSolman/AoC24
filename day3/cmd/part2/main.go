package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var iterations int

func main() {
    now := time.Now()
    sum := 0
    readInput(&sum)
    elapsed := time.Since(now)
    fmt.Printf("Answer: %v\nIterations: %v\nTime: %v\n", sum, iterations, elapsed)
}

func multiplyAll(nums []int, index, sum int) int {
    iterations += 1
    if index < len(nums)-1 {
        sum += nums[index] * nums[index+1]
        return multiplyAll(nums, index+2, sum)
    }

    return sum
}

func readInput(sum *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    
    pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`
    numericPattern := `\d+`
    skip := false
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        var found []string
        var numbers []string
        var nums []int
        line := scanner.Text()
        regex := regexp.MustCompile(pattern)
        found = regex.FindAllString(line, -1)
        regex = regexp.MustCompile(numericPattern)
        for _, i := range found {
            iterations += 1
            if i == "don't()"{
                skip = true
            } else if i == "do()" {
                skip = false
            }
            if !skip {
                numbers = regex.FindAllString(i, -1)
                for _, j := range numbers {
                    if num, err := strconv.Atoi(j); err == nil {
                        nums = append(nums, num)
                    }
                }
            }
        }
        *sum += multiplyAll(nums, 0, 0)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("error reading input: %v\n", err)
    }
    return nil
}
