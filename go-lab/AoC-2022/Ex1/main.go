// we go through all the lines of the file
// we keep count of each blank line. => bl + 1 = #elves
// simultaneously memorize max val of calories for each of the elves
// return that value. keeping it simple
//
package main

import (

  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {

  file, err:= os.Open("input.txt")
  if err != nil {
    fmt.Println("Error", err)
    return
  }
  defer file.Close()

  scanner:= bufio.NewScanner(file)

  var maxCalories, currentCalories int
  for scanner.Scan(){
    line := scanner.Text()
  
    if line == "" {
      if currentCalories > maxCalories {
        maxCalories = currentCalories
      }
      currentCalories = 0
    }else {
      calories, _ := strconv.Atoi(line)
      currentCalories += calories
    }  
  }

  if currentCalories > maxCalories {
    maxCalories = currentCalories
  }

  fmt.Println("The Elf carrying the most Calories has:", maxCalories)
}
