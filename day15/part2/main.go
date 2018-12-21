package main

import (
	"advent-of-code-2018/util/mathi"
	"advent-of-code-2018/util/term"
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type pos struct {
	x int
	y int
}

func (p pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (p pos) lessThan(p2 pos) bool {
	if p.y < p2.y {
		return true
	}
	if p.y == p2.y {
		return p.x < p2.x
	}
	return false
}

const (
	elf    rune = 'E'
	goblin rune = 'G'
)

type unit struct {
	ap   int
	hp   int
	p    pos
	race rune
}

func (u *unit) isElf() bool {
	return u.race == elf
}

func (u *unit) isGoblin() bool {
	return u.race == goblin
}

func (u *unit) isDead() bool {
	return u.hp <= 0
}

func (u *unit) inRange(u2 *unit) bool {
	return u.p.x == u2.p.x && mathi.Abs(u.p.y-u2.p.y) == 1 ||
		mathi.Abs(u.p.x-u2.p.x) == 1 && u.p.y == u2.p.y
}

func (u *unit) copy() *unit {
	return &unit{u.ap, u.hp, u.p, u.race}
}

func (u unit) String() string {
	return fmt.Sprintf("%q %d %s", u.race, u.hp, u.p)
}

func main() {
	start := time.Now()

	m, originalUnits := readInput()

	for ap := 4; ; ap++ {
		term.ClearScreen()
		fmt.Println("ap", ap)
		units := deepCopy(originalUnits)
		outcome, ok := runBattle(m, units, ap)
		if ok {
			fmt.Println("outcome", outcome)
			break
		}
	}

	fmt.Println(time.Since(start))
}

func runBattle(m [][]rune, units []*unit, ap int) (int, bool) {
	for _, u := range units {
		if u.isElf() {
			u.ap = ap
		}
	}

	for round := 1; ; round++ {
		term.Goto(1, 0)
		term.ClearLine()
		fmt.Println("round", round)
		sortUnits(units)
		for _, u := range units {
			if u.isDead() {
				continue
			}
			targets := getTargets(u, units)
			if len(targets) == 0 {
				printMap(m, units)
				term.ClearLine()
				fmt.Println("total hp", totalHP(units))
				return (round - 1) * totalHP(units), true
			}
			if !anyInRange(u, targets) {
				move(u, targets, m, units)
			}
			attack(u, targets, m)
			if anyDead(units, elf) {
				printMap(m, units)
				term.ClearLine()
				fmt.Println("total hp", totalHP(units))
				return 0, false
			}
		}
		printMap(m, units)
		term.ClearLine()
		fmt.Println("total hp", totalHP(units))
		// time.Sleep(200 * time.Millisecond)
	}
}

func getTargets(u *unit, units []*unit) []*unit {
	targets := make([]*unit, 0)
	for _, t := range units {
		if t.isDead() {
			continue
		}
		if u.isElf() && t.isGoblin() || u.isGoblin() && t.isElf() {
			targets = append(targets, t)
		}
	}
	return targets
}

func anyDead(units []*unit, race rune) bool {
	for _, u := range units {
		if u.race == race && u.isDead() {
			return true
		}
	}
	return false
}

func anyInRange(u *unit, units []*unit) bool {
	for _, u2 := range units {
		if u.inRange(u2) {
			return true
		}
	}
	return false
}

func move(u *unit, targets []*unit, m [][]rune, units []*unit) {
	path := pathToNearestTarget(u, targets, m, units)
	if path != nil {
		u.p = path[0]
	}
}

func pathToNearestTarget(u *unit, targets []*unit, m [][]rune, units []*unit) []pos {
	occupied := buildOccupiedMap(units)
	targetPositions := make(map[pos]bool)
	for _, t := range targets {
		around := getSurrounding(t.p)
		for _, p := range around {
			if m[p.y][p.x] == '.' && !occupied[p] {
				targetPositions[p] = true
			}
		}
	}
	if len(targetPositions) == 0 {
		return nil
	}
	paths := bfs(u.p, targetPositions, m, occupied)
	if len(paths) == 0 {
		return nil
	}
	sortPaths(paths)
	return paths[0]
}

func buildOccupiedMap(units []*unit) map[pos]bool {
	occupied := make(map[pos]bool)
	for _, u := range units {
		if !u.isDead() {
			occupied[u.p] = true
		}
	}
	return occupied
}

func getSurrounding(p pos) []pos {
	return []pos{
		pos{p.x, p.y - 1},
		pos{p.x - 1, p.y},
		pos{p.x + 1, p.y},
		pos{p.x, p.y + 1},
	}
}

type bfsNode struct {
	p    pos
	path []pos
}

type queue struct {
	items []bfsNode
}

func (q *queue) add(u bfsNode) {
	q.items = append(q.items, u)
}

func (q *queue) remove() bfsNode {
	node := q.items[0]
	q.items = q.items[1:]
	return node
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func bfs(from pos, to map[pos]bool, m [][]rune, occupied map[pos]bool) [][]pos {
	q := queue{[]bfsNode{}}
	q.add(bfsNode{from, []pos{}})
	visited := make(map[pos]bool)
	foundPaths := make([][]pos, 0)
	for !q.isEmpty() {
		node := q.remove()
		if len(foundPaths) > 0 && len(node.path) > len(foundPaths[0]) {
			break
		}
		if to[node.p] {
			foundPaths = append(foundPaths, node.path)
		}
		if visited[node.p] {
			continue
		}
		for _, p := range getSurrounding(node.p) {
			if m[p.y][p.x] == '.' && !occupied[p] {
				path := copyAppend(node.path, p)
				q.add(bfsNode{p, path})
			}
		}
		visited[node.p] = true
	}
	return foundPaths
}

func attack(u *unit, targets []*unit, m [][]rune) {
	sortUnits(targets)
	var target *unit
	for _, t := range targets {
		if u.inRange(t) && (target == nil || t.hp < target.hp) {
			target = t
		}
	}
	if target != nil {
		target.hp -= u.ap
	}
}

func totalHP(units []*unit) int {
	hp := 0
	for _, u := range units {
		if !u.isDead() {
			hp += u.hp
		}
	}
	return hp
}

func sortPositions(positions []pos) {
	sort.Slice(positions, func(i, j int) bool {
		p1 := positions[0]
		p2 := positions[1]
		return p1.lessThan(p2)
	})
}

func sortUnits(units []*unit) {
	sort.Slice(units, func(i, j int) bool {
		u1 := units[i]
		u2 := units[j]
		return u1.p.lessThan(u2.p)
	})
}

func sortPaths(paths [][]pos) {
	sort.Slice(paths, func(i, j int) bool {
		p1 := paths[i]
		p2 := paths[j]
		return p1[len(p1)-1].lessThan(p2[len(p2)-1])
	})
}

func printMap(m [][]rune, units []*unit) {
	for y := 0; y < len(m); y++ {
		row := m[y]
	ROW:
		for x := 0; x < len(row); x++ {
			for _, u := range units {
				if u.p.x == x && u.p.y == y && !u.isDead() {
					if u.hp == 200 {
						term.Color(term.Green)
					} else if u.hp > 100 {
						term.Color(term.Yellow)
					} else {
						term.Color(term.Red)
					}
					fmt.Print(string(u.race))
					term.Color(term.White)
					continue ROW
				}
			}
			fmt.Print(string(row[x]))
		}
		fmt.Println()
	}
}

func readInput() ([][]rune, []*unit) {
	m := make([][]rune, 0)
	units := make([]*unit, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		line := []rune(scanner.Text())
		for x, r := range line {
			if r == 'E' || r == 'G' {
				u := &unit{3, 200, pos{x, y}, r}
				units = append(units, u)
				line[x] = '.'
			}
		}
		m = append(m, line)
	}
	return m, units
}

func deepCopy(slice []*unit) []*unit {
	newSlice := make([]*unit, len(slice))
	for i := 0; i < len(slice); i++ {
		newSlice[i] = slice[i].copy()
	}
	return newSlice
}

func copyAppend(slice []pos, elem pos) []pos {
	newSlice := make([]pos, len(slice))
	copy(newSlice, slice)
	return append(newSlice, elem)
}
