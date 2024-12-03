package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var negative bool

func main() {
    sum := 0
    readInput(&sum)
    fmt.Println(sum)
}

func diffOfTwo(a, b int) int {
    if a > b {
        return a + (^b + 1)
    }
    return b + (^a + 1)
}

func adjacentChecking(nums []int, left, right, deviation int) int {
    if left <= right && deviation == 0 {
        if left-1 >= 0 && right+1 <= len(nums)-1 {
            diffLeft1 := diffOfTwo(nums[left-1], nums[left])
            diffLeft2 := diffOfTwo(nums[left+1], nums[left])
            diffRight1 := diffOfTwo(nums[right-1], nums[right])
            diffRight2 := diffOfTwo(nums[right+1], nums[right])

            if negative {
                if !(nums[left-1] > nums[left]) || !(diffLeft1 <= 3 && diffLeft1 >=1) {
                    deviation += 1
                }
                if !(nums[left+1] < nums[left]) || !(diffLeft2 <= 3 && diffLeft2 >=1) {
                    deviation += 1
                }
                if !(nums[right-1] > nums[right]) || !(diffRight1 <= 3 && diffRight1 >=1) {
                    deviation += 1
                }
                if !(nums[right+1] < nums[right]) || !(diffRight2 <= 3 && diffRight2 >=1) {
                    deviation += 1
                }
            }
            if !negative {
                if !(nums[left-1] < nums[left]) || !(diffLeft1 <= 3 && diffLeft1 >=1) {
                    deviation += 1
                }
                if !(nums[left+1] > nums[left]) || !(diffLeft2 <= 3 && diffLeft2 >=1) {
                    deviation += 1
                }
                if !(nums[right-1] < nums[right]) || !(diffRight1 <= 3 && diffRight1 >=1) {
                    deviation += 1
                }
                if !(nums[right+1] > nums[right]) || !(diffRight2 <= 3 && diffRight2 >=1) {
                    deviation += 1
                }
            }
            return adjacentChecking(nums, left+1, right-1, deviation)
        }
    }
    return deviation
}

func readInput(sum *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        subStr := strings.Split(line, " ")
        var nums []int
        for _, i := range subStr {
            num, err := strconv.Atoi(i)
            if err != nil { log.Fatalf("failed converting") }
            nums = append(nums, num)
        }
        if nums[0] > nums[len(subStr)-1] {
            negative = true
        } else if nums[0] < nums[len(subStr)-1] {
            negative = false
        } else {
            continue
        }

        if deviation := adjacentChecking(nums, 1, len(nums)-2, 0); deviation == 0 {
            *sum += 1
        }
    }
    return nil
}
