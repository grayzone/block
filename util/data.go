package util

import (
	"fmt"
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

func isEmptyArray(a [10]int) bool {
	sum := 0
	for _, v := range a {
		sum += v
	}
	if sum == 0 {
		return true
	}
	return false
}

func foundNCrossData(x int, y int, input [10][10]int) [4]int {
	var result [4]int
	if x == 0 {
		result[0] = -1
	}
	if x == 9 {
		result[1] = -1
	}
	if y == 0 {
		result[2] = -1
	}
	if y == 9 {
		result[3] = -1
	}
	if result[0] != -1 {
		result[0] = input[x-1][y]
	}
	if result[1] != -1 {
		result[1] = input[x+1][y]
	}
	if result[2] != -1 {
		result[2] = input[x][y-1]
	}
	if result[3] != -1 {
		result[3] = input[x][y+1]
	}
	return result
}

func checkNCrossData(x int, y int, input [10][10]int) ([10][10]int, bool) {
	var result [10][10]int
	crossData := foundNCrossData(x, y, input)
	point := input[x][y]

	if (point != crossData[0]) && (point != crossData[1]) && (point != crossData[2]) && (point != crossData[3]) {
		return result, false
	}
	result[x][y] = 1
	if point == crossData[0] {
		result[x-1][y] = 1
	}
	if point == crossData[1] {
		result[x+1][y] = 1
	}
	if point == crossData[2] {
		result[x][y-1] = 1
	}
	if point == crossData[3] {
		result[x][y+1] = 1
	}
	return result, true
}

func AdjoinData(input [10][10]int) [10][10]int {
	var result [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			point := input[i][j]
			var x1, x2, y1, y2 int

			if i == 0 {
				x1 = -1
			}
			if i == 9 {
				x2 = -1
			}
			if j == 0 {
				y1 = -1
			}
			if j == 9 {
				y2 = -1
			}
			if x1 != -1 {
				x1 = input[i-1][j]
			}
			if x2 != -1 {
				x2 = input[i+1][j]
			}
			if y1 != -1 {
				y1 = input[i][j-1]
			}
			if y2 != -1 {
				y2 = input[i][j+1]
			}

			fmt.Printf("point:%d, x1:%d, x2:%d, y1:%d, y2:%d\n", point, x1, x2, y1, y2)

		}
	}

	return result
}

func LeftData(input [10][10]int) [10][10]int {
	var result [10][10]int
	index := 0
	for i := 0; i < 10; i++ {
		a := input[i]
		fmt.Println("a=", a)
		if isEmptyArray(a) {
			continue
		}
		result[index] = a
		index++
	}
	return result
}
