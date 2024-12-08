package main

import (
	"bufio"
	"fmt"
	"os"
)

var iterations int

func main() {
    sum := 0
    readInput(&sum)
    fmt.Printf("Sum: %v | Iterations: %v\n", sum, iterations)
}

func crossLookup(xmas [][]string) int {
    sum := 0
    crossIndex := make(map[string]int)
    for i := 0; i < len(xmas); i++ {
        for j := 0; j < len(xmas[0]); j++ {
            iterations += 1
            if j > 1 { // LEFT DOWN/UP
                if i > 1 {
                    if lup := xmas[i][j] + xmas[i-1][j-1] + xmas[i-2][j-2]; lup == "MAS" || lup == "SAM" {
                        crossIndex[fmt.Sprintf("%v,%v", i-1,j-1)] += 1 
                    }
                }
                if i < len(xmas)-2 {
                    if ldown := xmas[i][j] + xmas[i+1][j-1] + xmas[i+2][j-2]; ldown == "MAS" || ldown == "SAM" {
                        crossIndex[fmt.Sprintf("%v,%v", i+1,j-1)] += 1 
                    }
                }
            }
            if j < len(xmas[0])-2 { // RIGHT UP/DOWN
                if i > 1 {
                    if rup := xmas[i][j] + xmas[i-1][j+1] + xmas[i-2][j+2]; rup == "MAS" || rup == "SAM" {
                        crossIndex[fmt.Sprintf("%v,%v", i-1,j+1)] += 1 
                    }
                }
                if i < len(xmas)-2 {
                    if rdown := xmas[i][j] + xmas[i+1][j+1] + xmas[i+2][j+2]; rdown == "MAS" || rdown == "SAM" {
                        crossIndex[fmt.Sprintf("%v,%v", i+1,j+1)] += 1 
                    }
                }
            }
        }
    }
    for _, i := range crossIndex {
        iterations += 1
        if i == 4 {
            sum += 1
        }
    }
    return sum
}

func readInput(sum *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    
    var xmas [][]string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        var xxmas []string
        for _, i := range line {
            iterations += 1
            xxmas = append(xxmas, string(i))
        }
        xmas = append(xmas, xxmas)
    }
    *sum += crossLookup(xmas)

    return nil
}
