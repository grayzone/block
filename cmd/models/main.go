package main

import (
	"math/rand"
	"time"

	"fmt"

	"github.com/grayzone/block/models"
)

func Test01() {
	var b models.BlockBox
	b.TestData2()
	b.GroupPoint()
	b.Print()

	t0 := time.Now()

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

					x5 := x4[i4].Step()

					for i5 := range x5 {
						x5[i5].PrintFlag()
					}
					//			x4[i4].PrintFlag()
				}
			}
			//	y[j].PrintFlag()
		}
		//	x[i].PrintFlag()
	}

	t1 := time.Now()
	fmt.Printf("spend time:%v\n", t1.Sub(t0))
}

func Test02() {
	t0 := time.Now()
	jobs := make(chan models.BlockBox)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				j.PrintFlag()
			} else {
				fmt.Println("all the jobs are done.")
				done <- true
				return
			}
		}
	}()

	var b models.BlockBox
	b.TestData2()
	b.GroupPoint()
	x1 := b.Step()
	for i1 := range x1 {
		x2 := x1[i1].Step()
		for i2 := range x2 {
			x3 := x2[i2].Step()
			for i3 := range x3 {
				x4 := x3[i3].Step()
				for i4 := range x4 {
					jobs <- x4[i4]
				}

			}

		}
	}

	close(jobs)
	fmt.Println("The task is done.")
	<-done
	t1 := time.Now()
	fmt.Printf("spend time:%v\n", t1.Sub(t0))
}

func Test03() {
	t0 := time.Now()
	var b models.BlockBox
	b.TestData2()
	b.GroupPoint()
	b.Print()
	done := make(chan bool)

	go func() {
		for {
			fmt.Println("Flag:", b.Flag)
			if b.Flag == 0 {
				fmt.Println("done.....")
				done <- true
				return
			}
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			index := r1.Intn(b.Flag) + 1
			tmp := b.OneClick(index)
			tmp.Print()
			b = tmp
			b.Mask = [10][10]int{}
			b.Status = [10][10]int{}
			/*
				b.Data = tmp.Data
				b.Flag = tmp.Flag
				b.FlagList = tmp.FlagList
				b.Mask = [10][10]int{}
				b.Status = [10][10]int{}

			*/

		}
	}()

	<-done
	t1 := time.Now()
	fmt.Printf("spend time:%v\n", t1.Sub(t0))
}

func Test04() {
	t0 := time.Now()
	var b models.BlockBox
	b.TestData2()
	b.GroupPoint()
	b.Print()

	count := b.Flag
	round := 0
	for {

		if count == 0 {
			return
		}
		buttom := b.FindButtonGroupIndex()
		fmt.Println("buttom:", buttom)

		index := 0
		if count > 1 {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			if round > 15 {
				index = r1.Intn(count) + 1
			} else {
				index = r1.Intn(len(buttom))
				index = buttom[index]
			}

		} else {
			index = buttom[0]
		}

		r := b.OneClick(index)
		r.Print()
		b = r
		count = b.Flag
		round++
	}

	t1 := time.Now()
	fmt.Printf("spend time:%v\n", t1.Sub(t0))

}

func main() {
	Test04()

}
