package conv

import (
	"fmt"
	"strconv"
)

func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("invalid int %s", str))
	}
	return i
}
