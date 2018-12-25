package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	e := 0
	f := e | 65536
	for e = 1765573; ; {
		b := f & 255
		e = e + b
		e = e & 16777215
		e = e * 65899
		e = e & 16777215
		if f < 256 {
			break
		}
		for b = 0; (b+1)*256 <= f; b++ {
		}
		f = b
	}

	fmt.Println(e)

	fmt.Println(time.Since(start))
}
