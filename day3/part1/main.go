package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func (c claim) xMax() int {
	return c.x + c.w - 1
}

func (c claim) yMax() int {
	return c.y + c.h - 1
}

func main() {
	claims := readInput()

	fabricSize := 1000
	overlapCount := 0

	for x := 0; x < fabricSize; x++ {
		for y := 0; y < fabricSize; y++ {
			if hasOverlap(x, y, claims) {
				overlapCount++
			}
		}
	}

	fmt.Println(overlapCount)
}

func hasOverlap(x, y int, claims []claim) bool {
	covered := false
	for _, c := range claims {
		if c.covers(x, y) {
			if covered {
				return true
			}
			covered = true
		}
	}
	return false
}

func (c claim) covers(x, y int) bool {
	return isBetween(x, c.x, c.xMax()) && isBetween(y, c.y, c.yMax())
}

func isBetween(i, min, max int) bool {
	return min <= i && i <= max
}

func readInput() []claim {
	var claims []claim
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		c := parseLine(line)
		claims = append(claims, c)
	}
	return claims
}

func parseLine(line string) claim {
	parts := strings.Split(line, " ")
	id := unsafeAtoi(parts[0][1:])
	pos := strings.Split(parts[2], ",")
	x := unsafeAtoi(pos[0])
	y := unsafeAtoi(pos[1][:len(pos[1])-1])
	dim := strings.Split(parts[3], "x")
	w := unsafeAtoi(dim[0])
	h := unsafeAtoi(dim[1])
	return claim{id, x, y, w, h}
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
