## [Problem Day 1](https://adventofcode.com/2025/day/1)



## Original solution for part 1:

```
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(rotation string, ans [2]int) [2]int {

	move := strings.Split(rotation, "") // Correction
	direction := 1
	spacesToMoveString := ""
	current := ans[1]

	if move[0] == "L" {
		direction = -1
	} else { // Not needed because direction is initialized to 1 anyway
		direction = 1
	}

	spacesToMoveString = strings.Join(move[1:], "")
	spacesToMove, _ := strconv.Atoi(spacesToMoveString)
	spacesToMove *= direction

	ans[1] = (((current+spacesToMove)%100 + 100) % 100)

	if ans[1] == 0 {
		ans[0] += 1
		return ans
	} else {
		return ans
	}

}

func main() {

	ans := [2]int{0, 50}

	file, err := os.Open("../Inputs/inputDay1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() //close file after all actions are completed

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := solve(line, ans)
		ans[0] = res[0]
		ans[1] = res[1]
	}
	fmt.Println(ans[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

```

#### Improvements I learned with our dear friend Claude:
- Use ```TrimSpace``` instead of ```Split``` because Go allows for subtring indexing. ```Trimspace``` will remove the ```\n``` and from there the number value is accessible via ```string[1:]```. TLDR Go allows for string indexing.

- You can return multiple individual variables of any types in Go via ```func() (int, string, float64, bool)... ``` 

