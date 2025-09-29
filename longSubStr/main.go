package main

import "fmt"

func main() {
	s := "abcabcbb"
	set := make(map[rune]int)
	l := 0
	lenMx := 0

	for r, v := range s {

		pi, ok := set[v]
		if !ok {
			set[v] = r
			cl := r - l + 1
			if cl > lenMx {
				lenMx = cl
				fmt.Printf("scan unique - %v \n", s[l:r+1])
			}
			continue
		}

		set[v] = r
		if pi >= l {
			l = pi + 1
		}
		fmt.Printf("scan %v \n", s[l:r+1])

	}

	fmt.Printf("max length %d \n", lenMx)
}
