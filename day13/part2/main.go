package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

const (
	leftTurn int = iota
	noTurn
	rightTurn
)

type pos struct {
	x int
	y int
}

func (p pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type cart struct {
	p         pos
	direction rune
	nextTurn  int
}

func (c *cart) move(trackMap [][]rune) {
	nextX, nextY := c.p.x, c.p.y
	switch c.direction {
	case '^':
		nextY--
	case 'v':
		nextY++
	case '<':
		nextX--
	case '>':
		nextX++
	}
	c.p = pos{nextX, nextY}
	next := trackMap[nextY][nextX]
	switch next {
	case '/':
		switch c.direction {
		case '^':
			c.turnRight()
		case 'v':
			c.turnRight()
		case '<':
			c.turnLeft()
		case '>':
			c.turnLeft()
		}
	case '\\':
		switch c.direction {
		case '^':
			c.turnLeft()
		case 'v':
			c.turnLeft()
		case '<':
			c.turnRight()
		case '>':
			c.turnRight()
		}
	case '+':
		switch c.nextTurn {
		case leftTurn:
			c.turnLeft()
			c.nextTurn = noTurn
		case noTurn:
			c.nextTurn = rightTurn
		case rightTurn:
			c.turnRight()
			c.nextTurn = leftTurn
		}
	}
}

func (c *cart) turnLeft() {
	switch c.direction {
	case '^':
		c.direction = '<'
	case 'v':
		c.direction = '>'
	case '<':
		c.direction = 'v'
	case '>':
		c.direction = '^'
	}
}

func (c *cart) turnRight() {
	switch c.direction {
	case '^':
		c.direction = '>'
	case 'v':
		c.direction = '<'
	case '<':
		c.direction = '^'
	case '>':
		c.direction = 'v'
	}
}

func (c cart) String() string {
	return fmt.Sprintf("%q %s", c.direction, c.p)
}

func main() {
	start := time.Now()

	trackMap, carts := readInput()

	// printMap(trackMap, carts)

	for i := 0; len(carts) > 1; i++ {
		for _, p := range sortCartsKeys(carts) {
			c := carts[p]
			if c == nil {
				continue
			}
			delete(carts, c.p)
			c.move(trackMap)
			if c2 := carts[c.p]; c2 != nil {
				fmt.Printf("crash on %s\n", c.p)
				delete(carts, c.p)
			} else {
				carts[c.p] = c
			}
		}

		// printMap(trackMap, carts)
	}

	for _, c := range carts {
		fmt.Println(c)
	}

	fmt.Println(time.Since(start))
}

func sortCartsKeys(carts map[pos]*cart) []pos {
	keys := make([]pos, 0)
	for p := range carts {
		keys = append(keys, p)
	}
	sort.Slice(keys, func(i int, j int) bool {
		p1, p2 := keys[i], keys[j]
		if p1.y < p2.y {
			return true
		}
		if p1.y == p2.y {
			return p1.x < p2.x
		}
		return false
	})
	return keys
}

func printMap(trackMap [][]rune, carts map[pos]*cart) {
	for y := 0; y < len(trackMap); y++ {
		row := trackMap[y]
		for x := 0; x < len(row); x++ {
			c := carts[pos{x, y}]
			if c != nil {
				fmt.Print(string(c.direction))
			} else {
				fmt.Print(string(row[x]))
			}
		}
		fmt.Println()
	}
}

func readInput() ([][]rune, map[pos]*cart) {
	trackMap := make([][]rune, 0)
	carts := make(map[pos]*cart)
	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		row := []rune(scanner.Text())
		for x := 0; x < len(row); x++ {
			r := row[x]
			if r == '^' || r == 'v' || r == '<' || r == '>' {
				p := pos{x, y}
				carts[p] = &cart{p, r, leftTurn}
				if r == '^' || r == 'v' {
					row[x] = '|'
				} else {
					row[x] = '-'
				}
			}
		}
		trackMap = append(trackMap, row)
	}
	return trackMap, carts
}
