package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
    testcases := []struct {
        nums []int
        expected bool
    }{
        {[]int{7,6,4,2,1}, true},
        {[]int{1,2,7,8,9}, false},
        {[]int{9,7,6,2,1}, false},
        {[]int{1,3,2,4,5}, true},
        {[]int{8,6,4,4,1}, true},
        {[]int{1,3,6,7,9}, true},
        {[]int{95,97,99,97,94}, false},
        {[]int{32,35,36,35,38,40,43,44}, true},
        {[]int{63,66,65,67,69,69}, false},
        {[]int{42,44,41,43,45,47,48,52}, false},
        {[]int{69,71,72,71,76}, false},
        {[]int{69,71,72,71,69}, false},
        {[]int{1,1,1,1,1,1,1}, false},
        {[]int{11,11,15,16,17}, false},
        {[]int{12,11,14,13,15}, false},
        {[]int{22,23,23,24,25}, true},
        {[]int{69,68,68,65,62}, true},
        {[]int{9,7,6,5,9}, true},
        {[]int{9,7,4,5,9}, false},
    }
    
    for _, tt := range testcases {
        if tt.nums[0] > tt.nums[len(tt.nums)-1] {
            negative = true
        } else if tt.nums[0] < tt.nums[len(tt.nums)-1] {
            negative = false
        }
        
        result := retryBruteforce(tt.nums)
        fmt.Printf("Report: %v: Result: %v\n", tt.nums, result)
        if result != tt.expected {
            t.Errorf("Result: %v\n", result)
        }
    }
}
