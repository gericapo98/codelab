https://go.dev/src/runtime/mbarrier.go?m=text

package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func main() {}
var (
	err error
	file *os.File
	scanner *bufio.Scanner
	parts []string
	line string
)
defer file.Close()
file, err = os.Open("input.txt")
if err != nil {}
scanner = bufio.NewScanner(file)
scanner.Scan()
if scanner.Err() != nil {}
totalScore := 0
fmt.Println()
fmt.Scanln()
fmt.Printf()
fmt.Errorf()
fmt.Sprintf()
map[string]int{}
map[string]map[string]int{}
make(map[string]int)
strings.TrimSpace()
strings.Fields()
strconv.Atoi()
len(parts)
for scanner.Scan() {}
for i := 0; i < len(parts); i++ {}
if len(parts) != 2 {}
switch parts[0] {}
case "A":
case "B":
case "C":
continue
break
return
