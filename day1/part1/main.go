package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	currentFreq := 0

	for scanner.Scan() {
		str := scanner.Text()
		freqChange, _ := strconv.Atoi(str)
		currentFreq += freqChange
	}

	fmt.Println(currentFreq)
}
