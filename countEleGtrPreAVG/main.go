package main

import "fmt"

func main() {
	a := []int{100, 200, 150, 300, 10}

	l := 0
	preSum := 0
	for i := 1; i < len(a); i++ {

		preSum += a[i-1]
		avg := preSum / i

		fmt.Printf("avg %d current %d \n", avg, a[i])

		if avg < a[i] {
			l++
		}
	}
	fmt.Printf("Count : %d \n", l)
}
