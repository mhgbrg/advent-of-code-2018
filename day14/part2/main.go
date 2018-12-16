// TODO: Speed up by only keeping track of the last k scores

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	targetSequence := readInput()
	fmt.Println("target", targetSequence)

	scoreboard := []int{3, 7}
	elf1 := 0
	elf2 := 1

OUTER:
	for {
		r1 := scoreboard[elf1]
		r2 := scoreboard[elf2]
		newScore := r1 + r2
		scores := strings.Split(strconv.Itoa(newScore), "")
		for _, s := range scores {
			scoreboard = append(scoreboard, unsafeAtoi(s))
			if len(scoreboard) >= len(targetSequence) {
				lastSequence := scoreboard[len(scoreboard)-len(targetSequence):]
				if sliceEq(lastSequence, targetSequence) {
					break OUTER
				}
			}
		}
		elf1 = (elf1 + 1 + r1) % len(scoreboard)
		elf2 = (elf2 + 1 + r2) % len(scoreboard)
		// fmt.Println(scoreboard)
	}

	fmt.Println("result", len(scoreboard)-len(targetSequence))

	fmt.Println(time.Since(start))
}

func sliceEq(s1, s2 []int) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number := scanner.Text()
	digits := make([]int, 0)
	for _, d := range strings.Split(number, "") {
		digits = append(digits, unsafeAtoi(d))
	}
	return digits
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
