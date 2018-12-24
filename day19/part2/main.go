package main

import (
	"fmt"
	"time"
)

type instruction struct {
	opName string
	input1 int64
	input2 int64
	output int64
}

type operation = func(instruction, []int64) int64

func main() {
	start := time.Now()

	output := decompiledProgram()
	fmt.Println(output)

	fmt.Println(time.Since(start))
}

func decompiledProgram() int {
	x := 0
	for w := 1; w <= 10551300; w++ {
		if 10551300%w == 0 {
			x = x + w
		}
	}
	return x
}
