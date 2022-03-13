package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func CreateNodeT(node *gotype.LNode, get int) *gotype.LNode {
	cur := node
	gotype.CreateNode(node, 8)
	for cur.Next != nil {
		cur = cur.Next
		if cur.Data == get {
			return cur
		}
	}
	return nil
}

//给定单链表中某个节点，删除该节点--并不用删除节点存储空间，修改内容即可
func RemoveNode(node *gotype.LNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	node.Data = node.Next.Data
	node.Next = node.Next.Next
	return true
}

func main() {
	fmt.Println("删除指定节点")
	head := &gotype.LNode{}
	retNode := CreateNodeT(head, 5)
	if retNode == nil {
		fmt.Println("get  error")
		return
	}
	fmt.Printf("删除节点%v前链表\n", retNode.Data)
	result := RemoveNode(retNode)
	if result {
		gotype.PrintNode("删除该节点后的链表", head)
	}
}
