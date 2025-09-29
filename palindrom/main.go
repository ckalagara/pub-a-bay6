package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "A1234567898765432b2B!8765346576885754a"

	pallidromV1(a)

	fmt.Printf("pali v2 %v \n", pallidromV2(a))
}

func pallidromV2(a string) bool {

	l := 0
	r := len(a) - 1
	res := false
	fmt.Printf("Input %s \n", a)
	for l < r {

		// if l > r {
		// 	break
		// }
		lv := rune(a[l])
		rv := rune(a[r])

		if !unicode.IsLetter(lv) {
			l++
			continue
		}
		if !unicode.IsLetter(rv) {
			r--
			continue
		}

		fmt.Printf("%d %s <> %d %s \n", l, string(lv), r, string(rv))

		if unicode.ToLower(lv) != unicode.ToLower(rv) {
			return false
		}
		res = true
		l++
		r--
	}
	return res
}

func pallidromV1(a string) {
	t := []rune{}

	for _, rn := range a {

		if !unicode.IsLetter(rn) {
			continue
		}

		t = append(t, unicode.ToLower(rn))
	}

	r := false
	n := len(t) - 1
	for i := 0; i < len(t)/2; i++ {

		if t[i] == t[n] {
			r = true
			n--
			continue
		}

		r = false

		break
	}

	fmt.Printf("palli %s %b \n", string(t), r)
}
