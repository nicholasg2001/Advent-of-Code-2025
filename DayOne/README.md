## Problem: 

You arrive at the secret entrance to the North Pole base ready to start decorating. Unfortunately, the password seems to have been changed, so you can't get in. A document taped to the wall helpfully explains:

"Due to new security protocols, the password is locked in the safe below. Please see the attached document for the new combination."

The safe has a dial with only an arrow on it; around the dial are the numbers 0 through 99 in order. As you turn the dial, it makes a small click noise as it reaches each number.

The attached document (your puzzle input) contains a sequence of rotations, one per line, which tell you how to open the safe. A rotation starts with an L or R which indicates whether the rotation should be to the left (toward lower numbers) or to the right (toward higher numbers). Then, the rotation has a distance value which indicates how many clicks the dial should be rotated in that direction.

So, if the dial were pointing at 11, a rotation of R8 would cause the dial to point at 19. After that, a rotation of L19 would cause it to point at 0.

Because the dial is a circle, turning the dial left from 0 one click makes it point at 99. Similarly, turning the dial right from 99 one click makes it point at 0.

So, if the dial were pointing at 5, a rotation of L10 would cause it to point at 95. After that, a rotation of R5 could cause it to point at 0.

The dial starts by pointing at 50.

You could follow the instructions, but your recent required official North Pole secret entrance security training seminar taught you that the safe is actually a decoy. The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

For example, suppose the attached document contained the following rotations:

L68
L30
R48
L5
R60
L55
L1
L99
R14
L82

Following these rotations would cause the dial to move as follows:

    The dial starts by pointing at 50.
    The dial is rotated L68 to point at 82.
    The dial is rotated L30 to point at 52.
    The dial is rotated R48 to point at 0.
    The dial is rotated L5 to point at 95.
    The dial is rotated R60 to point at 55.
    The dial is rotated L55 to point at 0.
    The dial is rotated L1 to point at 99.
    The dial is rotated L99 to point at 0.
    The dial is rotated R14 to point at 14.
    The dial is rotated L82 to point at 32.

Because the dial points at 0 a total of three times during this process, the password in this example is 3.

Analyze the rotations in your attached document. What's the actual password to open the door?




## Original solution:

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

