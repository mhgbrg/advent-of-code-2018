package main

import (
	"advent-of-code-2018/util/grid"
	"advent-of-code-2018/util/term"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	area := readInput()
	printArea(area)

	minutes := 1000000000

	term.ClearScreen()
	term.Goto(0, 0)
	fmt.Println(0)
	printArea(area)
	time.Sleep(200 * time.Millisecond)

	areas := make(map[string]int, 0)

	for i := 1; i <= minutes; i++ {
		area = mutateArea(area)
		result := calculateResult(area)

		term.Goto(0, 0)
		printArea(area)
		fmt.Println("minute", i, "result", result)

		areaStr := areaToString(area)
		if j, ok := areas[areaStr]; ok {
			loopLength := i - j
			minutesLeft := minutes - i
			iterationsLeft := minutes - ((minutesLeft/loopLength)*loopLength + i)
			minutes = i + iterationsLeft
		}
		areas[areaStr] = i

		// time.Sleep(100 * time.Millisecond)
	}

	fmt.Println(time.Since(start))
}

func mutateArea(area [][]rune) [][]rune {
	mutated := make([][]rune, len(area))
	for y := 0; y < len(area); y++ {
		mutated[y] = make([]rune, len(area[y]))
		for x := 0; x < len(area[y]); x++ {
			if x == 0 || x == len(area[y])-1 || y == 0 || y == len(area)-1 {
				mutated[y][x] = '%'
			} else {
				result := mutateAcre(grid.Position{X: x, Y: y}, area)
				mutated[y][x] = result
			}
		}
	}
	return mutated
}

func mutateAcre(pos grid.Position, area [][]rune) rune {
	around := pos.Around2()
	treeCount, lumberyardCount := 0, 0
	for _, p := range around {
		if area[p.Y][p.X] == '|' {
			treeCount++
		} else if area[p.Y][p.X] == '#' {
			lumberyardCount++
		}
	}
	switch area[pos.Y][pos.X] {
	case '.':
		if treeCount >= 3 {
			return '|'
		}
		return '.'
	case '|':
		if lumberyardCount >= 3 {
			return '#'
		}
		return '|'
	case '#':
		if lumberyardCount >= 1 && treeCount >= 1 {
			return '#'
		}
		return '.'
	}
	panic(fmt.Sprintf("Invalid acre %q", area[pos.Y][pos.X]))
}

func calculateResult(area [][]rune) int {
	treeCount, lumberyardCount := 0, 0
	for _, row := range area {
		for _, acre := range row {
			if acre == '|' {
				treeCount++
			} else if acre == '#' {
				lumberyardCount++
			}
		}
	}
	return treeCount * lumberyardCount
}

func printArea(area [][]rune) {
	for _, row := range area {
		fmt.Println(string(row))
	}
}

func areaToString(area [][]rune) string {
	var sb strings.Builder
	for _, row := range area {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

func readInput() [][]rune {
	area := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := []rune("%" + scanner.Text() + "%")
		area = append(area, row)
	}
	emptyRow := make([]rune, len(area[0]))
	for i := range emptyRow {
		emptyRow[i] = '%'
	}
	area = append([][]rune{emptyRow}, area...)
	area = append(area, emptyRow)
	return area
}
