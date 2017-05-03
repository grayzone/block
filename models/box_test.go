package models

import "testing"
import "fmt"

func TestBlockBoxPrint(t *testing.T) {
	fmt.Println("TestBlockBoxPrint")
	var b BlockBox
	b.Seed()
	b.Print()
}

func TestBlockBoxDown(t *testing.T) {
	fmt.Println("TestBlockBoxDown")
	var b BlockBox
	b.TestData()
	b.Print()
	fmt.Println("----------------------")
	b.Down()
	b.Print()
}

func TestBlockBoxSum(t *testing.T) {
	fmt.Println("TestBlockBoxSum")
	var b BlockBox
	var x BlockBox
	x.Seed()
	x.Print()
	fmt.Println("----------------------")
	var y BlockBox
	y.Seed()
	y.Print()
	fmt.Println("----------------------")
	b.Sum(x, y)
	b.Print()
}

func TestBlockBoxLeft(t *testing.T) {
	fmt.Println("TestBlockBoxLeft")
	var b BlockBox
	b.TestData()
	b.Print()
	fmt.Println("----------------------")
	b.Left()
	b.Print()
}

func TestBlockBoxFormat(t *testing.T) {
	fmt.Println("TestBlockBoxFormat")
	var b BlockBox
	b.TestData()
	b.Print()
	fmt.Println("----------------------")
	b.Format()
	b.Print()
}
