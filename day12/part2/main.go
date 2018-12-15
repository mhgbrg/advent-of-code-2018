package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	initState, combos := readInput()

	fmt.Println(initState)
	fmt.Println(combos)

	generations := 50000000000

	state := initState
	g := 0
	firstPotNumber := 0
	value := evaluate(state, firstPotNumber)
	var increase int

	fmt.Println(g, firstPotNumber, value, increase, state)

	for g = 1; g <= generations; g++ {
		newState, removedPots := mutate(state, combos)
		firstPotNumber += removedPots
		newValue := evaluate(newState, firstPotNumber)
		increase = newValue - value
		fmt.Println(g, firstPotNumber, newValue, increase, newState)
		value = newValue
		if newState == state {
			break
		}
		state = newState
	}

	result := (generations-g)*increase + value

	fmt.Println(result)

	fmt.Println(time.Since(start))
}

func mutate(state string, combos map[string]rune) (string, int) {
	firstPlantIndex := strings.Index(state, "#")
	lastPlantIndex := strings.LastIndex(state, "#")
	state = "....." + state[firstPlantIndex:lastPlantIndex+1] + "....."

	newState := []rune(state)

	for i := 2; i < len(state)-2; i++ {
		slice := state[i-2 : i+3]
		newState[i] = combos[slice]
	}

	return string(newState), firstPlantIndex - 5
}

func evaluate(state string, firstPotNumber int) int {
	value := 0
	for i := 0; i < len(state); i++ {
		if state[i] == '#' {
			potNumber := firstPotNumber + i
			value += potNumber
		}
	}
	return value
}

func readInput() (string, map[string]rune) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line1 := scanner.Text()
	initState := strings.Split(line1, ": ")[1]
	scanner.Scan() // blank line
	combos := make(map[string]rune)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		combo := parts[0]
		result := rune(parts[1][0])
		combos[combo] = result
	}
	return initState, combos
}
