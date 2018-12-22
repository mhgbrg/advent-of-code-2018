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

type operation struct {
	name string
	fn   func(instruction, []int) int
}

func main() {
	start := time.Now()

	samples := readSamples()
	program := readProgram()

	assignments := assignOpCodes(samples)
	result := runProgram(program, assignments)

	fmt.Println("result", result)

	fmt.Println(time.Since(start))
}

func assignOpCodes(samples []sample) map[int]operation {
	ops := map[string]operation{
		"addr": operation{"addr", addr},
		"addi": operation{"addi", addi},
		"mulr": operation{"mulr", mulr},
		"muli": operation{"muli", muli},
		"banr": operation{"banr", banr},
		"bani": operation{"bani", bani},
		"borr": operation{"borr", borr},
		"bori": operation{"bori", bori},
		"setr": operation{"setr", setr},
		"seti": operation{"seti", seti},
		"gtir": operation{"gtir", gtir},
		"gtri": operation{"gtri", gtri},
		"gtrr": operation{"gtrr", gtrr},
		"eqir": operation{"eqir", eqir},
		"eqri": operation{"eqri", eqri},
		"eqrr": operation{"eqrr", eqrr},
	}
	numOps := len(ops)
	assignments := make(map[int]operation)

	for len(assignments) < numOps {
		for _, s := range samples {
			possible := getPossibleOps(s, ops)
			if len(possible) == 1 {
				opCode := s.instr.opCode
				op := possible[0]
				assignments[opCode] = op
				delete(ops, op.name)
				fmt.Println(opCode, "==", op.name)
			}
		}
	}

	return assignments
}

func getPossibleOps(s sample, ops map[string]operation) []operation {
	possible := make([]operation, 0)
	for _, op := range ops {
		if behavesLike(s, op) {
			possible = append(possible, op)
		}
	}
	return possible
}

func behavesLike(s sample, op operation) bool {
	after := apply(op, s.instr, s.before)
	return slices.Eq(after, s.after)
}

func runProgram(program []instruction, assignments map[int]operation) int {
	state := []int{0, 0, 0, 0}
	for _, instr := range program {
		op := assignments[instr.opCode]
		state = apply(op, instr, state)
	}
	return state[0]
}

func apply(op operation, instr instruction, state []int) []int {
	newState := make([]int, len(state))
	copy(newState, state)
	newState[instr.output] = op.fn(instr, state)
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

func readSamples() []sample {
	samples := make([]sample, 0)
	f, err := os.Open("./samples.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
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

func readProgram() []instruction {
	f, err := os.Open("./program.txt")
	if err != nil {
		panic(err)
	}
	instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		instr := parseInstruction(line)
		instructions = append(instructions, instr)
	}
	return instructions
}

func parseInstruction(str string) instruction {
	parts := slices.Atoi(strings.Split(str, " "))
	return instruction{parts[0], parts[1], parts[2], parts[3]}
}
