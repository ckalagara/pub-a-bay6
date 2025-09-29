package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8.....

func main() {

	n := 10

	pa := 0
	pb := 1

	out := make([]int, 10)
	out[0] = pa
	out[1] = pb

	for i := 2; i < n; i++ {
		nxt := pa + pb
		out[i] = nxt
		pa = pb
		pb = nxt
	}
	fmt.Printf("fibonacci %v \n", out)
}
