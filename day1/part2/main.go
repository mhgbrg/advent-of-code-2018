package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	freqChanges := getInput()
	seenFreqs := make(map[int]bool)
	currentFreq := 0
	for {
		for _, freqChange := range freqChanges {
			seenFreqs[currentFreq] = true
			currentFreq += freqChange
			if seen := seenFreqs[currentFreq]; seen {
				fmt.Println(currentFreq)
				return
			}
		}
	}
}

func getInput() []int {
	var freqChanges []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		freqChange, _ := strconv.Atoi(str)
		freqChanges = append(freqChanges, freqChange)
	}
	return freqChanges
}
