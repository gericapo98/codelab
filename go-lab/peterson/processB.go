package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const sharedFile = "shared_memory.txt"

func main() {
	for {
		updateSharedFile(true, 0)
		for {
			flagOne, flagTwo, turn := readSharedFile()
			if !flagOne || turn != 0 {
				break
			}
			fmt.Println("Process B: Waiting...")
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Process B: Entering critical section")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Process B: Leaving critical section")
		updateSharedFile(false, 0)
		time.Sleep(500 * time.Millisecond)
	}
}

func readSharedFile() (bool, bool, int) {
	data, _ := os.ReadFile(sharedFile)
	parts := string(data)
	flagOne, _ := strconv.ParseBool(string(parts[0]))
	flagTwo, _ := strconv.ParseBool(string(parts[2]))
	turn, _ := strconv.Atoi(string(parts[4]))
	return flagOne, flagTwo, turn
}

func updateSharedFile(flagTwo bool, turn int) {
	data := fmt.Sprintf("%t,%t,%d", true, flagTwo, turn)
	os.WriteFile(sharedFile, []byte(data), 0644)
}

