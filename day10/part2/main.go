package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func (p *point) advance() {
	p.x += p.vx
	p.y += p.vy
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
	start := time.Now()

	points := readInput()

	fmt.Printf("t=%d: %d\n", 0, spread(points))

	for t := 1; spread(points) > 0; t++ {
		for i := 0; i < len(points); i++ {
			points[i].advance()
		}
		fmt.Printf("t=%d: %d\n", t, spread(points))
	}

	fmt.Println("final spread", spread(points))

	printPoints(points)

	fmt.Println(time.Since(start))
}

func spread(points []point) int {
	totalDist := 0
	for _, p1 := range points {
		minDist := math.MaxInt32
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}
			dist := distance(p1, p2)
			minDist = min(dist, minDist)
		}
		totalDist += minDist
	}
	return totalDist
}

func distance(p1, p2 point) int {
	xDiff := abs(p1.x - p2.x)
	yDiff := abs(p1.y - p2.y)
	if xDiff < 2 && yDiff < 2 {
		return 0
	}
	return xDiff + yDiff
}

func printPoints(points []point) {
	xMin, xMax, yMin, yMax := ranges(points)

	mat := make([][]rune, yMax-yMin+1)

	for y := yMin; y <= yMax; y++ {
		line := make([]rune, xMax-xMin+1)
		for x := xMin; x <= xMax; x++ {
			line[x-xMin] = ' '
		}
		mat[y-yMin] = line
	}

	for _, p := range points {
		mat[p.y-yMin][p.x-xMin] = '#'
	}

	for _, line := range mat {
		fmt.Println(string(line))
	}
}

func ranges(points []point) (int, int, int, int) {
	xMin := points[0].x
	xMax := points[0].x
	yMin := points[0].y
	yMax := points[0].y
	for i := 1; i < len(points); i++ {
		p := points[i]
		xMin = min(p.x, xMin)
		xMax = max(p.x, xMax)
		yMin = min(p.y, yMin)
		yMax = max(p.y, yMax)
	}
	return xMin, xMax, yMin, yMax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func readInput() []point {
	points := make([]point, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		exp := regexp.MustCompile("position=< *(-?\\d+), *(-?\\d+)> velocity=< *(-?\\d+), *(-?\\d+)>")
		matches := exp.FindStringSubmatch(line)
		x := unsafeAtoi(matches[1])
		y := unsafeAtoi(matches[2])
		vx := unsafeAtoi(matches[3])
		vy := unsafeAtoi(matches[4])
		p := point{x, y, vx, vy}
		points = append(points, p)
	}
	return points
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
