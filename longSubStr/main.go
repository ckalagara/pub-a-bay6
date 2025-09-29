package main

import "fmt"

func main() {
	s := "abcabcbb"
	set := make(map[rune]int)
	l := 0
	lenMx := 0

	for r, v := range s {

		sV, ok := set[v]

		if !ok {
			// not found
			set[v] = r
			cl := r - l + 1

			if cl > lenMx {
				lenMx = cl
			}
			continue
		}

		// found
		l = sV + 1
		set[v] = r

	}

	fmt.Printf("max length %d \n", lenMx)
}
