package main

import(
  "bufio"
  "fmt"
  "os"
  "unicode"
)

func extractDigits(line string) (int, int){
  var firstDigit, lastDigit rune
  var first bool
  first := true
  for _, ch := range line {
    if first == true {
      firstDigit = ch
    } else {
      lastDigit = ch 
      first = true
    } 
  }
  return int(firstDigit - '0'), int(lastDigit - '0')
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
  var l, r, sum int
  sum := 0
  for scanner.Scan(){
    l,r = extractDigits(scanner.Text())
    sum = sum + l + r
    lines = append(lines,scanner.Text())
  }
  
  if err := scanner.Err(); err != nil{
    fmt.Println("Error reading file:", err)
    return
  }
  
  fmt.Println("The total distance between left list and right list is:", sum)
  
}


