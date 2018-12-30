package main

import (
	"advent-of-code-2018/util/conv"
	"advent-of-code-2018/util/mathi"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type point struct {
	w int
	x int
	y int
	z int
}

type constellation []point

func main() {
	start := time.Now()

	points := readInput()
	fmt.Println(points)

	constellations := make([]constellation, 0)
	for _, p := range points {
		constellations = append(constellations, []point{p})
	}

	for {
		newConstellations := make([]constellation, 0)
		merged := make(map[int]bool)
	OUTER:
		for i := 0; i < len(constellations); i++ {
			if merged[i] {
				continue
			}
			c1 := constellations[i]
			for j := i + 1; j < len(constellations); j++ {
				if merged[j] {
					continue
				}
				c2 := constellations[j]
				if constellationDistance(c1, c2) <= 3 {
					c3 := merge(c1, c2)
					newConstellations = append(newConstellations, c3)
					merged[i], merged[j] = true, true
					continue OUTER
				}
			}
		}
		for i, c := range constellations {
			if !merged[i] {
				newConstellations = append(newConstellations, c)
			}
		}
		constellations = newConstellations
		if len(merged) == 0 {
			break
		}
	}

	fmt.Println(constellations)
	fmt.Println("number of constellations", len(constellations))

	fmt.Println(time.Since(start))
}

func constellationDistance(c1, c2 constellation) int {
	minDist := pointDistance(c1[0], c2[0])
	for i := 0; i < len(c1); i++ {
		p1 := c1[i]
		for j := 0; j < len(c2); j++ {
			p2 := c2[j]
			dist := pointDistance(p1, p2)
			minDist = mathi.Min(dist, minDist)
		}
	}
	return minDist
}

func pointDistance(p1, p2 point) int {
	return mathi.Abs(p1.w-p2.w) +
		mathi.Abs(p1.x-p2.x) +
		mathi.Abs(p1.y-p2.y) +
		mathi.Abs(p1.z-p2.z)
}

func merge(c1, c2 constellation) constellation {
	c3 := make(constellation, 0)
	c3 = append(c3, c1...)
	return append(c3, c2...)
}

func readInput() []point {
	scanner := bufio.NewScanner(os.Stdin)
	points := make([]point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		w := conv.Atoi(parts[0])
		x := conv.Atoi(parts[1])
		y := conv.Atoi(parts[2])
		z := conv.Atoi(parts[3])
		points = append(points, point{w, x, y, z})
	}
	return points
}
