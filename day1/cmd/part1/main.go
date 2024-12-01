package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
    var (
        list1 []int
        list2 []int
        sum int
    )

    readInput(&list1, &list2)
    
    sort.Ints(list1)
    sort.Ints(list2)

    for i := 0; i < len(list1); i++ {
        diff := largerofTwo(list1[i], list2[i]) 
        sum += diff
    }    
    fmt.Println(sum) 
}

func largerofTwo(a, b int) int {
    if a > b {
        return a + (^b + 1)
    }
    return b + (^a + 1)
}

func readInput(list1, list2 *[]int) error {
    f, err := os.Open("input.txt")
    if err != nil {
        return err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        subStr := strings.Split(line, "   ")

        num1, err := strconv.Atoi(subStr[0])
        if err != nil {
            log.Fatalf("failed converting string to int: %v\n", err)
        }
        *list1 = append(*list1, num1)

        num2, err := strconv.Atoi(subStr[1])
        if err != nil {
            log.Fatalf("failed converting string to int: %v\n", err)
        }
        *list2 = append(*list2, num2)
    }

    return nil
}
