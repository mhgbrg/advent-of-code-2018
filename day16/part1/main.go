package main

import (
	"advent-of-code-2018/util/slices"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type sample struct {
	before []int
	instr  instruction
	after  []int
}

type instruction struct {
	opCode int
	input1 int
	input2 int
	output int
}

type operation = func(instruction, []int) int

func main() {
	start := time.Now()

	samples := readInput()
	fmt.Println(len(samples))

	count := 0
	for _, s := range samples {
		if behavesLikeCount(s) >= 3 {
			count++
		}
	}
	fmt.Println("count", count)

	fmt.Println(time.Since(start))
}

func behavesLikeCount(s sample) int {
	ops := map[string]operation{
		"addr": addr,
		"addi": addi,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"borr": borr,
		"bori": bori,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
	}
	count := 0
	fmt.Print(s, " - ")
	for opName, opFn := range ops {
		if behavesLike(s, opFn) {
			fmt.Print(opName, " ")
			count++
		}
	}
	fmt.Printf("(%d)\n", count)
	return count
}

func behavesLike(s sample, op func(instruction, []int) int) bool {
	after := apply(op, s.instr, s.before)
	return slices.Eq(after, s.after)
}

func apply(op func(instruction, []int) int, instr instruction, state []int) []int {
	newState := make([]int, len(state))
	copy(newState, state)
	newState[instr.output] = op(instr, state)
	return newState
}

func addr(instr instruction, state []int) int {
	return state[instr.input1] + state[instr.input2]
}

func addi(instr instruction, state []int) int {
	return state[instr.input1] + instr.input2
}

func mulr(instr instruction, state []int) int {
	return state[instr.input1] * state[instr.input2]
}

func muli(instr instruction, state []int) int {
	return state[instr.input1] * instr.input2
}

func banr(instr instruction, state []int) int {
	return state[instr.input1] & state[instr.input2]
}

func bani(instr instruction, state []int) int {
	return state[instr.input1] & instr.input2
}

func borr(instr instruction, state []int) int {
	return state[instr.input1] | state[instr.input2]
}

func bori(instr instruction, state []int) int {
	return state[instr.input1] | instr.input2
}

func setr(instr instruction, state []int) int {
	return state[instr.input1]
}

func seti(instr instruction, state []int) int {
	return instr.input1
}

func gtir(instr instruction, state []int) int {
	if instr.input1 > state[instr.input2] {
		return 1
	}
	return 0
}

func gtri(instr instruction, state []int) int {
	if state[instr.input1] > instr.input2 {
		return 1
	}
	return 0
}

func gtrr(instr instruction, state []int) int {
	if state[instr.input1] > state[instr.input2] {
		return 1
	}
	return 0
}

func eqir(instr instruction, state []int) int {
	if instr.input1 == state[instr.input2] {
		return 1
	}
	return 0
}

func eqri(instr instruction, state []int) int {
	if state[instr.input1] == instr.input2 {
		return 1
	}
	return 0
}

func eqrr(instr instruction, state []int) int {
	if state[instr.input1] == state[instr.input2] {
		return 1
	}
	return 0
}

func readInput() []sample {
	samples := make([]sample, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		beforeStr := scanner.Text()
		before := slices.Atoi(strings.Split(beforeStr[9:len(beforeStr)-1], ", "))
		scanner.Scan()
		instr := parseInstruction(scanner.Text())
		scanner.Scan()
		afterStr := scanner.Text()
		after := slices.Atoi(strings.Split(afterStr[9:len(afterStr)-1], ", "))
		samples = append(samples, sample{before, instr, after})
		scanner.Scan()
	}
	return samples
}

func parseInstruction(str string) instruction {
	parts := slices.Atoi(strings.Split(str, " "))
	return instruction{parts[0], parts[1], parts[2], parts[3]}
}
