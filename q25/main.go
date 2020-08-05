package main

import "fmt"

// LinkNode define a node
type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {
	n1 := LinkNode{Val: 1}
	n2 := LinkNode{Val: 2}
	n3 := LinkNode{Val: 3}
	n4 := LinkNode{Val: 4}
	n5 := LinkNode{Val: 5}
	n6 := LinkNode{Val: 6}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	DisplayLink(&n1)
	nl := ReverseLink(&n1)
	DisplayLink(nl)
}

// ReverseLinkFirstN reverse firse {n} node
func ReverseLinkFirstN(n uint64, head *LinkNode) bool {
	i := n
	if i == 1 {
		return true
	}
	return true
}

// ReverseLink reverse whole link
func ReverseLink(head *LinkNode) *LinkNode {
	var pre *LinkNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}

// DisplayLink print a link's all nodes
func DisplayLink(head *LinkNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val)
		if p.Next != nil {
			fmt.Print(" => ")
		}
		p = p.Next
	}
	fmt.Println("")
}
