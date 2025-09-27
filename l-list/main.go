package main

import "fmt"

type node struct {
	val  int
	next *node
}

type lList struct {
	head, tail *node
	length     int
}

func (ll *lList) print() {
	fmt.Printf("length %d \n", ll.length)
	if ll.head == nil {
		return
	}

	nxt := ll.head
	for {
		fmt.Println(nxt.val)
		if nxt.next == nil {
			break
		}
		nxt = nxt.next
	}
}

func (ll *lList) addhead(v int) {
	// add head if missing
	if ll.head == nil {
		n := &node{
			val: v,
		}
		ll.head = n
		ll.tail = n
		ll.length++
		return
	}

	n := &node{
		val:  v,
		next: ll.head,
	}

	ll.head = n
	ll.length++

}

func (ll *lList) add(val int) {
	if ll.head == nil {
		ll.addhead(val)
		return
	}

	n := &node{
		val: val,
	}

	ll.tail.next = n
	ll.tail = n
	ll.length++
}

func (ll *lList) delete(v int) {
	// no nodes
	if ll.length == 0 {
		return
	}

	// head match
	if ll.head.val == v {
		ll.head = ll.head.next
		ll.length--
		return
	}

	p := ll.head
	c := ll.head.next

	for {

		if c == nil {
			return
		}

		if c.val == v {
			p.next = c.next
			ll.length--
			return
		}
		p = c
		c = c.next
	}
}

func main() {
	ll := lList{}
	ll.add(10)
	ll.add(20)
	ll.add(30)
	ll.addhead(5)
	ll.delete(20)

	ll.print()
}
