package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type marble struct {
	number int
	prev   *marble
	next   *marble
}

func (m marble) String() string {
	return fmt.Sprintf("%d <- %d -> %d", m.prev.number, m.number, m.next.number)
}

func (m *marble) addAfter(number int) *marble {
	nm := &marble{number, m, m.next}
	m.next.prev = nm
	m.next = nm
	return nm
}

func (m *marble) remove() {
	m.prev.next = m.next
	m.next.prev = m.prev
}

func (m *marble) nprev(number int) *marble {
	nm := m
	for i := 0; i < number; i++ {
		nm = nm.prev
	}
	return nm
}

func main() {
	start := time.Now()

	numberOfPlayers, lastMarbleNumber := readInput()
	lastMarbleNumber *= 100

	fmt.Println("people", numberOfPlayers)
	fmt.Println("lastMarble", lastMarbleNumber)

	scores := make(map[int]int)

	currentMarble := &marble{0, nil, nil}
	currentMarble.prev = currentMarble
	currentMarble.next = currentMarble

	player := 0
	for marbleNumber := 1; marbleNumber <= lastMarbleNumber; marbleNumber++ {
		if marbleNumber%23 == 0 {
			rm := currentMarble.nprev(7)
			rm.remove()
			currentMarble = rm.next
			scores[player] += marbleNumber + rm.number
		} else {
			nm := currentMarble.next.addAfter(marbleNumber)
			currentMarble = nm
		}
		player = (player + 1) % numberOfPlayers
	}

	player, score := mapMax(scores)

	fmt.Println("player", player+1, "score", score)

	fmt.Println(time.Since(start))
}

func mapMax(m map[int]int) (int, int) {
	mv := 0
	var mk int
	for k, v := range m {
		if v > mv {
			mv = v
			mk = k
		}
	}
	return mk, mv
}

func readInput() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	numberOfPlayers := unsafeAtoi(parts[0])
	lastMarble := unsafeAtoi(parts[len(parts)-2])
	return numberOfPlayers, lastMarble
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
