package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

var directions = map[rune]Position{
	'|':  {0, 1},
	'-':  {1, 0},
	'L':  {-1, 1},
	'J':  {-1, -1},
	'7':  {1, -1},
	'F':  {1, 1},
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func findStart(grid [][]rune) Position {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				return Position{x, y}
			}
		}
	}
	return Position{-1, -1}
}

func isValidMove(grid [][]rune, pos Position, from Position) bool {
	y, x := pos.y, pos.x
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) || grid[y][x] == '.' {
		return false
	}
	if grid[y][x] == 'S' && from != (Position{-1, -1}) {
		return true
	}
	return true
}

func dfs(grid [][]rune, pos Position, from Position, visited map[Position]bool, distance int, maxDist *int) {
	visited[pos] = true
	if distance > *maxDist {
		*maxDist = distance
	}

	for _, d := range directions {
		next := Position{pos.x + d.x, pos.y + d.y}
		if !visited[next] && isValidMove(grid, next, pos) {
			dfs(grid, next, pos, visited, distance+1, maxDist)
		}
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	grid := parseGrid(string(input))
	start := findStart(grid)
	if start == (Position{-1, -1}) {
		fmt.Println("Start position 'S' not found in the grid")
		return
	}

	visited := make(map[Position]bool)
	maxDist := 0
	dfs(grid, start, Position{-1, -1}, visited, 0, &maxDist)

	fmt.Println("Maximum distance from 'S':", maxDist)
}
