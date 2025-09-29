package main

import (
	"fmt"
)

func main() {
	a := []int{-1, 5, -1, 4, -1}
	smMx := a[0]
	smCr := a[0]
	var s, e, l int

	for r := 1; r < len(a); r++ {
		v := a[r]

		if v > smCr+v {
			l = r
			smCr = v
		} else {
			smCr = smCr + v
		}

		if smCr > smMx {
			smMx = smCr
			s = l
			e = r
		}

	}
	fmt.Printf("sub array wit cur sun %d max sum %d > %v\n", smCr, smMx, a[s:e+1])

}
