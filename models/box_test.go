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

func testAutoPlay(t *testing.T) {
	var b BlockBox
	b.TestData2()
	b.GroupPoint()
	b.Print()

	//	var flaglist []int
	g := b.Step()
	for i := 0; i < 100; i++ {
		for _, v := range g {
			v.Step()
		}
	}
}
