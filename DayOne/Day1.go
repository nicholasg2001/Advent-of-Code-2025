package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solve(rotation string, ans int, position int) (int, int) {

	rotation = strings.TrimSpace(rotation)
	direction := 1

	if rotation[0] == 'L' {
		direction = -1
	}

	spacesToMove, _ := strconv.Atoi(rotation[1:])
	spacesToMove *= direction

	position = (((position+spacesToMove)%100 + 100) % 100)

	if position == 0 {
		ans += 1
	}

	return ans, position
}

func main() {

	ans := 0
	position := 50

	file, err := os.Open("../Inputs/inputDay1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() //close file after all actions are completed

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans, position = solve(scanner.Text(), ans, position)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
