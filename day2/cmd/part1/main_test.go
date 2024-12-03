package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
    testcases := []struct {
        nums []string
        expected bool
    }{
        {[]string{"7","6","4","2","1"}, true},
        {[]string{"1","2","7","8","9"}, false},
        {[]string{"9","7","6","2","1"}, false},
        {[]string{"1","3","2","4","5"}, false},
        {[]string{"8","6","4","4","1"}, false},
        {[]string{"1","3","6","7","9"}, true},
    }
    
    for _, tt := range testcases {
        if tt.nums[0] > tt.nums[len(tt.nums)-1] {
            negative = true
        } else if tt.nums[0] < tt.nums[len(tt.nums)-1] {
            negative = false
        }
        result := adjacentChecking(tt.nums, 0, len(tt.nums), true, true)
        fmt.Printf("testcase: %v: Exp: %v: Is: %v\n", tt.nums, tt.expected, result)
        if result != tt.expected {
            t.Errorf("testcase: %v: Expected: %v\n", result, tt.expected)
        }
    }
}
