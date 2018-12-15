package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ids := getInput()
	twos, threes := 0, 0
	for _, id := range ids {
		charCounts := make(map[rune]int, len(id))
		for _, char := range id {
			charCounts[char]++
		}
		containsTwo, containsThree := false, false
		for _, count := range charCounts {
			if count == 2 && !containsTwo {
				twos++
				containsTwo = true
			} else if count == 3 && !containsThree {
				threes++
				containsThree = true
			}
		}
	}
	result := twos * threes
	fmt.Println("twos:", twos)
	fmt.Println("threes:", threes)
	fmt.Println("result:", result)
}

func getInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
