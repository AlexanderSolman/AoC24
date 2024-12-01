package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
    testcases := []struct {
        a,b int
        expected int
    }{
        {1,3,2},
        {2,3,1},
        {3,3,0},
        {3,4,1},
        {3,5,2},
        {4,9,5},
    }
    // 2+1+0+1+2+5

    for _, tt := range testcases {
        result := largerofTwo(tt.a, tt.b)
        fmt.Printf("Result: %v | %v = %v\n", tt.a, tt.b, tt.expected) 
        if result != tt.expected {
            t.Errorf("difference between %v and %v != %v\n", tt.a, tt.b, tt.expected)
        }
    }
}
