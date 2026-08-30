/* package main

import(
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

// So for each string out of the input.txt file there import (
// a 2d slice, with rows as the mappings and columns as the 
// destination range, source range and range length.
// For starters I need to, 
// 
func parseInput(filename string)([]int, map[string][][3]int){
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return nil, nil
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  seeds := []int{}
  mappings := make(map[string][][3]int)
  

  
}

func main(){
  seeds, mappings := parseInput("input.txt")
  
  categories := []string {
    "seed-to-soil",
    "soil-to-fertilizer",
    "fertilizer-to-water",
    "water-to-light",
    "light-to-temeperature",
    "temperature-to-humidity",
    "humidity-to-location",
  }
  fmt.Println(seeds)
  fmt.Println(mappings)
  /*
  locations := []int{}
  for _, seed := range seeds {
    value := seed
    for _, category := range categories {
      value = mapValue(value, mappings[category])
    }
    locations = append(locations, value)
g  }
  */


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parse the input file to extract seeds and mappings
func parseInput(filename string) ([]int, map[string][][3]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()
  
	scanner := bufio.NewScanner(file)
	seeds := []int{}
	mappings := make(map[string][][3]int)
	var currentMap string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			// Parse seeds
			seedStrs := strings.Fields(line[len("seeds:"):])
			for _, s := range seedStrs {
				seed, _ := strconv.Atoi(s)
				seeds = append(seeds, seed)
			}
		} else if strings.HasSuffix(line, "map:") {
			// start a new map
			currentMap = strings.TrimSuffix(line, " map:")
			mappings[currentMap] = [][3]int{}
		} else if currentMap != "" {
			// parse a map line
			parts := strings.Fields(line)
			if len(parts) == 3 {
				destStart, _ := strconv.Atoi(parts[0])
				srcStart, _ := strconv.Atoi(parts[1])
				length, _ := strconv.Atoi(parts[2])
				mappings[currentMap] = append(mappings[currentMap], [3]int{destStart, srcStart, length})
			}
		}
	}

	return seeds, mappings
}

func mapValue(value int, ranges [][3]int) int {
	for _, r := range ranges {
		destStart, srcStart, length := r[0], r[1], r[2]
		if value >= srcStart && value < srcStart+length {
			return destStart + (value - srcStart)
		}
	}
	return value 
}

func main() {
	
	seeds, mappings := parseInput("input.txt")

	categories := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	locations := []int{}
	for _, seed := range seeds {
		value := seed
		for _, category := range categories {
			value = mapValue(value, mappings[category])
		}
		locations = append(locations, value)
	}

	
	minLocation := locations[0]
	for _, loc := range locations {
		if loc < minLocation {
			minLocation = loc
		}
	}

	fmt.Println("Lowest location:", minLocation)
}


