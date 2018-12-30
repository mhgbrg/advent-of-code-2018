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

type nanobot struct {
	x int
	y int
	z int
	r int
}

func (bot nanobot) inRange(other nanobot) bool {
	distance := mathi.Abs(bot.x-other.x) +
		mathi.Abs(bot.y-other.y) +
		mathi.Abs(bot.z-other.z)
	return distance <= bot.r
}

func main() {
	start := time.Now()

	bots := readInput()
	fmt.Println(bots)

	strongestBot := bots[0]
	for i := 1; i < len(bots); i++ {
		bot := bots[i]
		if bot.r > strongestBot.r {
			strongestBot = bot
		}
	}
	fmt.Println("strongest bot", strongestBot)

	inRangeCount := 0
	for _, bot := range bots {
		if strongestBot.inRange(bot) {
			inRangeCount++
		}
	}
	fmt.Println("in range count", inRangeCount)

	fmt.Println(time.Since(start))
}

func readInput() []nanobot {
	bots := make([]nanobot, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		reg := regexp.MustCompile("^pos=<(.+)>, r=(\\d+)$")
		matches := reg.FindStringSubmatch(line)
		pos := slices.Atoi(strings.Split(matches[1], ","))
		x := pos[0]
		y := pos[1]
		z := pos[2]
		r := conv.Atoi(matches[2])
		bot := nanobot{x, y, z, r}
		bots = append(bots, bot)
	}
	return bots
}
