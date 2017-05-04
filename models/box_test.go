package models

import "testing"

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
}
