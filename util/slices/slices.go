package slices

import "math/rand"

func Shuffle(slice []interface{}) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
