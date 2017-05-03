package util

import (
	"math/rand"
	"strconv"
	"time"

	"strings"

	"github.com/astaxie/beego"
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

func DropZeroData(s string) [10][10]int {
	array := strings.Split(s, ",")
	var result [10][10]int
	index := 0
	beego.Debug("input:", s)
	for i := 0; i < 10; i++ {
		noZero := 0
		for j := 0; j < 10; j++ {
			a := array[index]
			index++
			num, _ := strconv.Atoi(a)
			if num != 0 {
				result[i][noZero] = num
				noZero++
			}
		}
	}
	return result
}

func IsEmptyArray(a [10]int) bool {
	sum := 0
	for _, v := range a {
		sum += v
	}
	if sum == 0 {
		return true
	}
	return false
}

func LeftData(input [10][10]int) [10][10]int {
	var result [10][10]int
	index := 0
	for i := 0; i < 10; i++ {
		a := input[i]
		beego.Debug("a=", a)
		if IsEmptyArray(a) {
			continue
		}
		result[index] = a
		index++
	}
	return result
}
