package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
    testcases := []struct {
        nums []int
        expected int 
    }{
        {[]int{7,6,4,2,1}, 0},
        {[]int{1,2,7,8,9}, 1},
        {[]int{9,7,6,2,1}, 1},
        {[]int{1,3,2,4,5}, 1},
        {[]int{8,6,4,4,1}, 1},
        {[]int{1,3,6,7,9}, 0},
        {[]int{1,3,2,7,9}, 2},
        {[]int{7,4,4,2,4}, 2},
        {[]int{5,5,5,5,4}, 2},
        {[]int{1,4,7,10,13}, 0},
    }
    
    for _, tt := range testcases {
        if tt.nums[0] > tt.nums[len(tt.nums)-1] {
            negative = true
        } else if tt.nums[0] < tt.nums[len(tt.nums)-1] {
            negative = false
        }
        result := adjacentChecking(tt.nums, 0, len(tt.nums), 0)
        fmt.Printf("testcase: %v: ExpDev: %v: ActDev: %v\n", tt.nums, tt.expected, result)
        if result != tt.expected {
            t.Errorf("testcase: %v: Expected: %v\n", result, tt.expected)
        }
    }
}
