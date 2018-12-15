package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type treeNode struct {
	children []*treeNode
	metadata []int
}

func (node *treeNode) value() int {
	if len(node.children) == 0 {
		return arraySum(node.metadata)
	}
	sum := 0
	for _, childRef := range node.metadata {
		i := childRef - 1
		if i < len(node.children) {
			sum += node.children[i].value()
		}
	}
	return sum
}

func main() {
	start := time.Now()

	head := readInput()
	sum := head.value()

	fmt.Println(sum)

	fmt.Println(time.Since(start))
}

func arraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func readInput() *treeNode {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := unsafeArrayAtoi(strings.Split(input, " "))
	return parseTree(parts)
}

func parseTree(parts []int) *treeNode {
	tn, _ := parseTreeAux(parts)
	return tn
}

func parseTreeAux(parts []int) (*treeNode, int) {
	if len(parts) == 0 {
		return nil, 0
	}

	numberOfChildren := parts[0]
	numberOfMetadata := parts[1]

	children := make([]*treeNode, numberOfChildren)

	offset := 2
	for i := 0; i < numberOfChildren; i++ {
		childNode, length := parseTreeAux(parts[offset:])
		children[i] = childNode
		offset += length
	}

	length := offset + numberOfMetadata

	metadata := parts[offset:length]

	return &treeNode{children, metadata}, length
}

func unsafeArrayAtoi(arr []string) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i] = unsafeAtoi(v)
	}
	return newArr
}

func unsafeAtoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
