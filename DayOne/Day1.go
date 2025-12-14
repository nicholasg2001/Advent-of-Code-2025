package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solvePartOne(rotation string, ans int, position int) (int, int) {

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

func solvePartTwo(rotation string, ans int, position int) (int, int) {

	rotation = strings.TrimSpace(rotation)
	direction := 1

	if rotation[0] == 'L' {
		direction = -1
	}

	spacesToMove, _ := strconv.Atoi(rotation[1:])
	spacesToMove *= direction
	absSpacesToMove := abs(spacesToMove)

	initPosition := position
	cycles := 0

	var firstZero int

	if initPosition == 0 {
		firstZero = 100
	} else if direction == 1 {
		firstZero = 100 - initPosition
	} else {
		firstZero = initPosition
	}

	if absSpacesToMove >= firstZero {
		cycles++
		remainingDistance := absSpacesToMove - firstZero
		cycles += remainingDistance / 100
	}

	sum := position + spacesToMove
	position = ((sum%100 + 100) % 100)

	ans += cycles

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
		ans, position = solvePartTwo(scanner.Text(), ans, position)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
