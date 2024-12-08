package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

func calculateMulSum(input string) int {
  pattern := `mul\((\d{1,3}),(\d{1,3})\)`
  re := regexp.MustCompile(pattern)

  matches := re.FindAllStringSubmatch(input, -1)
  
  totalSum := 0
  for _, match := range matches {
    X, _ := strconv.Atoi(match[1])
    Y, _ := strconv.Atoi(match[2])
    totalSum += x*y
  }
  return totalSum
}

func main(){
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  var input string
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    input += scanner.Text()
  }
  result := calculateMulSum(input)
	fmt.Println("The total sum of valid multiplications is:", result)
}
