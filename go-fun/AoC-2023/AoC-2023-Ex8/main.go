package main

import (
	"bufio"
	"fmt"
	"os"
//	"strconv"
	"strings"
)


type Graph map[string][2]string


func ParseInput(filename string) (Graph, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	graph := make(Graph)
	var instructions string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "L") || strings.HasPrefix(line, "R") {
			instructions = line
		} else if line != "" && strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			node := strings.TrimSpace(parts[0])
			children := strings.TrimSpace(parts[1])
			children = strings.Trim(children, "()")
			childNodes := strings.Split(children, ",")
			graph[node] = [2]string{strings.TrimSpace(childNodes[0]), strings.TrimSpace(childNodes[1])}
		}
	}

	return graph, instructions, nil
}


func NavigateGraph(graph Graph, instructions string) int {
	currentNode := "AAA"
	steps := 0
	instructionIndex := 0

	for currentNode != "ZZZ" {
		direction := instructions[instructionIndex]
		if direction == 'L' {
			currentNode = graph[currentNode][0]
		} else {
			currentNode = graph[currentNode][1]
		}
		steps++
		instructionIndex = (instructionIndex + 1) % len(instructions)
	}

	return steps
}

func main() {

	filename := "input.txt"


	graph, instructions, err := ParseInput(filename)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}


	steps := NavigateGraph(graph, instructions)
	fmt.Println("Steps to reach ZZZ:", steps)
}

