package main

import (
  "bufio"
  "fmt"
  "os"
  "unicode"
)

func extractDigits(line string) (int,int){
    var firstDigit, lastDigit rune
    for _, ch:= range line {
      if unicode.IsDigit(ch){
        if firstDigit == 0 {
          firstDigit = ch        
        }
        lastDigit=ch
      }
    }

    if firstDigit == 0 || lastDigit == 0 {
      fmt.Println("Error: Line does not contain valid digits:", line)
      return 0,0
    } 
  return int(firstDigit - '0'), int(lastDigit - '0')
}

func combineDigits(first, last int) int {
  fmt.Println("Combining digits, resulting to: 10*%d + %d = %d", first, last, 10*first + last)
  return 10*first + last
}

func calculateSum(lines []string) int{
  total := 0
  for _, line := range lines {
    first, last := extractDigits(line)
    if first == 0 && last == 0 {
			fmt.Println("Skipping line with no digits:", line)
			continue
		}
    calibrationValue := combineDigits(first, last)
    total += calibrationValue
  }
  return total
}

func main(){
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return 
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)

  for scanner.Scan(){
    lines = append(lines, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  total := calculateSum(lines)
  fmt.Println("Total calibration value:", total)
}
