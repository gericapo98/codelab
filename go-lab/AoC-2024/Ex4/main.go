package main

// Count occurrences of a given word ("XMAS") in a grid.
// Check horizontally, vertically, and diagonally in all directions.
//

import (
	"bufio"
	"fmt"
	"os"
  //"strings"
)

var directions = [8][2]int{
	{0, 1}, {1, 0}, 
	{0, -1},{-1, 0}, 
	{1, 1}, {-1, -1},
  {1, -1}, {-1, 1}, 
}

func countWord(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])

	wordLen := len(word)
	totalCount := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				count := 0

				for i := 0; i < wordLen; i++ {
					nr, nc := r+i*dir[0], c+i*dir[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == rune(word[i]) {
						count++
					} else {
						break
					}
				}
				if count == wordLen {
					totalCount++
				}
			}
		}
	}
	return totalCount
}

func readGridFromFile(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	  	grid = append(grid, []rune(line))
	}

	if  err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func main() {
	filename := "input.txt"
	grid := readGridFromFile(filename)
	word := "XMAS"
	result := countWord(grid, word)
	fmt.Printf("The word '%s' appears %d times in the grid.\n", word, result)
}

