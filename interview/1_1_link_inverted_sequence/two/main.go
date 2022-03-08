package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func InsertReverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var cur *gotype.LNode
	var next *gotype.LNode
	cur = node.Next.Next
	node.Next.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

func ReversePrint(node *gotype.LNode) {
	if node == nil {
		return
	}
	ReversePrint(node.Next)
	fmt.Print(node.Data, " ")
}

func main() {
	head := &gotype.LNode{Data: 0}
	fmt.Println("插入法")
	gotype.CreateNode(head, 100)
	ReversePrint(head)
	gotype.PrintNode("逆序前：", head)
	InsertReverse(head)
	gotype.PrintNode("逆序后：", head)
}
