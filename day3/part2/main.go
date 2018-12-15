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

	for _, c := range claims {
		if !isOverlapped(c, claims) {
			fmt.Println(c.id)
			break
		}
	}
}

func isOverlapped(c claim, claims []claim) bool {
	for _, c2 := range claims {
		if c == c2 {
			continue
		}
		if c.overlaps(c2) {
			return true
		}
	}
	return false
}

func (c claim) overlaps(c2 claim) bool {
	if c.x > c2.xMax() || c2.x > c.xMax() {
		return false
	}

	if c.y > c2.yMax() || c2.y > c.yMax() {
		return false
	}

	return true
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
	x := unsafeAtoi(pos[0]) + 1
	y := unsafeAtoi(pos[1][:len(pos[1])-1]) + 1
	dim := strings.Split(parts[3], "x")
	w := unsafeAtoi(dim[0])
	h := unsafeAtoi(dim[1])
	return claim{id, x, y, w, h}
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
