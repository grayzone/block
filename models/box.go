package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type BlockBox struct {
	Data [10][10]int
}

func (b BlockBox) Print() {
	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			fmt.Printf("%d ", b.Data[10-j][10-i-1])
		}
		fmt.Println()
	}
}

func (b *BlockBox) Seed() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Data[i][j] = r1.Intn(5) + 1
		}
	}
}

func (b *BlockBox) TestData() {
	b.Data = [10][10]int{
		{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 2, 3, 4, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}}

}

func (b *BlockBox) Parse(s string) {
	array := strings.Split(s, ",")
	index := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			a := array[index]
			num, _ := strconv.Atoi(a)
			b.Data[i][j] = num
			index++
		}
	}
}

func (b *BlockBox) Down() {
	var tmp [10][10]int
	for i := 0; i < 10; i++ {
		noZero := 0
		for j := 0; j < 10; j++ {
			num := b.Data[i][j]

			if num != 0 {
				tmp[i][noZero] = num
				noZero++
			}
		}
	}
	b.Data = tmp
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

func (b *BlockBox) Left() {
	var tmp [10][10]int
	index := 0
	for i := 0; i < 10; i++ {
		a := b.Data[i]
		if isEmptyArray(a) {
			continue
		}
		tmp[index] = a
		index++
	}
	b.Data = tmp
}

func (b *BlockBox) Format() {
	b.Down()
	b.Left()
}

func (b *BlockBox) Sum(x, y BlockBox) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Data[i][j] = x.Data[i][j] + y.Data[i][j]
		}
	}
}

func (b BlockBox) foundNCrossData(x int, y int) [4]int {
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
		result[0] = b.Data[x-1][y]
	}
	if result[1] != -1 {
		result[1] = b.Data[x+1][y]
	}
	if result[2] != -1 {
		result[2] = b.Data[x][y-1]
	}
	if result[3] != -1 {
		result[3] = b.Data[x][y+1]
	}
	return result
}
