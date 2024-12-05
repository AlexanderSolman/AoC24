package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
    negative bool
    global_iterations int
    adj_iterations int
    brute_iterations int
    accepted = map[int]int{
        -3: -3,
        -2: -2,
        -1: -1,
        1: 1,
        2: 2,
        3: 3,
    }
)


func main() {
    start := time.Now()
    sum1 := 0
    sum2 := 0
    if err := readInput(&sum1, &sum2); err != nil {
        fmt.Printf("solution failed: %v\n", err)
    }
    end := time.Since(start)
    fmt.Println("Adjacent:", sum1,"Bruteforce:", sum2)
    fmt.Println("Adjacent Iterations:", global_iterations+adj_iterations,"Bruteforce Iterations:", global_iterations+brute_iterations)
    fmt.Printf("Execution time: %v\n", end.Seconds())
}

func diffOfTwo(a, b int) int {
    if a > b {
        return a + (^b + 1)
    }
    return b + (^a + 1)
}

/*----------- DOC SOLUTION -------------*/

func adjacentChecking(nums []int, left, right int, neg, pos int) bool {
    adj_iterations += 1
    if left < right {
        if left-1 >= 0 && right+1 <= len(nums)-1 {
            
            if (nums[left] - nums[left-1]) < 0 { 
                neg += 1 
            } else {
                pos += 1
            }
            if (nums[left+1] - nums[left]) < 0 { 
                neg += 1 
            } else {
                pos += 1
            }
            if (nums[right] - nums[right-1]) < 0 { 
                neg += 1 
            } else {
                pos += 1
            }
            if (nums[right+1] - nums[right]) < 0 { 
                neg += 1 
            } else {
                pos += 1
            }

            if _, ok := accepted[diffOfTwo(nums[left-1], nums[left])]; !ok { return false }
            if _, ok := accepted[diffOfTwo(nums[left+1], nums[left])]; !ok { return false }
            if _, ok := accepted[diffOfTwo(nums[right-1], nums[right])]; !ok { return false }
            if _, ok := accepted[diffOfTwo(nums[right+1], nums[right])]; !ok { return false }
            
            return adjacentChecking(nums, left+1, right-1, neg, pos)
        }
    }
    if (pos > 0 && neg == 0) || (pos == 0 && neg > 0) {
        return true
    }
    return false
}

func retryAdjacent(nums []int) bool {
    if adjacentChecking(nums, 1, len(nums)-2, 0, 0) { return true }
    for i := 0; i < len(nums); i++ {
        adjustednums := make([]int, 0, len(nums)-1)
        adjustednums = append(adjustednums, nums[:i]...)
        adjustednums = append(adjustednums, nums[i+1:]...)
        if adjacentChecking(adjustednums, 1, len(adjustednums)-2, 0, 0) {
            return true
        }
    }
    return false
}

/*----------- BRUTE FORCE SOLUTION ---------------*/

func bruteForce(nums []int) bool {
    neg := 0
    for i := 1; i < len(nums); i++ {
        brute_iterations += 1
        diff := nums[i] - nums[i-1]

        if diffOfTwo(nums[i], nums[i-1]) < 1 || diffOfTwo(nums[i], nums[i-1]) > 3 {
            return false
        }
        if diff < 0 {
            neg -= 1
        } else if diff > 0 {
            neg += 1
        }
    }
    if neg < 0 { neg *= -1 }
    if neg != len(nums)-1 {
        return false
    }
    return true
}

func retryBruteforce(nums []int) bool {
    if bruteForce(nums) { return true }
    for i := 0; i < len(nums); i++ {
        adjustednums := make([]int, 0, len(nums)-1)
        adjustednums = append(adjustednums, nums[:i]...)
        adjustednums = append(adjustednums, nums[i+1:]...)
        if bruteForce(adjustednums) {
            return true
        }
    }
    return false
}

func readInput(sum1, sum2 *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        global_iterations += 1
        line := scanner.Text()
        subStr := strings.Split(line, " ")
        var nums []int
        for _, i := range subStr {
            global_iterations += 1
            num, err := strconv.Atoi(i)
            if err != nil { fmt.Errorf("failed converting: %v\n", err) }
            nums = append(nums, num)
        }

        if nums[0] > nums[len(nums)-1] {
            negative = true
        } else if nums[0] < nums[len(nums)-1] {
            negative = false
        } else {
            if nums[0] > nums[len(nums)-1] {
                negative = true
            } else if nums[0] < nums[len(nums)-1] {
                negative = false
            }
        }

        if retryAdjacent(nums) {
            *sum1 += 1
        }
        if retryBruteforce(nums) {
            *sum2 += 1
        }
    }
    return nil
}
