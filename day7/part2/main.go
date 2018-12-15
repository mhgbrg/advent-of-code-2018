package main

import (
	"bufio"
	"fmt"
	"math"
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

	baseTime := 60
	numberOfWorkers := 5

	completed := make(map[rune]int)
	workers := make([]int, numberOfWorkers)
	totalTime := 0

	for len(completed) < len(jobs) {
		for i := 0; i < len(jobs); i++ {
			job := jobs[i]
			if completed[job] > 0 {
				continue
			}
			maxDepTime := getMaxDepTime(depsMap[job], completed)
			if maxDepTime == -1 {
				continue
			}
			worker, workerTime := arrayMin(workers)
			startTime := max(maxDepTime, workerTime)
			completionTime := startTime + baseTime + int(job-'A'+1)
			completed[job] = completionTime
			workers[worker] = completionTime
			totalTime = max(totalTime, completionTime)
			fmt.Printf("%q %d-%d (worker %d)\n", job, startTime, completionTime, worker+1)
		}
	}

	fmt.Println("completed", completed)
	fmt.Println("totalTime", totalTime)

	fmt.Println(time.Since(start))
}

func getMaxDepTime(deps []rune, completed map[rune]int) int {
	maxTime := 0
	for _, dep := range deps {
		time := completed[dep]
		if time == 0 {
			return -1
		}
		if time > maxTime {
			maxTime = time
		}
	}
	return maxTime
}

func arrayMin(arr []int) (int, int) {
	minValue := math.MaxInt32
	var minIndex int
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}
	return minIndex, minValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
