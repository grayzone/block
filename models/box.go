package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
}

type BlockBox struct {
	Data [10][10]int
	Mask [10][10]int
	Flag int
}

func (b BlockBox) Print() {
	fmt.Println("-------------Data--------------|--------------Mask------------")
	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			if b.Data[10-j][10-i-1] != 0 {
				fmt.Printf("%02d ", b.Data[10-j][10-i-1])
			} else {
				fmt.Printf("   ")
			}

		}
		fmt.Print(" | ")
		for j := 10; j > 0; j-- {
			if b.Mask[10-j][10-i-1] != 0 {
				fmt.Printf("%02d ", b.Mask[10-j][10-i-1])
			} else {
				fmt.Printf("   ")
			}
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
		{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
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

func BlockBoxSum(x, y [10][10]int) [10][10]int {
	var result [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			result[i][j] = x[i][j] + y[i][j]
		}
	}
	return result
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

func (b BlockBox) foundNCrossMask(x int, y int, cross [4]int) [4]int {
	var result [4]int
	point := b.Data[x][y]

	for i := range cross {
		if point == cross[i] {
			result[i] = 1
		}
	}
	return result
}

func (b *BlockBox) SetNCrossMask(x int, y int, mask [4]int, value int) {
	b.Mask[x][y] = value
	if mask[0] != 0 {
		b.Mask[x-1][y] = value
	}
	if mask[1] != 0 {
		b.Mask[x+1][y] = value
	}
	if mask[2] != 0 {
		b.Mask[x][y-1] = value
	}
	if mask[3] != 0 {
		b.Mask[x][y+1] = value
	}
}

func (b *BlockBox) checkNCrossMask(x int, y int, mask [4]int) {
	index := 0
	if mask[0] != 0 {
		if b.Mask[x-1][y] != 0 {
			index = b.Mask[x-1][y]
		}
	}
	if index == 0 {
		if mask[1] != 0 {
			if b.Mask[x+1][y] != 0 {
				index = b.Mask[x+1][y]
			}
		}
	}

	if index == 0 {
		if mask[2] != 0 {
			if b.Mask[x][y-1] != 0 {
				index = b.Mask[x][y-1]
			}
		}
	}

	if index == 0 {
		if mask[3] != 0 {
			if b.Mask[x][y+1] != 0 {
				index = b.Mask[x][y+1]
			}
		}
	}

	if index == 0 {
		b.Flag++
		b.SetNCrossMask(x, y, mask, b.Flag)
	} else {
		b.SetNCrossMask(x, y, mask, index)
	}
}

func (b *BlockBox) checkNCrossData(x, y int) bool {

	point := b.Data[x][y]
	if point == 0 {
		return false
	}
	crossData := b.foundNCrossData(x, y)

	bSingle := (point != crossData[0]) && (point != crossData[1]) && (point != crossData[2]) && (point != crossData[3])
	if bSingle {
		return false
	}

	maskData := b.foundNCrossMask(x, y, crossData)
	//	fmt.Printf("p(%d,%d)=%d cross=%v mask=%v\n", x, y, point, crossData, maskData)
	b.checkNCrossMask(x, y, maskData)

	return true
}

func (b *BlockBox) Adjoin() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.checkNCrossData(i, j)
		}
	}
}

func (b BlockBox) FoundGroupBlock(index int) []Point {
	var result []Point
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if b.Mask[i][j] != index {
				continue
			}
			var p Point
			p.X = i
			p.Y = j
			result = append(result, p)
		}
	}
	return result
}

func (b BlockBox) RemoveGroupBlock(index int) [10][10]int {
	var result BlockBox
	result = b
	group := b.FoundGroupBlock(index)
	for i := range group {
		p := group[i]
		result.Data[p.X][p.Y] = 0
	}
	return result.Data
}

func (b BlockBox) AutoPlay() {
	b.Format()
	b.Adjoin()
	for b.Flag > 0 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		index := r1.Intn(b.Flag) + 1
		fmt.Println("index:", index)
		b.RemoveGroupBlock(index)
		b.Format()
		b.Adjoin()

		b.Print()
	}
}
