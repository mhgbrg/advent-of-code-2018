package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type guardAction int

const (
	beginShift guardAction = iota
	fallAsleep
	wakeUp
)

type event struct {
	timestamp string
	minute    int
	action    guardAction
	guardID   int
}

type guardData struct {
	totalCount   int
	minuteCounts map[int]int
}

func (gi *guardData) update(fallAsleepMinute, wakeUpMinute int) {
	timeAsleep := wakeUpMinute - fallAsleepMinute
	gi.totalCount += timeAsleep
	for m := fallAsleepMinute; m < wakeUpMinute; m++ {
		gi.minuteCounts[m]++
	}
}

func main() {
	events := readInput()

	guardDataMap := buildGuardDataMap(events)
	guardID, minute, totalCount, minuteCount := getMostAsleep(guardDataMap)

	fmt.Println("guard", guardID, "count", totalCount)
	fmt.Println("minute", minute, "count", minuteCount)
	fmt.Println("result", guardID*minute)
}

func buildGuardDataMap(events []event) map[int]*guardData {
	guardDataMap := make(map[int]*guardData)

	var guardID int
	var fallAsleepMinute int
	for _, e := range events {
		if e.action == beginShift {
			guardID = e.guardID
		} else if e.action == fallAsleep {
			fallAsleepMinute = e.minute
		} else if e.action == wakeUp {
			wakeUpMinute := e.minute
			if _, ok := guardDataMap[guardID]; !ok {
				guardDataMap[guardID] = &guardData{0, make(map[int]int)}
			}
			guardDataMap[guardID].update(fallAsleepMinute, wakeUpMinute)
		}
	}

	return guardDataMap
}

func getMostAsleep(guardDataMap map[int]*guardData) (int, int, int, int) {
	var guardID int
	maxTotalCount := 0
	for id, data := range guardDataMap {
		if data.totalCount > maxTotalCount {
			maxTotalCount = data.totalCount
			guardID = id
		}
	}
	minute, minuteCount := getMax(guardDataMap[guardID].minuteCounts)
	return guardID, minute, maxTotalCount, minuteCount
}

func getMax(m map[int]int) (int, int) {
	maxValue := 0
	var maxKey int
	for k, v := range m {
		if v > maxValue {
			maxValue = v
			maxKey = k
		}
	}
	return maxKey, maxValue
}

func readInput() []event {
	var input []event
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		input = append(input, parsedLine)
	}
	return input
}

func parseLine(line string) event {
	timestamp := line[:18]
	minute := unsafeAtoi(timestamp[len(timestamp)-3 : len(timestamp)-1])
	action, guardID := parseAction(line[19:])
	return event{
		timestamp,
		minute,
		action,
		guardID,
	}
}

func parseAction(action string) (guardAction, int) {
	beginShiftRegexp := regexp.MustCompile("^Guard #(\\d+) begins shift$")
	fallAsleepRegexp := regexp.MustCompile("^falls asleep$")
	wakeUpRegexp := regexp.MustCompile("^wakes up$")

	if beginShiftRegexp.MatchString(action) {
		guardID := unsafeAtoi(beginShiftRegexp.FindStringSubmatch(action)[1])
		return beginShift, guardID
	} else if fallAsleepRegexp.MatchString(action) {
		return fallAsleep, 0
	} else if wakeUpRegexp.MatchString(action) {
		return wakeUp, 0
	}

	panic(fmt.Sprintf("no match for actionStr '%s'", action))
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
