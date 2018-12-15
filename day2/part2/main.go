package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ids := getInput()
	for i, id1 := range ids {
		for j, id2 := range ids {
			if i == j {
				continue
			}
			alike := compareIds(id1, id2)
			if alike {
				fmt.Println(id1)
				fmt.Println(id2)
				printCommonChars(id1, id2)
				return
			}
		}
	}
	fmt.Println("no solution!")
}

func compareIds(id1, id2 string) bool {
	strikeOne := false
	for i, char1 := range id1 {
		char2 := rune(id2[i])
		if char1 != char2 {
			if strikeOne {
				return false
			}
			strikeOne = true
		}
	}
	return true
}

func printCommonChars(str1, str2 string) {
	for i, char1 := range str1 {
		char2 := rune(str2[i])
		if char1 == char2 {
			fmt.Printf("%c", char1)
		}
	}
	fmt.Println()
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
