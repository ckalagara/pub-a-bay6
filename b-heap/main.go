package main

import "fmt"

func main() {

	h := BHeap{
		data: []int{8, 11, 19, 15, 22, 23, 30, 35, 17, 25},
	}
	h.remove()
	h.add(6)

}

type BHeap struct {
	data []int
}

func findLeft(i int) int {
	return 2*i + 1
}

func findRight(i int) int {
	return 2*i + 2
}

func findParent(i int) int {
	return (i - 1) / 2
}

// add
// add the element to the last
// compare the parent if small swap other wise exit
func (h *BHeap) add(v int) {
	h.data = append(h.data, v)

	if len(h.data) == 1 {
		return
	}

	ti := len(h.data) - 1

	for {
		// compare tail with it parent
		pi := findParent(ti)

		if h.data[ti] < h.data[pi] {
			// swap with parent
			h.swap(ti, pi)
			// parentIndex as target
			ti = pi
		} else {
			break
		}
	}
	fmt.Printf("Updated %v \n", h.data)
}

func (h *BHeap) swap(x, y int) {
	h.data[x], h.data[y] = h.data[y], h.data[x]
}

// remove
// serve the first item in heap and remove
// take last element form the bottom and add to heap
// sort by comparing left & right and continue looping until we reach it
func (h *BHeap) remove() {
	fmt.Printf("before %v \n", h.data)
	// take the last element and add to front
	ti := len(h.data) - 1
	h.data[0] = h.data[ti]
	h.data = h.data[:ti]

	l := len(h.data) - 1

	ti = 0

	for {
		fmt.Printf(">> %v \n", h.data)

		li := findLeft(ti)
		ri := findRight(ti)

		leftExist := li <= l
		rightExist := ri <= l

		if rightExist && leftExist {
			// parent with both childs

			leftSmall := h.data[li] < h.data[ti]
			rightSmall := h.data[ri] < h.data[ti]

			if leftSmall && rightSmall {

				// compare left right, find which smallest node
				if h.data[li] < h.data[ri] {
					// swap with left
					h.swap(ti, li)
					ti = li
					continue

				} else {
					// swap with right
					h.swap(ti, ri)
					ti = ri
					continue
				}
			}

			if leftSmall {
				// swap with left
				h.swap(ti, li)
				ti = li
				continue
			}

			if rightSmall {
				// swap with right
				h.swap(ti, ri)
				ti = ri
			}

		} else if leftExist {
			// parent with just left

			if h.data[li] < h.data[ti] {
				// swap with left
				h.swap(ti, li)
				ti = li
				continue
			}

		} else {
			// no child nodes
			break
		}

	}

	fmt.Printf("removed %v \n", h.data)
}

// 1
// 8     2
// 3 4   5 6
// 7

// 1
// 3	 	2
// 8 4		 5 6
// 7
