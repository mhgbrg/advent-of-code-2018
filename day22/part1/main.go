package main

import (
	"advent-of-code-2018/util/grid"
	"fmt"
	"time"
)

const (
	rocky int = iota
	wet
	narrow
)

func main() {
	start := time.Now()

	// depth := 510
	// target := grid.Position{X: 10, Y: 10}
	depth := 8112
	target := grid.Position{X: 13, Y: 743}
	erosionLevelCache := make(map[grid.Position]int)
	totalRiskLevel := 0
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			pos := grid.Position{X: x, Y: y}
			geoIndex := calcGeoIndex(pos, target, erosionLevelCache)
			erosionLevel := calcErosionLevel(geoIndex, depth)
			erosionLevelCache[pos] = erosionLevel
			regionType := getRegionType(erosionLevel)
			riskLevel := getRiskLevel(regionType)
			totalRiskLevel += riskLevel
			fmt.Println(pos, "geo", geoIndex, "ero", erosionLevel, "type", regionType, "risk", riskLevel)
		}
	}
	fmt.Println("total risk level", totalRiskLevel)

	fmt.Println(time.Since(start))
}

func calcGeoIndex(pos grid.Position, target grid.Position, cache map[grid.Position]int) int {
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
	return cache[pos.Left()] * cache[pos.Above()]
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

func getRiskLevel(regionType int) int {
	return regionType
}
