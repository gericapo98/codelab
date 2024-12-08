package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	var times, distances []int
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "Time:") {
			timeParts := strings.Fields(line[5:])
			for _, part := range timeParts {
				time, _ := strconv.Atoi(part)
				times = append(times, time)
			}
		} else if strings.HasPrefix(line, "Distance:") {
			distanceParts := strings.Fields(line[9:])
			for _, part := range distanceParts {
				distance, _ := strconv.Atoi(part)
				distances = append(distances, distance)
			}
		}
	}

	return times, distances
}

func calculateWaysToWin(times []int, distances []int) int {
	totalWays := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		record := distances[i]

		ways := 0
		for hold := 1; hold <= time; hold++ {
			speed := hold
			travelTime := time - hold
			distance := speed * travelTime

			if distance > record {
				ways++
			}
		}

		totalWays *= ways
	}

	return totalWays
}

func main() {
	times, distances := parseInput("input.txt")

	if len(times) == 0 || len(distances) == 0 {
		fmt.Println("Invalid input.")
		return
	}


	result := calculateWaysToWin(times, distances)
	fmt.Println("The total number of ways to win:", result)
}

