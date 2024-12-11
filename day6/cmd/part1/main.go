package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type positions struct {
    i, j int
}

var iterations int

func main() {
    sum := 0
    if err := readInput(&sum); err != nil {
        log.Fatalf("%v\n", err)
    }
    fmt.Printf("Sum: %v | Iterations: %v\n", sum, iterations)
}

func stalkTheGuard(lab map[positions]string, guard, prevPos, rowLimit positions, currentDir string, visited map[positions]positions) map[positions]positions {
    iterations += 1
    if guard.i != prevPos.i || guard.j != prevPos.j {
        if guard.i >= 0 && guard.i <= rowLimit.i && guard.j >= 0 && guard.j <= rowLimit.j {
            prevPos = guard
            visited[guard] = guard
            if currentDir == "^" {
                if _, ok := lab[guard]; ok {
                    delete(visited, guard)
                    guard = positions{i: guard.i+1, j: guard.j+1}
                    currentDir = ">"
                } else {
                    guard = positions{i: guard.i-1, j: guard.j}
                }
                return stalkTheGuard(lab, guard, prevPos, rowLimit, currentDir, visited)
            }
            if currentDir == ">" {
                if _, ok := lab[guard]; ok {
                    delete(visited, guard)
                    guard = positions{i: guard.i+1, j: guard.j-1}
                    currentDir = "v"
                } else {
                    guard = positions{i: guard.i, j: guard.j+1}
                }
                return stalkTheGuard(lab, guard, prevPos, rowLimit, currentDir, visited)
            }
            if currentDir == "v" {
                if _, ok := lab[guard]; ok {
                    delete(visited, guard)
                    guard = positions{i: guard.i-1, j: guard.j-1}
                    currentDir = "<"
                } else {
                    guard = positions{i: guard.i+1, j: guard.j}
                }
                return stalkTheGuard(lab, guard, prevPos, rowLimit, currentDir, visited)
            }
            if currentDir == "<" {
                if _, ok := lab[guard]; ok {
                    delete(visited, guard)
                    guard = positions{i: guard.i-1, j: guard.j+1}
                    currentDir = "^"
                } else {
                    guard = positions{i: guard.i, j: guard.j-1}
                }
                return stalkTheGuard(lab, guard, prevPos, rowLimit, currentDir, visited)
            }
        }
    }

    return visited
}

func readInput(sum *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    
    lab := make(map[positions]string)
    guard := positions{}
    k := 0
    guardInitialDir := ""
    row := 0

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        for index, ele := range line {
            iterations += 1
            item := string(ele)
            if item == "#" {
                lab[positions{i: k, j: index}] = "#"
            }
            if item == "^" || item == ">" || item == "v" || item == "<" {
                guard = positions{i: k, j: index}    
                guardInitialDir = item   
            }
        }
        k += 1
        row = len(line)
    }
    *sum += len(stalkTheGuard(lab, guard, positions{i: -1, j: -1}, positions{i: k-1, j: row-1}, guardInitialDir, make(map[positions]positions)))

    return nil
}
