// Implemented solution with strings first, later refactored to using a linked
// list in an attempt to make the solution more efficient. This gave a slight
// performance gain (~1 s) per react.

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

func main() {
	start := time.Now()

	p := readInput()
	react(&p)
	fmt.Println(p.length)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

type polymer struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	unit rune
	prev *node
	next *node
}

func (n node) String() string {
	return string(n.unit)
}

func (poly polymer) String() string {
	str := ""
	for n := poly.head; n != nil; n = n.next {
		str += n.String()
	}
	return str
}

func (poly *polymer) append(unit rune) {
	n := node{unit, poly.tail, nil}
	if poly.head == nil {
		poly.head = &n
	} else {
		poly.tail.next = &n
	}
	poly.tail = &n
	poly.length++
}

func (poly *polymer) remove(n *node) {
	if n == poly.tail && n == poly.head {
		poly.head = nil
		poly.tail = nil
	} else if n == poly.head {
		poly.head = n.next
		poly.head.prev = nil
	} else if n == poly.tail {
		poly.tail = n.prev
		poly.tail.next = nil
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
	poly.length--
}

func react(poly *polymer) {
	for n := poly.head; n.next != nil; {
		unit1 := n.unit
		unit2 := n.next.unit
		if reacts(unit1, unit2) {
			poly.remove(n)
			poly.remove(n.next)
			n = poly.head
		} else {
			n = n.next
		}
	}
}

func reacts(unit1, unit2 rune) bool {
	return reactsAux(unit1, unit2) || reactsAux(unit2, unit1)
}

func reactsAux(unit1, unit2 rune) bool {
	return unicode.IsUpper(unit1) &&
		unicode.IsLower(unit2) &&
		unicode.ToLower(unit1) == unit2
}

func readInput() polymer {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	return parsePolymer(str)
}

func parsePolymer(str string) polymer {
	poly := polymer{}
	for _, unit := range str {
		poly.append(unit)
	}
	return poly
}
