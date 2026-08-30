package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var leftList, rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text()) 
		if len(line) == 2 {
			left, _ := strconv.Atoi(line[0]) 
			right, _ := strconv.Atoi(line[1]) 
			leftList = append(leftList, left)
			rightList = append(rightList, right)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	var totalDistance int
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
	}

	fmt.Println("The total distance between left and right lists is:", totalDistance)
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

