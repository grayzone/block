package models

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBlockBoxPrint(t *testing.T) {
	var b BlockBox
	b.Seed()
	b.Print()
}

func TestBlockBoxDown(t *testing.T) {
	var b BlockBox
	b.TestData()
	b.Print()
	b.Down()
	b.Print()
}

func TestBlockBoxLeft(t *testing.T) {
	var b BlockBox
	b.TestData()
	b.Print()
	b.Left()
	b.Print()
}

func TestBlockBoxFormat(t *testing.T) {
	var b BlockBox
	b.TestData()
	b.Print()
	b.Format()
	b.Print()
}

func TestBlockBoxGroupPoint01(t *testing.T) {
	var b BlockBox
	b.Seed()
	b.Print()
	b.GroupPoint()
	b.Print()
	for i := 1; i < b.Flag+1; i++ {
		group := b.FoundGroupBlock(i)
		t.Logf("(%d):%v\n", i, group)
	}
}

func TestBlockBoxGroupPoint02(t *testing.T) {
	var b BlockBox
	b.Data = [10][10]int{
		{2, 1, 4, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 1, 3, 0, 0, 0, 0, 0, 0},
		{1, 5, 2, 1, 2, 1, 0, 0, 0, 0},
		{5, 3, 1, 4, 5, 2, 0, 0, 0, 0},
		{1, 5, 4, 2, 4, 0, 0, 0, 0, 0},
		{1, 2, 2, 2, 0, 0, 0, 0, 0, 0},
		{5, 3, 3, 5, 3, 5, 0, 0, 0, 0},
		{4, 4, 1, 2, 1, 2, 2, 5, 0, 0},
		{4, 4, 2, 4, 3, 2, 5, 5, 3, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	b.Print()
	b.GroupPoint()
	b.Print()
	for i := 1; i < b.Flag+1; i++ {
		group := b.FoundGroupBlock(i)
		t.Logf("(%d):%v\n", i, group)
	}
}

func TestRemoveGroupBlock(t *testing.T) {
	var b BlockBox
	b.Seed()
	b.Print()
	b.GroupPoint()

	for i := 1; i < b.Flag+1; i++ {
		b.RemoveGroupBlock(i)
		b.GroupPoint()
		b.Print()
	}
}

func TestRandomPlay(t *testing.T) {
	var b BlockBox
	b.Seed()
	b.GroupPoint()
	b.Print()
	count := 0
	flag := b.Flag
	for flag > 0 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		index := r1.Intn(flag) + 1
		result := b.OneClick(index)
		flag = result.Flag
		count++
		fmt.Printf("index:%d, count:%d flag:%d\n", index, count, flag)
		result.Print()
		b.Data = result.Data
		b.Mask = [10][10]int{}
		b.Status = [10][10]int{}
	}
}

func TestAutoPlay(t *testing.T) {
	var b BlockBox
	b.TestData2()
	b.GroupPoint()
	b.Print()

	// 1
	x1 := b.Step()
	for i1 := range x1 {
		// 2
		x2 := x1[i1].Step()
		for i2 := range x2 {
			// 3
			x3 := x2[i2].Step()
			for i3 := range x3 {
				// 4
				x4 := x3[i3].Step()
				for i4 := range x4 {

					//	x5 := x4[i4].Step()
					x4[i4].PrintFlag()

				}
			}
			//	y[j].PrintFlag()
		}
		//	x[i].PrintFlag()
	}
}
