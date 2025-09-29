package main

import "fmt"

func main() {

	a := []int{1, 3, 4, 6, 9, 7}

	i, v := findPeak(a)

	fmt.Printf("Peak in %v, is at %d val %d \n", a, i, v)

}

func findPeak(a []int) (index, peak int) {

	l, r := 0, len(a)-1

	for l < r {
		m := l + (r-l)/2

		if a[m] > a[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l, a[l]
}
