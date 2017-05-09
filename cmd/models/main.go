package main

import (
	"time"

	"fmt"

	"github.com/grayzone/block/models"
)

func main() {
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
