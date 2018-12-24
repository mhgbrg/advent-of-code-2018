package main

import (
	"advent-of-code-2018/util/conv"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type instruction struct {
	opName string
	input1 int
	input2 int
	output int
}

type operation = func(instruction, []int) int

func main() {
	start := time.Now()

	ipReg, program := readInput()

	result := runProgram(program, ipReg)
	fmt.Println("result", result)

	fmt.Println(time.Since(start))
}

func runProgram(program []instruction, ipReg int) int {
	state := []int{0, 0, 0, 0, 0, 0}
	for ip := 0; ip < len(program); ip++ {
		instr := program[ip]
		state[ipReg] = ip
		// fmt.Printf("ip=%d %v %v ", ip, state, instr)
		state = apply(instr, state)
		// fmt.Println(state)
		ip = state[ipReg]
	}
	return state[0]
}

func apply(instr instruction, state []int) []int {
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
	op := ops[instr.opName]
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

func readInput() (int, []instruction) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ipReg := conv.Atoi(strings.Replace(scanner.Text(), "#ip ", "", 1))
	instructions := make([]instruction, 0)
	for scanner.Scan() {
		instr := parseInstruction(scanner.Text())
		instructions = append(instructions, instr)
	}
	return ipReg, instructions
}

func parseInstruction(str string) instruction {
	parts := strings.Split(str, " ")
	return instruction{
		parts[0],
		conv.Atoi(parts[1]),
		conv.Atoi(parts[2]),
		conv.Atoi(parts[3]),
	}
}
