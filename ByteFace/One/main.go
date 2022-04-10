package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func main()  {
	//root := &gotype.LNode{}
	//gotype.CreateNode(root,7) //root的数据为空，指向后续的链表
	//reverse(root)
	//root2 := &gotype.LNode{}
	//gotype.CreateNode(root2,1) //root的数据为空，指向后续的链表
	//reverse(root2)
	//root3 := &gotype.LNode{}
	//gotype.CreateNode(root3,2) //root的数据为空，指向后续的链表
	//reverse(root3)
	root4 := &gotype.LNode{}
	gotype.CreateNode(root4,5) //root的数据为空，指向后续的链表
	reverse(root4)
}

func reverse(root *gotype.LNode) *gotype.LNode {
	if root == nil || root.Next == nil {
		return nil
	}

	mid := root.Next
	end := root.Next
	//找到中间节点
	for end != nil {
		if end.Next != nil {
			end = end.Next.Next
			mid = mid.Next
		} else {
			end = end.Next
		}
	}
	var cur *gotype.LNode
	var start = &gotype.LNode{} 	//第二个单链表的起始指针

	start.Next = mid.Next
	if mid.Next == nil {
		for root.Next != nil {
			fmt.Println(root.Next.Data)
			root = root.Next
		}
	}
	if start == nil || start.Next == nil {
		return nil
	}
	//var cur *gotype.LNode       //定义当前结点
	var next *gotype.LNode      //后继结点
	cur = start.Next.Next //从链表的第二个结点开始
	start.Next.Next = nil //链表的第一个结点为尾结点
	//遍历的结点依次插入到头结点的后面
	for cur != nil {
		next = cur.Next      //保存后继结点
		cur.Next = start.Next //放到头结点后面
		start.Next = cur
		cur = next
	}
	mid.Next = nil
	for root.Next != nil && start.Next != nil {
		fmt.Println(root.Next.Data)
		fmt.Println(start.Next.Data)
		root = root.Next
		start = start.Next
	}
	for root.Next != nil {
		fmt.Println(root.Next.Data)
		root = root.Next
	}
	return nil
}
