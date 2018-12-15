package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	jobs, depsMap := readInput()

	fmt.Println("jobs", jobs)
	fmt.Println("depsMap", depsMap)

	order := make([]rune, 0)
	completed := make(map[rune]bool)

OUTER:
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		if completed[job] {
			continue
		}
		deps := depsMap[job]
		for _, dep := range deps {
			if !completed[dep] {
				continue OUTER
			}
		}
		order = append(order, job)
		completed[job] = true
		i = -1
	}

	fmt.Println("completed", completed)
	fmt.Println(string(order))

	fmt.Println(time.Since(start))
}

func readInput() ([]rune, map[rune][]rune) {
	jobsMap := make(map[rune]bool)
	deps := make(map[rune][]rune)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		job, dep := parseLine(line)
		jobsMap[job] = true
		jobsMap[dep] = true
		deps[job] = append(deps[job], dep)
	}
	jobs := make([]rune, 0)
	for k := range jobsMap {
		jobs = append(jobs, k)
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] < jobs[j]
	})
	return jobs, deps
}

func parseLine(line string) (rune, rune) {
	exp := regexp.MustCompile("^Step ([A-Z]) .+ step ([A-Z])")
	matches := exp.FindStringSubmatch(line)
	job := rune(matches[2][0])
	dep := rune(matches[1][0])
	return job, dep
}
