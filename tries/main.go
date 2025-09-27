package main

import "fmt"

type node struct {
	childs [26]*node
	val    string
	end    bool
}

func getNode(char string) *node {
	n := new(node)
	n.end = false
	n.val = char
	for i, _ := range n.childs {
		n.childs[i] = nil
	}
	return n
}

type trie struct {
	head *node
	sz   int
}

func (t *trie) insert(v string) {
	if t.head == nil {
		t.head = getNode("")
	}

	c := t.head

	for _, k := range v {

		index := k - 'a'
		// char entry has no node
		if c.childs[index] == nil {
			n := getNode(string(k))
			c.childs[index] = n
			c = n
			continue
		}

		c = c.childs[index]
	}
	c.end = true

}

func (t *trie) search(v string) bool {
	c := t.head

	for _, k := range v {
		index := k - 'a'

		if c.childs[index] == nil {
			// not found
			return false
		}

		c = c.childs[index]
	}

	if c != nil && c.end == true {
		// word found
		return true
	}
	return false
}

func main() {
	t := trie{}

	t.insert("card")
	f := t.search("car")

	fmt.Printf("Found %v \n", f)
}
