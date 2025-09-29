package main

import "fmt"

const (
	debug = false
)

func main() {
	a := []int{-5, 3, 2, 1, -3, -3, 7, 2, 2}
	fmt.Printf("Input 		%v \n", a)
	// bubbleSort(copyIt(a))
	// insertionSort(copyIt(a))
	// selectionSort(copyIt(a))
	mergeSort(copyIt(a))
}

// vedio ref: greg hoog | sorting:.....

func copyIt(a []int) []int {
	dst := make([]int, len(a))
	copy(dst, a)
	return dst
}

func mergeSort(a []int) []int {
	r := mergeSortCore(a)
	fmt.Printf("mergeSort 	%v \n", r)
	return r
}

func mergeSortCore(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	m := len(a) / 2

	la := mergeSortCore(a[:m])
	ra := mergeSortCore(a[m:])
	if debug {
		fmt.Printf("mergesort %v %v \n", la, ra)
	}
	l, r := 0, 0

	sa := []int{}

	for l < len(la) && r < len(ra) {
		if la[l] <= ra[r] {
			sa = append(sa, la[l])
			l++
		} else {
			sa = append(sa, ra[r])
			r++
		}
	}

	sa = append(sa, la[l:]...)
	sa = append(sa, ra[r:]...)

	return sa
}

func selectionSort(a []int) []int {
	ite := 0

	n := len(a)

	for i := 0; i < n-1; i++ {
		ite++
		target := a[i]

		for j := i + 1; j < n; j++ {
			ite++
			currentLower := i
			if target > a[j] {
				currentLower = j
				target = a[j]
			}

			if currentLower != i {
				a[i], a[currentLower] = a[currentLower], a[i]
			}

		}
	}
	fmt.Printf("Selection sort %v, iterations %d \n", a, ite)
	return a
}

func insertionSort(a []int) []int {
	ite := 0
	n := len(a)
	for i := 1; i < n; i++ {
		ite++
		// compare 0, 1
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
			if debug {
				fmt.Printf("Insertionsort - iteration %d._, %v %v \n", i, a[:i+1], a[i+1:])
			}
			for j := i - 1; j > 0; j-- {

				ite++
				if a[j] < a[j-1] {
					a[j], a[j-1] = a[j-1], a[j]
					if debug {
						fmt.Printf("Insertionsort - iteration %d.%d, %v %v \n", i, j, a[:i+1], a[i+1:])
					}
					continue
				}

				break

			}
		}

	}

	fmt.Printf("Insertion sort %v, iterations %d \n", a, ite)
	return a
}

// bubbleSort -- each outer loop iteration pushes the largest number to the end
// In next out loop we can skip seven l(a)-l-1
func bubbleSort(a []int) []int {

	ite := 0

	for l := 0; l < len(a); l++ {
		ite++

		for r := 0; r < len(a)-l-1; r++ {
			ite++
			if a[r] > a[r+1] {
				a[r], a[r+1] = a[r+1], a[r]
			}
			if debug {
				fmt.Printf("Bubblesort - Iteration %d.%d, %v \n", l, r, a)
			}

		}

	}

	fmt.Printf("Bubblesort -  %v , iterations %d\n", a, ite)
	return a
}
