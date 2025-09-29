package main

import "fmt"

func overlap(aEnd, bStart int) bool {

	if aEnd > bStart {
		return true
	}

	return false
}

func main() {

	a := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{9, 18},
	}

	var b [][]int

	w := a[0]

	for i := 1; i < len(a); i++ {

		if overlap(w[1], a[i][0]) {

			if w[0] > a[i][0] {
				w[0] = a[i][0]
			}

			if w[1] < a[i][1] {
				w[1] = a[i][1]
			}
		} else {

			b = append(b, w, a[i])
			w = a[i]
		}

		fmt.Printf("Iteratio %d: target %v, window %v \n", i, a[i], w)

	}

	fmt.Printf("result %v \n", b)

}
