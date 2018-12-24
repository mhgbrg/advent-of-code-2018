package main

import (
	"advent-of-code-2018/util/grid"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	regex := readInput()
	graph := constructGraph(regex)
	distance := findFurthest(graph)
	fmt.Println("result", distance)

	fmt.Println(time.Since(start))
}

func constructGraph(regex string) map[grid.Position][]grid.Position {
	graph := make(map[grid.Position][]grid.Position, 0)
	pos := grid.Position{X: 0, Y: 0}
	constructGraphAux(regex, pos, graph)
	return graph
}

func constructGraphAux(regex string, pos grid.Position, graph map[grid.Position][]grid.Position) {
	for i := 0; i < len(regex); i++ {
		nextPosMap := map[rune]grid.Position{
			'N': pos.Above(),
			'S': pos.Below(),
			'W': pos.Left(),
			'E': pos.Right(),
		}
		d := regex[i]
		switch d {
		case 'N', 'S', 'W', 'E':
			nextPos := nextPosMap[rune(d)]
			graph[pos] = append(graph[pos], nextPos)
			pos = nextPos
		case '(':
			splits := parseSplit(regex[i:])
			for _, s := range splits {
				constructGraphAux(s, pos, graph)
				i += len(s) + 1
			}
		default:
			panic(fmt.Sprintf("Should not have gotten rune %q, regex=%s, i=%d", d, regex, i))
		}
	}
}

func parseSplit(regex string) []string {
	level := 0
	splits := make([]string, 0)
	start := 1
	for i, foundEnd := 0, false; i < len(regex) && !foundEnd; i++ {
		switch regex[i] {
		case '(':
			level++
		case ')':
			level--
			if level == 0 {
				splits = append(splits, regex[start:i])
				foundEnd = true
			}
		case '|':
			if level == 1 {
				splits = append(splits, regex[start:i])
				start = i + 1
			}
		}
	}
	return splits
}

type bfsNode struct {
	pos      grid.Position
	distance int
}

type queue struct {
	items []bfsNode
}

func (q *queue) add(pos grid.Position, distance int) {
	q.items = append(q.items, bfsNode{pos, distance})
}

func (q *queue) remove() (grid.Position, int) {
	node := q.items[0]
	q.items = q.items[1:]
	return node.pos, node.distance
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func findFurthest(graph map[grid.Position][]grid.Position) int {
	q := queue{}
	q.add(grid.Origo(), 0)
	visited := make(map[grid.Position]bool)
	furthestDistance := 0
	for !q.isEmpty() {
		pos, distance := q.remove()
		if visited[pos] {
			continue
		}
		visited[pos] = true
		if distance > furthestDistance {
			furthestDistance = distance
		}
		neighbours := graph[pos]
		for _, adjacentPos := range neighbours {
			q.add(adjacentPos, distance+1)
		}
	}
	return furthestDistance
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	regex := scanner.Text()
	return regex[1 : len(regex)-1]
}
