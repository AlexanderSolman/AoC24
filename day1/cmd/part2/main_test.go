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
        {3,3,9},
        {4,1,4},
        {2,0,0},
        {1,0,0},
        {3,3,9},
        {3,3,9},
    }
    // 2+1+0+1+2+5

    for _, tt := range testcases {
        result := similiarity(tt.a, tt.b)
        fmt.Printf("Result: %v | %v = %v\n", tt.a, tt.b, tt.expected) 
        if result != tt.expected {
            t.Errorf("difference between %v and %v != %v\n", tt.a, tt.b, tt.expected)
        }
    }
}
