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

type memPos struct {
    order int
    direction string
}

type watcher struct {
    direction string
    pos positions
    past map[positions]memPos
}

var iterations int

func main() {
    sum := 0
    if err := readInput(&sum); err != nil {
        log.Fatalf("%v\n", err)
    }
    fmt.Printf("Sum: %v | Iterations: %v\n", sum, iterations)
}

func boundsCheck(guard watcher, bounds positions) bool {
    if !(guard.pos.i >= 0 && guard.pos.i <= bounds.i) || 
        !(guard.pos.j >= 0 && guard.pos.j <= bounds.j) {
        return false
    }
    return true
}

func stalkTheGuard(guard watcher, lab map[positions]string, bounds positions, order int) {
    for boundsCheck(guard, bounds) {
        iterations += 1
        if _, ok := guard.past[guard.pos]; !ok {
            guard.past[guard.pos] = memPos{order: order, direction: guard.direction}
            order += 1
        }
        walkThePath(&guard, lab)
    }
}

func trapTheGuard(guard watcher, lab map[positions]string, bounds positions) bool {
    currentLoop := make(map[positions]string)
    for boundsCheck(guard, bounds) {
        iterations += 1
        if value, ok := currentLoop[guard.pos]; ok {
            if value == guard.direction {
                return true
            }
        } else {
            currentLoop[guard.pos] = guard.direction
        }
        walkThePath(&guard, lab)
    }
    return false
}

func walkThePath(guard *watcher, lab map[positions]string) {
    switch guard.direction {
    case "^":
        if checkCollision(watcher{direction: guard.direction, pos: positions{i: guard.pos.i-1, j: guard.pos.j}}, lab) {
            guard.pos.i -= 1
        } else {
            guard.direction = ">"
        }
    case ">":
        if checkCollision(watcher{direction: guard.direction, pos: positions{i: guard.pos.i, j: guard.pos.j+1}}, lab) {
            guard.pos.j += 1
        } else {
            guard.direction = "v"
        }
    case "v":
        if checkCollision(watcher{direction: guard.direction, pos: positions{i: guard.pos.i+1, j: guard.pos.j}}, lab) {
            guard.pos.i += 1
        } else {
            guard.direction = "<"
        }
    case "<":
        if checkCollision(watcher{direction: guard.direction, pos: positions{i: guard.pos.i, j: guard.pos.j-1}}, lab) {
            guard.pos.j -= 1
        } else {
            guard.direction = "^"
        }
    }
}

func checkCollision(guard watcher, lab map[positions]string) bool {
    if value, ok := lab[guard.pos]; ok {
        if value == "#" {
            return false
        }
    }
    return true
}

func readInput(sum *int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    
    lab := make(map[positions]string)
    guard := watcher{}
    k := 0
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
                guard = watcher{direction: item, pos: positions{i: k, j: index}, past: make(map[positions]memPos)} 
            }
        }
        k += 1
        row = len(line)
    }
    
    guardInitialLife := guard

    stalkTheGuard(guard, lab, positions{i: k-1, j: row-1}, 0)

    guardsOldLife := make(map[int]positions)
    for k, i := range guard.past {
        guardsOldLife[i.order] = k
    }
   
    past := guard.past
    guard = guardInitialLife
    guard.past = past
    i := 2

    for i <= len(guardsOldLife) {
        if _, ok := guardsOldLife[i]; ok {
            lab[positions{i: guardsOldLife[i].i, j: guardsOldLife[i].j}] = "#"
        }
        if trapTheGuard(guard, lab, positions{i: k-1, j: row-1}) {
            *sum += 1
        }
        delete(lab, positions{i: guardsOldLife[i].i, j: guardsOldLife[i].j})
        i += 1
    }

    return nil
}

