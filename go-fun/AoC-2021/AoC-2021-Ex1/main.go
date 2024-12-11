package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var prevDepth, currentDepth, count int
    prevDepth = -1

    for scanner.Scan() {
        line := scanner.Text()
        currentDepth, err = strconv.Atoi(line)
        if err != nil {
            fmt.Println("Error converting line to int:", err)
            continue
        }

        if prevDepth != -1 && currentDepth > prevDepth {
            count++
        }
        prevDepth = currentDepth
    }

    
    fmt.Println("The amount of measurement increases is:", count)
}

