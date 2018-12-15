package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()
	points := readInput()

	gridSize := 400
	safeSize := 0

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			c := coord{x, y}
			if safe(c, points) {
				safeSize++
			}
		}
	}

	fmt.Println(safeSize)

	fmt.Println(time.Since(start))
}

func safe(c coord, points []coord) bool {
	limit := 10000
	totalDistance := 0
	for _, p := range points {
		totalDistance += distance(c, p)
		if totalDistance >= limit {
			return false
		}
	}
	return totalDistance < limit
}

func distance(c1, c2 coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(m map[int]int) (int, int) {
	mv := math.MaxInt32
	var mk int

	for k, v := range m {
		if v < mv {
			mv = v
			mk = k
		}
	}

	for k, v := range m {
		if k == mk {
			continue
		} else if v == mv {
			return -1, 0
		}
	}

	return mk, mv
}

func max(m map[int]int) (int, int) {
	mv := 0
	var mk int

	for k, v := range m {
		if v > mv && v < math.MaxInt32 {
			mv = v
			mk = k
		}
	}

	return mk, mv
}

func readInput() []coord {
	var input []coord
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		c := strings.Split(line, ", ")
		x := unsafeAtoi(c[0])
		y := unsafeAtoi(c[1])
		input = append(input, coord{x, y})
	}
	return input
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
