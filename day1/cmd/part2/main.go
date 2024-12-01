package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type list_map struct {
    Number int
    Count int
}

func main() {
    var list1 []int
    list2 := make(map[int]list_map)
    var sum int

    readInput(&list1, &list2)

    for _, i := range list1 {
        if v, ok := list2[i]; ok {
            sum += similiarity(i, v.Count)
        }
    }
    fmt.Println(sum)
}

// iterate list1, look-up in list2, check v.count and return list1[i] * list2.v.count

func similiarity(num1, num2 int) int { 
    return num1 * num2
}

func readInput(list1 *[]int, list2 *map[int]list_map) error {
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
        if v, ok := (*list2)[num2]; !ok {
            (*list2)[num2] = list_map{
                Number: num2,
                Count: 1,
            }
        } else {
            (*list2)[num2] = list_map{
                Number: num2,
                Count: v.Count+1,
            }
        }
    }

    return nil
}
