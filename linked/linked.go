package linked

import (
	"fmt"
	)

type ListNode struct {
	Value int
	Next *ListNode
}

func PrintLinked(node *ListNode)  {
	if node != nil {
		if node.Next != nil {
			PrintLinked(node.Next)
		}
		fmt.Println(node.Value)
		//debug.PrintStack()
	}
}

func PrintLinked2(node *ListNode)  {
	var stack Stack

	pNode := node
	for pNode != nil {
		stack.Push(pNode)
		pNode = pNode.Next
	}

	for !stack.IsEmpty() {
		Node,_ := stack.Top()
		pNode = Node.(*ListNode)
		fmt.Println(pNode.Value)
		stack.Pop()
	}
}