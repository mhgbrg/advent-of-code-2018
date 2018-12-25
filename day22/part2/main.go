package main

import (
	"advent-of-code-2018/util/grid"
	"container/heap"
	"fmt"
	"time"
)

const (
	rocky int = iota
	wet
	narrow
)

const (
	none int = iota
	torch
	climbingGear
)

func main() {
	start := time.Now()

	// depth := 510
	// target := grid.Position{X: 10, Y: 10}
	depth := 8112
	target := grid.Position{X: 13, Y: 743}

	// TODO: Replace regionTypeMap with memoized function that calculates
	// regionType lazily
	regionTypeMap := buildRegionTypeMap(depth, target)
	fmt.Println("region type map done")
	fmt.Println(time.Since(start))
	distance := shortestPath(grid.Origo(), target, depth, regionTypeMap)
	fmt.Println("result", distance)

	fmt.Println(time.Since(start))
}

func buildRegionTypeMap(depth int, target grid.Position) map[grid.Position]int {
	xLimit := 1000 // set arbitrarily to include the x values of all visited positions
	erosionLevelCache := make(map[grid.Position]int)
	regionTypeMap := make(map[grid.Position]int)
	for y := 0; y <= depth; y++ {
		for x := 0; x <= xLimit; x++ {
			pos := grid.Position{X: x, Y: y}
			geoIndex := calcGeoIndex(pos, target, erosionLevelCache)
			erosionLevel := calcErosionLevel(geoIndex, depth)
			erosionLevelCache[pos] = erosionLevel
			regionType := getRegionType(erosionLevel)
			regionTypeMap[pos] = regionType
		}
	}
	return regionTypeMap
}

func calcGeoIndex(pos grid.Position, target grid.Position, erosionLevelCache map[grid.Position]int) int {
	if pos == grid.Origo() {
		return 0
	}
	if pos == target {
		return 0
	}
	if pos.Y == 0 {
		return pos.X * 16807
	}
	if pos.X == 0 {
		return pos.Y * 48271
	}
	left, lok := erosionLevelCache[pos.Left()]
	above, aok := erosionLevelCache[pos.Above()]
	if !lok || !aok {
		panic(fmt.Sprintf("no prev erosion level for %v", pos))
	}
	return left * above
}

func calcErosionLevel(geoIndex int, depth int) int {
	return (geoIndex + depth) % 20183
}

func getRegionType(erosionLevel int) int {
	switch erosionLevel % 3 {
	case 0:
		return rocky
	case 1:
		return wet
	case 2:
		return narrow
	}
	panic(fmt.Sprintf("%d %% 3 == %d", erosionLevel, erosionLevel%3))
}

type state struct {
	pos  grid.Position
	tool int
}

type pqItem struct {
	state    state
	distance int
}

type priorityQueue []pqItem

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(pqItem)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func shortestPath(
	from grid.Position,
	to grid.Position,
	depth int,
	regionTypeMap map[grid.Position]int,
) int {
	switchCost := 7
	pq := make(priorityQueue, 0)
	heap.Push(&pq, pqItem{state{from, torch}, 0})
	visited := make(map[state]bool)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(pqItem)
		if visited[item.state] {
			continue
		}
		visited[item.state] = true
		if item.state.pos == to {
			if item.state.tool == torch {
				return item.distance
			}
			heap.Push(&pq, pqItem{state{item.state.pos, torch}, item.distance + switchCost})
		} else {
			nextStates := make([]state, 0)
			for _, pos := range item.state.pos.Around1() {
				if pos.X < 0 || pos.Y < 0 || pos.Y > depth {
					continue
				}
				nextStates = append(nextStates, state{pos, none})
				nextStates = append(nextStates, state{pos, torch})
				nextStates = append(nextStates, state{pos, climbingGear})
			}
			for _, nextState := range nextStates {
				if isValidState(state{item.state.pos, nextState.tool}, regionTypeMap) &&
					isValidState(state{nextState.pos, nextState.tool}, regionTypeMap) &&
					!visited[nextState] {
					distance := item.distance + 1
					if item.state.tool != nextState.tool {
						distance += switchCost
					}
					heap.Push(&pq, pqItem{nextState, distance})
				}
			}
		}
	}
	panic("no route to target")
}

func isValidState(s state, regionTypeMap map[grid.Position]int) bool {
	regionType, ok := regionTypeMap[s.pos]
	if !ok {
		panic(fmt.Sprintf("no region type at %v", s.pos))
	}
	switch regionType {
	case rocky:
		return s.tool == torch || s.tool == climbingGear
	case wet:
		return s.tool == climbingGear || s.tool == none
	case narrow:
		return s.tool == torch || s.tool == none
	}
	panic(fmt.Sprintf("invalid region type %d at %v", regionType, s.pos))
}
