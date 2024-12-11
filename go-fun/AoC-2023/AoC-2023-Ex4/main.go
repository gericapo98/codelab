package main

import (
  "bufio"
  "fmt"
  "os"
 //"strconv"
  "strings"
)

func main(){
  file, err := os.Open("input.txt")
  if err != nil{
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  totalScore := 0
  scanner := bufio.NewScanner(file)

  for scanner.Scan(){
    line := scanner.Text()

    parts := strings.Split(line, " | ")
    if len(parts) != 2 {
      fmt.Println("Malformed input:", line)
      continue
    }
    
    winningNumbers := strings.Fields(parts[0])
    numbersYouHave := strings.Fields(parts[1])
    
    winningSet := make(map[string]bool)
    for i, num := range winningNumbers {
      fmt.Printf("%v\n", winningNumbers[i])
      winningSet[num] = true
    }
        
    cardScore, matchCount := 0,0
    for _, num := range numbersYouHave {
      if winningSet[num] {
        matchCount++
        if matchCount == 1 {
          cardScore += 1
        } else {
          cardScore *= 2
        }
      }
    }
    totalScore += cardScore
  }
  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading input:", err)
    return
  }
  fmt.Println("Total score:", totalScore)
}
