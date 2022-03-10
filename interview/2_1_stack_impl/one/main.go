package main

import (
	"fmt"
)

type LNode struct {
	Data  int
	Right *LNode
	Down  *LNode
}

func (p *LNode) Insert(headRef *LNode, data int) *LNode {
	newNode := &LNode{Data: data, Down: headRef}
	headRef = newNode
	return headRef
}

//创建链表
func CreateNode() *LNode {
	node := &LNode{}
	node = node.Insert(nil, 31)
	node = node.Insert(node, 8)
	node = node.Insert(node, 6)
	node = node.Insert(node, 3)

	node.Right = node.Insert(node.Right, 21)
	node.Right = node.Insert(node.Right, 11)

	node.Right.Right = node.Insert(node.Right.Right, 50)
	node.Right.Right = node.Insert(node.Right.Right, 22)
	node.Right.Right = node.Insert(node.Right.Right, 15)

	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 55)
	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 40)
	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 39)
	node.Right.Right.Right = node.Insert(node.Right.Right.Right, 30)

	return node
}

//合并有序链表
func merge(a *LNode, b *LNode) *LNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	//把两个链表中较小的节点赋值给result
	result := &LNode{}
	if a.Data < b.Data {
		result = a
		result.Down = merge(a.Down, b)
	} else {
		result = b
		result.Down = merge(a, b.Down)
	}
	return result
}

//把链表扁平化处理
func Flatten(root *LNode) *LNode {
	if root == nil || root.Right == nil {
		return root
	}
	//递归处理root.Right链表
	root.Right = Flatten(root.Right)
	//把root节点对应的链表与右边的链表合并
	root = merge(root, root.Right)
	return root
}

func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	tmp := node
	for tmp != nil {
		fmt.Print(tmp.Data, " ")
		tmp = tmp.Down
	}
	fmt.Println()
}

func main() {
	fmt.Println("链表扁平化")
	head := CreateNode()
	head = Flatten(head)
	PrintNode("扁平化后的链表：", head)
}
