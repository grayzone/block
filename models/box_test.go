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

func TestBlockBoxSum(t *testing.T) {
	var b BlockBox
	var x BlockBox
	x.Seed()
	x.Print()
	var y BlockBox
	y.Seed()
	y.Print()
	b.Data = BlockBoxSum(x.Data, y.Data)
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

func TestBlockBoxAdjoin(t *testing.T) {
	var b BlockBox
	b.Seed()
	b.Print()
	b.Format()
	b.Adjoin()
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
	b.Format()
	b.Adjoin()

	for i := 1; i < b.Flag+1; i++ {
		var result BlockBox
		result.Data = b.RemoveGroupBlock(i)
		result.Format()
		result.Adjoin()
		result.Print()
	}
}

func TestAutoPlay(t *testing.T) {
	var b BlockBox
	b.Seed()
	b.Format()
	b.Adjoin()
	b.Print()
	count := 0
	for b.Flag > 0 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		index := r1.Intn(b.Flag) + 1
		fmt.Printf("index:%d, count:%d\n", index, count)
		count++
		var result BlockBox
		result.Data = b.RemoveGroupBlock(index)
		result.Format()
		result.Adjoin()

		result.Print()
		b = result
	}
}
