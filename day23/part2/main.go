package main

import (
	"advent-of-code-2018/util/conv"
	"advent-of-code-2018/util/mathi"
	"advent-of-code-2018/util/slices"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type position struct {
	x int
	y int
	z int
}

func (pos position) String() string {
	return fmt.Sprintf("(%d, %d, %d)", pos.x, pos.y, pos.z)
}

func (pos position) distanceFromOrigo() int {
	return distance(position{0, 0, 0}, pos)
}

type cube struct {
	xMin int
	xMax int
	yMin int
	yMax int
	zMin int
	zMax int
}

func (c cube) split() []cube {
	dx := (c.xMax - c.xMin) / 2
	dy := (c.yMax - c.yMin) / 2
	dz := (c.zMax - c.zMin) / 2
	cubes := make([]cube, 0)

	cubes = append(cubes, cube{c.xMin, c.xMin + dx, c.yMin, c.yMin + dy, c.zMin, c.zMin + dz})
	cubes = append(cubes, cube{c.xMin + dx + 1, c.xMax, c.yMin, c.yMin + dy, c.zMin, c.zMin + dz})
	cubes = append(cubes, cube{c.xMin, c.xMin + dx, c.yMin + dy + 1, c.yMax, c.zMin, c.zMin + dz})
	cubes = append(cubes, cube{c.xMin + dx + 1, c.xMax, c.yMin + dy + 1, c.yMax, c.zMin, c.zMin + dz})

	cubes = append(cubes, cube{c.xMin, c.xMin + dx, c.yMin, c.yMin + dy, c.zMin + dz + 1, c.zMax})
	cubes = append(cubes, cube{c.xMin + dx + 1, c.xMax, c.yMin, c.yMin + dy, c.zMin + dz + 1, c.zMax})
	cubes = append(cubes, cube{c.xMin, c.xMin + dx, c.yMin + dy + 1, c.yMax, c.zMin + dz + 1, c.zMax})
	cubes = append(cubes, cube{c.xMin + dx + 1, c.xMax, c.yMin + dy + 1, c.yMax, c.zMin + dz + 1, c.zMax})

	return cubes
}

func (c cube) maxWidth() int {
	return mathi.Max(c.xMax-c.xMin, mathi.Max(c.yMax-c.yMin, c.zMax-c.zMin))
}

type nanobot struct {
	pos position
	r   int
}

func (bot nanobot) String() string {
	return fmt.Sprintf("(%d, %d, %d, r=%d)", bot.pos.x, bot.pos.y, bot.pos.z, bot.r)
}

func (bot nanobot) posInRange(pos position) bool {
	distance := distance(bot.pos, pos)
	return distance <= bot.r
}

func (bot nanobot) cubeInRange(c cube) bool {
	return bot.distanceToCube(c) <= bot.r
}

func (bot nanobot) distanceToCube(c cube) int {
	var xDist, yDist, zDist int
	if bot.pos.x > c.xMin && bot.pos.x < c.xMax {
		xDist = 0
	} else {
		xDist = mathi.Min(mathi.Abs(bot.pos.x-c.xMin), mathi.Abs(bot.pos.x-c.xMax))
	}
	if bot.pos.y > c.yMin && bot.pos.y < c.yMax {
		yDist = 0
	} else {
		yDist = mathi.Min(mathi.Abs(bot.pos.y-c.yMin), mathi.Abs(bot.pos.y-c.yMax))
	}
	if bot.pos.z > c.zMin && bot.pos.z < c.zMax {
		zDist = 0
	} else {
		zDist = mathi.Min(mathi.Abs(bot.pos.z-c.zMin), mathi.Abs(bot.pos.z-c.zMax))
	}
	return xDist + yDist + zDist
}

func main() {
	start := time.Now()

	bots := readInput()
	fmt.Println("all", len(bots))
	bots = filterByOverlapCount(bots, 900)
	fmt.Println("in cluster", len(bots))

	cubes := cubeSearch(bots)
	fmt.Println("found cubes", cubes)
	for _, c := range cubes {
		for x := c.xMin; x <= c.xMax; x++ {
			for y := c.yMin; y <= c.yMax; y++ {
				for z := c.zMin; z <= c.zMax; z++ {
					pos := position{x, y, z}
					count := posCountInRange(pos, bots)
					fmt.Println("pos", pos, "count", count)
					if count == len(bots) {
						dist := pos.distanceFromOrigo()
						fmt.Println("!!!!!!! all in range. distance from origo", dist)
					}
				}
			}
		}
	}

	fmt.Println(time.Since(start))
}

func filterByOverlapCount(bots []nanobot, count int) []nanobot {
	filteredBots := make([]nanobot, 0)
	for _, bot := range bots {
		overlapsCount := countOverlaps(bot, bots)
		if overlapsCount >= count {
			filteredBots = append(filteredBots, bot)
		}
	}
	return filteredBots
}

func cubeSearch(bots []nanobot) []cube {
	xMin, xMax := bots[0].pos.x, bots[0].pos.x
	yMin, yMax := bots[0].pos.y, bots[0].pos.y
	zMin, zMax := bots[0].pos.z, bots[0].pos.z
	for i := 1; i < len(bots); i++ {
		bot := bots[i]
		xMin = mathi.Min(bot.pos.x, xMin)
		xMax = mathi.Max(bot.pos.x, xMax)
		yMin = mathi.Min(bot.pos.y, yMin)
		yMax = mathi.Max(bot.pos.y, yMax)
		zMin = mathi.Min(bot.pos.z, zMin)
		zMax = mathi.Max(bot.pos.z, zMax)
	}
	c := cube{xMin, xMax, yMin, yMax, zMin, zMax}
	return cubeSearchAux(c, bots)
}

func cubeSearchAux(c cube, bots []nanobot) []cube {
	if c.maxWidth() <= 1 {
		return []cube{c}
	}
	cubes := c.split()
	results := make([]cube, 0)
	for _, c2 := range cubes {
		count := cubeCountInRange(c2, bots)
		if count >= len(bots) {
			auxResults := cubeSearchAux(c2, bots)
			results = append(results, auxResults...)
		}
	}
	return results
}

func distance(pos1, pos2 position) int {
	return mathi.Abs(pos1.x-pos2.x) +
		mathi.Abs(pos1.y-pos2.y) +
		mathi.Abs(pos1.z-pos2.z)
}

func doesOverlap(bot1, bot2 nanobot) bool {
	dist := distance(bot1.pos, bot2.pos)
	return dist <= bot1.r+bot2.r
}

func countOverlaps(bot nanobot, bots []nanobot) int {
	count := 0
	for _, bot2 := range bots {
		if doesOverlap(bot, bot2) {
			count++
		}
	}
	return count
}

func posCountInRange(pos position, bots []nanobot) int {
	count := 0
	for _, bot := range bots {
		if bot.posInRange(pos) {
			count++
		}
	}
	return count
}

func cubeCountInRange(c cube, bots []nanobot) int {
	count := 0
	for _, bot := range bots {
		if bot.cubeInRange(c) {
			count++
		}
	}
	return count
}

func readInput() []nanobot {
	bots := make([]nanobot, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		reg := regexp.MustCompile("^pos=<(.+)>, r=(\\d+)$")
		matches := reg.FindStringSubmatch(line)
		pos := slices.Atoi(strings.Split(matches[1], ","))
		x := pos[0]
		y := pos[1]
		z := pos[2]
		r := conv.Atoi(matches[2])
		bot := nanobot{position{x, y, z}, r}
		bots = append(bots, bot)
	}
	return bots
}
