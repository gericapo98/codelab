package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func nextValue(sequence []int) int {
	rows := [][]int{sequence}

	for {

		current := rows[len(rows)-1]
		if len(current) == 1 {
			break
		}

		next := make([]int, len(current)-1)
		for i := 0; i < len(current)-1; i++ {
			next[i] = current[i+1] - current[i]
		}

		rows = append(rows, next)
		if allZeros(next) {
			break
		}
	}


	for i := len(rows) - 1; i > 0; i-- {
		rows[i] = append(rows[i], 0)
		for j := len(rows[i]) - 1; j > 0; j-- {
			rows[i-1][j] += rows[i][j]
		}
	}

	return rows[0][len(rows[0])-1]
}


func allZeros(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		sequence := make([]int, len(parts))
		for i, part := range parts {
			sequence[i], _ = strconv.Atoi(part)
		}


		totalSum += nextValue(sequence)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	fmt.Println("Sum of extrapolated values:", totalSum)
}
