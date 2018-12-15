package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	serial := readInput()
	gridSize := 300
	squareSize := 3

	maxPowerLevel := -math.MaxInt32
	var bestX, bestY int
	for x := 1; x <= gridSize-squareSize-1; x++ {
		for y := 1; y <= gridSize-squareSize-1; y++ {
			powerLevel := squarePowerLevel(x, y, squareSize, serial)
			if powerLevel > maxPowerLevel {
				maxPowerLevel = powerLevel
				bestX = x
				bestY = y
			}
		}
	}

	fmt.Printf("(%d,%d): %d\n", bestX, bestY, maxPowerLevel)

	fmt.Println(time.Since(start))
}

func squarePowerLevel(x, y, squareSize, serial int) int {
	powerLevel := 0
	for dx := 0; dx < squareSize; dx++ {
		for dy := 0; dy < squareSize; dy++ {
			powerLevel += cellPowerLevel(x+dx, y+dy, serial)
		}
	}
	return powerLevel
}

func cellPowerLevel(x, y, serial int) int {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10
	// Begin with a power level of the rack ID times the Y coordinate.
	powerLevel := rackID * y
	// Increase the power level by the value of the grid serial number (your puzzle input).
	powerLevel += serial
	// Set the power level to itself multiplied by the rack ID.
	powerLevel *= rackID
	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	powerLevel = powerLevel / 100 % 10
	// Subtract 5 from the power level.
	return powerLevel - 5
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
