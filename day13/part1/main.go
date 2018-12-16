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

type cart struct {
	x         int
	y         int
	direction rune
	nextTurn  int
}

func (c *cart) move(trackMap [][]rune) {
	nextX, nextY := c.x, c.y
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
	c.x = nextX
	c.y = nextY
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
	return fmt.Sprintf("%q (%d, %d)", c.direction, c.x, c.y)
}

func main() {
	start := time.Now()

	trackMap, carts := readInput()

	// printMap(trackMap, carts)

OUTER:
	for i := 0; ; i++ {
		sortCarts(carts)
		for _, c := range carts {
			c.move(trackMap)
			if crash, cx, cy := checkForCrash(carts); crash {
				fmt.Printf("crash on (%d, %d)\n", cx, cy)
				break OUTER
			}
		}

		// printMap(trackMap, carts)
	}

	fmt.Println(time.Since(start))
}

func sortCarts(carts []*cart) {
	sort.Slice(carts, func(i, j int) bool {
		c1, c2 := carts[i], carts[j]
		if c1.y < c2.y {
			return true
		}
		if c1.y == c2.y {
			return c1.x < c2.x
		}
		return false
	})
}

func checkForCrash(carts []*cart) (bool, int, int) {
	for _, c1 := range carts {
		for _, c2 := range carts {
			if c1 == c2 {
				continue
			}
			if c1.x == c2.x && c1.y == c2.y {
				return true, c1.x, c1.y
			}
		}
	}
	return false, 0, 0
}

func printMap(trackMap [][]rune, carts []*cart) {
	for y := 0; y < len(trackMap); y++ {
		row := trackMap[y]
		for x := 0; x < len(row); x++ {
			cartsOnPos := make([]*cart, 0)
			for _, c := range carts {
				if c.x == x && c.y == y {
					cartsOnPos = append(cartsOnPos, c)
				}
			}

			if len(cartsOnPos) == 1 {
				fmt.Print(string(cartsOnPos[0].direction))
			} else if len(cartsOnPos) > 1 {
				fmt.Print("X")
			} else {
				fmt.Print(string(row[x]))
			}
		}
		fmt.Println()
	}
}

func readInput() ([][]rune, []*cart) {
	trackMap := make([][]rune, 0)
	carts := make([]*cart, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		row := []rune(scanner.Text())
		for x := 0; x < len(row); x++ {
			r := row[x]
			if r == '^' || r == 'v' || r == '<' || r == '>' {
				carts = append(carts, &cart{x, y, r, leftTurn})
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
