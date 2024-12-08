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
    fmt.Printf("Sum: %v Iterations: %v\n", sum, iterations)
}

func crossLookup(xmas [][]string)int {
    sum := 0
    for i := 0; i < len(xmas); i++ {
        for j := 0; j < len(xmas[0]); j++ {
            iterations += 1
            if j > 2 { // LEFT DOWN/UP
                if i > 2 {
                    if lup := xmas[i][j] + xmas[i-1][j-1] + xmas[i-2][j-2] + xmas[i-3][j-3]; lup == "XMAS" || lup == "SAMX" {
                        sum += 1
                    }
                }
                if i < len(xmas)-3 {
                    if ldown := xmas[i][j] + xmas[i+1][j-1] + xmas[i+2][j-2] + xmas[i+3][j-3]; ldown == "XMAS" || ldown == "SAMX" {
                        sum += 1
                    }
                }
            }
            if j < len(xmas[0])-3 { // RIGHT UP/DOWN
                if i > 2 {
                    if rup := xmas[i][j] + xmas[i-1][j+1] + xmas[i-2][j+2] + xmas[i-3][j+3]; rup == "XMAS" || rup == "SAMX" {
                        sum += 1
                    }
                }
                if i < len(xmas)-3 {
                    if rdown := xmas[i][j] + xmas[i+1][j+1] + xmas[i+2][j+2] + xmas[i+3][j+3]; rdown == "XMAS" || rdown == "SAMX" {
                        sum += 1
                    }
                }
            }
            if i > 2 { // UP/DOWN
                if up := xmas[i][j] + xmas[i-1][j] + xmas[i-2][j] + xmas[i-3][j]; up == "XMAS" || up == "SAMX" {
                    sum += 1
                }
            }
            if i < len(xmas)-3 {
                if down := xmas[i][j] + xmas[i+1][j] + xmas[i+2][j] + xmas[i+3][j]; down == "XMAS" || down == "SAMX" {
                    sum += 1
                }
            }
            if j > 2 {
                if left := xmas[i][j] + xmas[i][j-1] + xmas[i][j-2] + xmas[i][j-3]; left == "XMAS" || left == "SAMX" {
                    sum += 1
                }
            }
            if j < len(xmas[0])-3 {
                if right := xmas[i][j] + xmas[i][j+1] + xmas[i][j+2] + xmas[i][j+3]; right == "XMAS" || right == "SAMX" {
                    sum += 1
                }
            }
        }
    }
    return sum / 2
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
