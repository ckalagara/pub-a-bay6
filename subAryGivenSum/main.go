package main

import "fmt"

func main() {

	a := []int{2, -1, 2, 1, -2, 3}
	k := 3
	m := 2

	for l := 0; l < len(a)-1; l++ {
		rSum := 0
		maxEle := a[l]
		for r := l; r < len(a); r++ {
			rSum = rSum + a[r]

			if maxEle < a[r] {
				maxEle = a[r]
			}

			if rSum == k && maxEle <= m {
				fmt.Printf("Subset %v , maxEle %d \n", a[l:r+1], maxEle)
			}

		}
	}

}
