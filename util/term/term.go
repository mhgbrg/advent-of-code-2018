package term

import "fmt"

const (
	Black int = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func Color(color int) {
	fmt.Printf("\033[0;%dm", color)
}

func ClearLine() {
	fmt.Print("\033[K")
}

func ClearScreen() {
	fmt.Print("\033[2J\033[0;0H")
}

func Goto(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}
