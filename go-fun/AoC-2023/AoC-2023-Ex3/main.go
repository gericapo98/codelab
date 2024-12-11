package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var schematic [][]rune
	scanner := bufio.NewScanner(file)

	
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := len(schematic)
	cols := len(schematic[0])
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalSum := 0

	
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; {
			
			if unicode.IsDigit(schematic[r][c]) {
				start := c
				for c < cols && unicode.IsDigit(schematic[r][c]) {
					c++
				}


				num, _ := strconv.Atoi(string(schematic[r][start:c]))
				isPartNumber := false


				for i := start; i < c; i++ {
					for _, dir := range directions {
						nr, nc := r+dir[0], i+dir[1]
						if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
							if schematic[nr][nc] != '.' && !unicode.IsDigit(schematic[nr][nc]) {
								isPartNumber = true
								break
							}
						}
					}
					if isPartNumber {
						break
					}
				}

				if isPartNumber {
					totalSum += num
				}
			} else {
				c++
			}
		}
	}

	fmt.Println("Sum of all part numbers:", totalSum)
}

