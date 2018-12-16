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

	numberOfRecipes := readInput()
	fmt.Println(numberOfRecipes)

	scoreboard := []int{3, 7}
	elf1 := 0
	elf2 := 1

	for len(scoreboard) < numberOfRecipes+10 {
		r1 := scoreboard[elf1]
		r2 := scoreboard[elf2]
		newScore := r1 + r2
		scores := strings.Split(strconv.Itoa(newScore), "")
		for _, s := range scores {
			scoreboard = append(scoreboard, unsafeAtoi(s))
		}
		elf1 = (elf1 + 1 + r1) % len(scoreboard)
		elf2 = (elf2 + 1 + r2) % len(scoreboard)
		// fmt.Println(scoreboard)
	}

	fmt.Println(scoreboard[numberOfRecipes : numberOfRecipes+10])

	fmt.Println(time.Since(start))
}

func readInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return unsafeAtoi(scanner.Text())
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
