package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func Reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre *gotype.LNode
	var cur *gotype.LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func main() {
	head := &gotype.LNode{}
	fmt.Println("就地逆序")
	gotype.CreateNode(head, 100)
	gotype.PrintNode("逆序前：", head)
	Reverse(head)
	gotype.PrintNode("逆序后：", head)
}
