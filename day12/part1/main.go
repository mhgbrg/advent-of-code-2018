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

	generations := 20

	state := initState
	for g := 1; g <= generations; g++ {
		state = ".." + state + ".."
	}

	fmt.Println(0, state)

	for g := 1; g <= generations; g++ {
		state = mutate(state, combos)
		fmt.Println(g, state)
	}

	result := 0
	for i := 0; i < len(state); i++ {
		if state[i] == '#' {
			result += i - generations*2
		}
	}

	fmt.Println(result)

	fmt.Println(time.Since(start))
}

func mutate(state string, combos map[string]rune) string {
	newState := []rune(state)

	for i := 2; i < len(state)-2; i++ {
		slice := state[i-2 : i+3]
		newState[i] = combos[slice]
	}

	// TODO: Keep as rune slice
	return string(newState)
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
