package slices

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Shuffle(slice []interface{}) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Atoi(slice []string) []int {
	intSlice := make([]int, len(slice))
	for i, str := range slice {
		j, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Sprintf("invalid int %s", str))
		}
		intSlice[i] = j
	}
	return intSlice
}

func Eq(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func StrContains(slice []string, str string) bool {
	for _, elem := range slice {
		if elem == str {
			return true
		}
	}
	return false
}
