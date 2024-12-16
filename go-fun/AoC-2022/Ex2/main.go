package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  //"strconv"
)
/*
* simple logic:
* go through each line and check the first strings
* keep score variable to accumulate and add on it 
* depending on the game results.
* print it
* */
func main(){

  file, err := os.Open("input.txt")
  if err != nil{
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  //game := make([]strings, 2)
  moveScore := map[string]int{"X":1, "Y":2, "Z":3}
  gameScore := map[string]map[string]int{
    "A": {"X": 3, "Y": 6, "Z": 0}, 
		"B": {"X": 0, "Y": 3, "Z": 6}, 
		"C": {"X": 6, "Y": 0, "Z": 3},
  } // I have to test this...

  score := 0
  scanner := bufio.NewScanner(file)

  for scanner.Scan(){
    line := scanner.Text()
    parts := strings.Fields(line)
    opponent := parts[0]
		player := parts[1]
    score += moveScore[player] + gameScore[opponent][player]
  }

  fmt.Println("Total Score:", score)
}
