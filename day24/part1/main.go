package main

import (
	"advent-of-code-2018/util/conv"
	"advent-of-code-2018/util/mathi"
	"advent-of-code-2018/util/slices"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	teamImmuneSystem int = iota
	teamInfection
)

type army map[*group]bool

type group struct {
	id         int
	team       int
	units      int
	hp         int
	attackDmg  int
	attackType string
	initiative int
	weaknesses []string
	immunities []string
}

func (g *group) effectivePower() int {
	return g.units * g.attackDmg
}

func (g *group) isWeakAgainst(attackType string) bool {
	return slices.StrContains(g.weaknesses, attackType)
}

func (g *group) isImmuneAgainst(attackType string) bool {
	return slices.StrContains(g.immunities, attackType)
}

func (g *group) powerAgainst(g2 *group) int {
	if g2.isWeakAgainst(g.attackType) {
		return g.effectivePower() * 2
	} else if g2.isImmuneAgainst(g.attackType) {
		return 0
	}
	return g.effectivePower()
}

func (g group) String() string {
	var teamStr string
	if g.team == teamImmuneSystem {
		teamStr = "immune system"
	} else {
		teamStr = "infection"
	}
	return fmt.Sprintf("%s %d", teamStr, g.id)
}

func main() {
	start := time.Now()

	immuneSystem, infection := readInput()

	for i := 1; len(immuneSystem) > 0 && len(infection) > 0; i++ {
		fmt.Println("round", i)
		targets := selectTargets(immuneSystem, infection)
		attack(targets, immuneSystem, infection)
	}

	unitsLeft := 0
	for g := range immuneSystem {
		unitsLeft += g.units
	}
	for g := range infection {
		unitsLeft += g.units
	}
	fmt.Println("units left", unitsLeft)

	fmt.Println(time.Since(start))
}

func selectTargets(immuneSystem, infection army) map[*group]*group {
	targets := make(map[*group]*group)

	immuneSystemTargets, infectionTargets := make(army), make(army)
	for k, v := range immuneSystem {
		immuneSystemTargets[k] = v
	}
	for k, v := range infection {
		infectionTargets[k] = v
	}

	groups := combineArmies(immuneSystem, infection)
	sortInSelectPriorityOrder(groups)

	for _, g := range groups {
		var otherTeam map[*group]bool

		if g.team == teamImmuneSystem {
			otherTeam = infectionTargets
		} else {
			otherTeam = immuneSystemTargets
		}
		target := selectTarget(g, otherTeam)
		if target != nil {
			targets[g] = target
			delete(otherTeam, target)
		}
	}

	return targets
}

func combineArmies(immuneSystem, infection army) []*group {
	groups := make([]*group, 0)
	for g := range immuneSystem {
		groups = append(groups, g)
	}
	for g := range infection {
		groups = append(groups, g)
	}
	return groups
}

func sortInSelectPriorityOrder(groups []*group) {
	sort.Slice(groups, func(i, j int) bool {
		g1, g2 := groups[i], groups[j]
		if g1.effectivePower() == g2.effectivePower() {
			return g1.initiative > g2.initiative
		}
		return g1.effectivePower() > g2.effectivePower()
	})
}

func sortInAttackPriorityOrder(groups []*group) {
	sort.Slice(groups, func(i, j int) bool {
		g1, g2 := groups[i], groups[j]
		return g1.initiative > g2.initiative
	})
}

func selectTarget(g *group, targets army) *group {
	var target *group
	for t := range targets {
		// fmt.Println(g, "?>", t, g.powerAgainst(t), "dmg")
		if g.powerAgainst(t) > 0 &&
			(target == nil ||
				g.powerAgainst(t) > g.powerAgainst(target) ||
				g.powerAgainst(t) == g.powerAgainst(target) && t.effectivePower() > target.effectivePower() ||
				g.powerAgainst(t) == g.powerAgainst(target) && t.effectivePower() == target.effectivePower() && t.initiative > target.initiative) {
			target = t
		}
	}
	return target
}

func attack(targets map[*group]*group, immuneSystem, infection army) {
	groups := combineArmies(immuneSystem, infection)
	sortInAttackPriorityOrder(groups)
	for _, attacker := range groups {
		if attacker.units <= 0 {
			continue
		}
		defender, ok := targets[attacker]
		if !ok || defender.units <= 0 {
			continue
		}
		power := attacker.powerAgainst(defender)
		unitsLost := power / defender.hp
		fmt.Println(attacker, ">", defender, power, "dmg", mathi.Min(unitsLost, defender.units), "units")
		defender.units -= unitsLost
		if defender.units <= 0 {
			fmt.Println(defender, "dead")
			if defender.team == teamImmuneSystem {
				delete(immuneSystem, defender)
			} else {
				delete(infection, defender)
			}
		}
	}
}

func readInput() (army, army) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Immune System:
	immuneSystem, infection := make(army), make(army)
	id := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		g := parseLine(id, teamImmuneSystem, line)
		immuneSystem[g] = true
		id++
	}
	scanner.Scan() // Infection:
	id = 1
	for scanner.Scan() {
		line := scanner.Text()
		g := parseLine(id, teamInfection, line)
		infection[g] = true
		id++
	}
	return immuneSystem, infection
}

func parseLine(id, team int, line string) *group {
	// TODO: Clean up. Match with only three regexes
	units := conv.Atoi(matchOneGroup(line, "(\\d+) units"))
	hp := conv.Atoi(matchOneGroup(line, "(\\d+) hit points"))
	weaknesses := matchAndSplit(line, "weak to ([a-z, ]+)")
	immunities := matchAndSplit(line, "immune to ([a-z, ]+)")
	attackDmgStr, attackType := matchTwoGroups(line, "(\\d+) ([a-z]+) damage")
	attackDmg := conv.Atoi(attackDmgStr)
	initiative := conv.Atoi(matchOneGroup(line, "initiative (\\d+)"))
	return &group{
		id,
		team,
		units,
		hp,
		attackDmg,
		attackType,
		initiative,
		weaknesses,
		immunities,
	}
}

func matchOneGroup(str string, regex string) string {
	return regexp.MustCompile(regex).FindStringSubmatch(str)[1]
}

func maybeMatchOneGroup(str string, regex string) (string, bool) {
	matches := regexp.MustCompile(regex).FindStringSubmatch(str)
	if len(matches) == 0 {
		return "", false
	}
	return string(matches[1]), true
}

func matchAndSplit(str string, regex string) []string {
	match, found := maybeMatchOneGroup(str, regex)
	if !found {
		return []string{}
	}
	return strings.Split(match, ", ")
}

func matchTwoGroups(str string, regex string) (string, string) {
	matches := regexp.MustCompile(regex).FindStringSubmatch(str)
	return matches[1], matches[2]
}
