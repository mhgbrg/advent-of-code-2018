package main

import (
	"advent-of-code-2018/util/conv"
	"advent-of-code-2018/util/grid"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

const (
	down int = iota
	left
	right
)

type state struct {
	well         grid.Position
	clay         map[grid.Position]bool
	runningWater map[grid.Position]bool
	staleWater   map[grid.Position]bool
	limits       grid.Limits
}

func (s *state) occupied(pos grid.Position) bool {
	return s.clay[pos] || s.staleWater[pos]
}

func main() {
	start := time.Now()

	clay := readInput()
	well := grid.Position{X: 500, Y: 0}
	s := &state{
		well,
		clay,
		make(map[grid.Position]bool),
		make(map[grid.Position]bool),
		grid.CalcLimits(clay),
	}
	// term.ClearScreen()
	goDown(s, well)
	// printMap(s)
	waterCount := countWater(s)
	fmt.Println("water", waterCount)

	fmt.Println(time.Since(start))
}

func goDown(s *state, pos grid.Position) {
	if pos.Y < 0 || pos.Y > s.limits.YMax || s.runningWater[pos] {
		return
	}
	s.runningWater[pos] = true
	if !s.occupied(pos.Below()) {
		goDown(s, pos.Below())
	} else {
		split(s, pos)
	}
}

func split(s *state, pos grid.Position) {
	lpos, lstale := goLeft(s, pos)
	rpos, rstale := goRight(s, pos)
	if lstale && rstale {
		for p := pos.Left(); p != lpos.Left(); p = p.Left() {
			s.staleWater[p] = true
		}
		for p := pos.Right(); p != rpos.Right(); p = p.Right() {
			s.staleWater[p] = true
		}
		s.staleWater[pos] = true
		split(s, pos.Above())
	} else if rstale {
		goDown(s, lpos.Below())
	} else if lstale {
		goDown(s, rpos.Below())
	} else {
		goDown(s, lpos.Below())
		goDown(s, rpos.Below())
	}
}

func goLeft(s *state, pos grid.Position) (grid.Position, bool) {
	s.runningWater[pos] = true
	if !s.occupied(pos.Below()) {
		return pos, false
	}
	if s.occupied(pos.Left()) {
		return pos, true
	}
	return goLeft(s, pos.Left())
}

func goRight(s *state, pos grid.Position) (grid.Position, bool) {
	s.runningWater[pos] = true
	if !s.occupied(pos.Below()) {
		return pos, false
	}
	if s.occupied(pos.Right()) {
		return pos, true
	}
	return goRight(s, pos.Right())
}

func countWater(s *state) int {
	water := make(map[grid.Position]bool)
	for pos := range s.staleWater {
		if pos.Y >= s.limits.YMin {
			water[pos] = true
		}
	}
	for pos := range s.runningWater {
		if pos.Y >= s.limits.YMin {
			water[pos] = true
		}
	}
	return len(water)
}

func printMap(s *state) {
	for y := 0; y <= s.limits.YMax; y++ {
		for x := s.limits.XMin - 1; x <= s.limits.XMax+1; x++ {
			pos := grid.Position{X: x, Y: y}
			if pos == s.well {
				fmt.Print("+")
			} else if s.clay[pos] {
				fmt.Print("#")
			} else if s.staleWater[pos] {
				fmt.Print("~")
			} else if s.runningWater[pos] {
				fmt.Print("|")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func readInput() map[grid.Position]bool {
	clay := make(map[grid.Position]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		xFirst := regexp.MustCompile("^x=(\\d+), y=(\\d+)..(\\d+)")
		if xFirst.MatchString(line) {
			parts := xFirst.FindStringSubmatch(line)
			x := conv.Atoi(parts[1])
			yStart := conv.Atoi(parts[2])
			yEnd := conv.Atoi(parts[3])
			for y := yStart; y <= yEnd; y++ {
				clay[grid.Position{X: x, Y: y}] = true
			}
		} else {
			yFirst := regexp.MustCompile("^y=(\\d+), x=(\\d+)..(\\d+)")
			parts := yFirst.FindStringSubmatch(line)
			y := conv.Atoi(parts[1])
			xStart := conv.Atoi(parts[2])
			xEnd := conv.Atoi(parts[3])
			for x := xStart; x <= xEnd; x++ {
				clay[grid.Position{X: x, Y: y}] = true
			}
		}
	}
	return clay
}
