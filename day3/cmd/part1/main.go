package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


func main() {
    sum := 0
    readInput(&sum)
    fmt.Printf("%v\n", sum)
}

func multiplyAll(nums []int, index, sum int) int {
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
    
    pattern := `mul\(\d{1,3},\d{1,3}\)`
    numericPattern := `\d+`
    
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
            numbers = regex.FindAllString(i, -1)
            for _, j := range numbers {
                if num, err := strconv.Atoi(j); err == nil {
                    nums = append(nums, num)
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
