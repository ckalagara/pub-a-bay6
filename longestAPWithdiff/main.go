package main

import "fmt"

func main() {

	a := []int32{8, 1, -1, 0, 3, 6, 2, 4, 5, 7, 9, -5, -7}

	fmt.Printf("Longest %d \n", longAPWithDiff(a, 2))

}

func longAPWithDiff(a []int32, k int32) int32 {

	m := make(map[int32]int)

	lo := a[0]
	hi := a[0]

	for i := 1; i < len(a); i++ {

		m[a[i]] = i

		if a[i] > hi {
			hi = a[i]
		}

		if a[i] < lo {
			lo = a[i]
		}

	}

	fmt.Printf("High %d, Low %d \n", hi, lo)

	var l, lx int32

	i := lo

	cMatch := []int32{}

	// if _, ok := m[i+k]; ok {
	// 	l++

	// 	cMatch = append(cMatch, i)
	// 	i = i + k
	// }

	for i <= hi {

		if _, ok := m[i]; ok {
			l++
			cMatch = append(cMatch, i)
			fmt.Printf("match %v, %d \n", cMatch, l)
			i = i + k
			continue
		} else {
			cMatch = []int32{}
			if lx < l {
				lx = l
			}
			l = 0
			i = i + k
		}
		fmt.Printf("match %v, %d \n", cMatch, l)
	}

	if l > lx {
		return l
	}

	return lx

}
