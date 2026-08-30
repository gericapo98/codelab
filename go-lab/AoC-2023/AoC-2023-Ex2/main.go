/* a list of games are given in the input, each of which contains 
* an id, and three rounds for the game consisting of sets of cubes of up to three different colors.
* the game is impossible when : for any of the colors, if the set of the cubes for that color 
* has a cardinality higher than the number initially given.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mred, mgreen, mblue := 12, 13, 14

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfValidGameIDs := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Split game ID and cube sets
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			fmt.Printf("Malformed line: %s\n", line)
			continue
		}

		gameID, _ := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Game")))
		cubeSets := strings.Split(parts[1], "; ")

		valid := true
		for _, set := range cubeSets {
			cubes := strings.Split(set, ", ")
			red, green, blue := 0, 0, 0

			for _, cube := range cubes {
				parts := strings.Fields(cube)
				if len(parts) != 2 {
					fmt.Printf("Unexpected cube format: %s\n", cube)
					continue
				}
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]

				switch color {
				case "red":
					red += count
				case "green":
					green += count
				case "blue":
					blue += count
				default:
					fmt.Printf("Unknown color: %s\n", color)
				}
			}
			if red > mred || green > mgreen || blue > mblue {
				valid = false
				break
			}
		}
		if valid {
			sumOfValidGameIDs += gameID
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum of valid game IDs:", sumOfValidGameIDs)
}

