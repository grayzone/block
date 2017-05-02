package util

import (
	"math/rand"
	"time"
)

func GetSeedData() [10][10]int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var result [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			result[i][j] = r1.Intn(5) + 1
		}
	}
	return result
}
