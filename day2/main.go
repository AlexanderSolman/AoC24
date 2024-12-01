package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
    var list1 []int
    var list2 []int

    readInput(&list1, &list2)

    x := 0
    for _, i := range list1 {
        fmt.Println(i)
        x += 1
        if x == 10 { break }
    }
    x = 0
    fmt.Println("-----------")
    for _, i := range list2 {
        fmt.Println(i)
        x += 1
        if x == 10 { break }
    }
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
