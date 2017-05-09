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
	Data        [10][10]int
	Mask        [10][10]int
	Status      [10][10]int
	Flag        int
	FlagList    []int
	MaxFlagList []int
	IsClicked   bool
}

func (b BlockBox) Print() {
	b.PrintFlag()
	fmt.Println("-------------Data------------- | --------------Mask------------ | --------------Status------------")
	fmt.Println("|01|02|03|04|05|06|07|08|09|10 @ |01|02|03|04|05|06|07|08|09|10|@ |01|02|03|04|05|06|07|08|09|10|")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			if b.Data[10-j][10-i-1] != 0 {
				fmt.Printf("|%02d", b.Data[10-j][10-i-1])
			} else {
				fmt.Printf("|  ")
			}

		}
		fmt.Print(" @ ")
		for j := 10; j > 0; j-- {
			if b.Mask[10-j][10-i-1] != 0 {
				fmt.Printf("|%02d", b.Mask[10-j][10-i-1])
			} else {
				fmt.Printf("|  ")
			}
		}
		fmt.Print(" @ ")
		for j := 10; j > 0; j-- {
			if b.Status[10-j][10-i-1] != 0 {
				fmt.Printf("|%02d", b.Status[10-j][10-i-1])
			} else {
				fmt.Printf("|  ")
			}
		}
		fmt.Println()
	}
}
func (b BlockBox) PrintFlag() {
	fmt.Printf("flag:%d, flag list:%v, max flag list:%v, is clicked : %v\n", b.Flag, b.FlagList, b.MaxFlagList, b.IsClicked)
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

func (b *BlockBox) TestData2() {
	s1 := rand.NewSource(99)
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Data[i][j] = r1.Intn(5) + 1
		}
	}
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

func (b *BlockBox) down() {
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

func (b *BlockBox) left() {
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
	b.down()
	b.left()
}

func (b *BlockBox) findPointGroup(x, y int) []Point {
	var result []Point
	point := b.Data[x][y]
	if point == 0 {
		return result
	}
	cross := b.findNCrossData(x, y)
	bSingle := (point != cross[0]) && (point != cross[1]) && (point != cross[2]) && (point != cross[3])
	if bSingle {
		return result
	}
	if point == cross[0] {
		if b.Status[x-1][y] == 0 {
			result = append(result, Point{X: x - 1, Y: y})
			b.Status[x-1][y]++
			ps := b.findPointGroup(x-1, y)
			if len(ps) > 0 {
				result = append(result, ps...)
			}
		}

	}
	if point == cross[1] {
		if b.Status[x+1][y] == 0 {
			result = append(result, Point{X: x + 1, Y: y})
			b.Status[x+1][y]++
			ps := b.findPointGroup(x+1, y)
			if len(ps) > 0 {
				result = append(result, ps...)
			}
		}
	}
	if point == cross[2] {
		if b.Status[x][y-1] == 0 {
			result = append(result, Point{X: x, Y: y - 1})
			b.Status[x][y-1]++
			ps := b.findPointGroup(x, y-1)
			if len(ps) > 0 {
				result = append(result, ps...)
			}

		}

	}
	if point == cross[3] {
		if b.Status[x][y+1] == 0 {
			result = append(result, Point{X: x, Y: y + 1})
			b.Status[x][y+1]++
			ps := b.findPointGroup(x, y+1)
			if len(ps) > 0 {
				result = append(result, ps...)
			}
		}
	}
	//	fmt.Printf("x:%d,y:%d,point:%d, cross:%v,group:%v\n", x, y, point, cross, result)
	return result
}

func (b *BlockBox) setPointGroupMask(x, y int) {
	ps := b.findPointGroup(x, y)
	//	fmt.Printf("x:%d,y:%d,group:%d\n", x, y, ps)
	if len(ps) == 0 {
		return
	}
	mask := 0
	for _, v := range ps {
		if b.Mask[v.X][v.Y] != 0 {
			mask = b.Mask[v.X][v.Y]
		}
	}
	if mask == 0 {
		b.Flag++
		for _, v := range ps {
			b.Mask[v.X][v.Y] = b.Flag
		}
	} else {
		for _, v := range ps {
			b.Mask[v.X][v.Y] = mask
		}
	}
}

func (b BlockBox) findNCrossData(x int, y int) [4]int {
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

func (b *BlockBox) GroupPoint() {
	b.Format()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.setPointGroupMask(i, j)
		}
	}
}

func (b BlockBox) findGroupBlock(index int) []Point {
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

func (b BlockBox) FindButtonGroupIndex() []int {
	var result []int
	for i := 1; i < b.Flag+1; i++ {
		points := b.findGroupBlock(i)
		for _, p := range points {
			if p.Y < 5 {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

func (b *BlockBox) RemoveGroupBlock(index int) {
	if index == 0 {
		return
	}
	group := b.findGroupBlock(index)
	for i := range group {
		p := group[i]
		b.Data[p.X][p.Y] = 0
	}
}

func (b BlockBox) OneClick(index int) BlockBox {
	b.RemoveGroupBlock(index)
	b.Format()
	b.FlagList = append(b.FlagList, index)
	b.MaxFlagList = append(b.MaxFlagList, b.Flag)
	var result BlockBox
	result.Data = b.Data
	result.FlagList = b.FlagList
	result.MaxFlagList = b.MaxFlagList
	result.GroupPoint()
	return result
}

func (b *BlockBox) Step() []BlockBox {
	var result []BlockBox
	b.GroupPoint()
	for i := 1; i < b.Flag+1; i++ {
		s := *b
		s.RemoveGroupBlock(i)
		s.Format()
		s.GroupPoint()
		s.FlagList = append(s.FlagList, i)
		result = append(result, s)
	}
	b.IsClicked = true
	return result
}

func (b *BlockBox) Remove(x, y int) {
	b.GroupPoint()
	//	fmt.Printf("x:%d,y:%d,mask index:%d\n", x, y, b.Mask[x][y])
	index := b.Mask[x][y]
	b.RemoveGroupBlock(index)
	b.Format()
	b.Print()
}
