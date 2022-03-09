package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//判断单链表是否有环
func IsLoop(head *gotype.LNode) *gotype.LNode {
	if head == nil && head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

func FindLoopNode(head *gotype.LNode, meetNode *gotype.LNode) *gotype.LNode {
	first := head.Next
	second := meetNode
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}

func main() {
	fmt.Println("查找环")
	head := &gotype.LNode{}
	gotype.CreateNode(head, 8)
	//构造环
	loop := head.Next
	for loop.Next != nil {
		loop = loop.Next
	}
	loop.Next = head.Next.Next.Next

	meetNode := IsLoop(head)
	if meetNode != nil {
		fmt.Println("有环")
		loopNode := FindLoopNode(head, meetNode)
		fmt.Println("环的入口节点为：", loopNode.Data)
	} else {
		fmt.Println("无环")
	}
}
