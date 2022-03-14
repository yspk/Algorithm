package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//方法功能：获取二叉树从根节点root到node节点的路径
func GetPathFromRoot(root *gotype.BNode,node *gotype.BNode,s *gotype.SliceStack) bool {
	if root == nil {
		return false
	}
	if root.Data.(int) == node.Data.(int) {
		s.Push(root)
		return true
	}
	//如果node节点在root节点的左子树或者右子树上
	//那么root是node的祖先节点，把它加到栈里
	if GetPathFromRoot(root.LeftChild,node,s) || GetPathFromRoot(root.RightChild,node,s) {
		s.Push(root)
		return true
	}
	return false
}

//查找二叉树中两个节点最近的共同父节点
func FindParentNode(root,node1,node2 *gotype.BNode) *gotype.BNode {
	stack1 := gotype.NewSliceStack()
	stack2 := gotype.NewSliceStack()
	//获取从root到node1的路径
	GetPathFromRoot(root,node1,stack1)
	//获取从root到node2的路径
	GetPathFromRoot(root,node2,stack2)
	var commonParent *gotype.BNode
	for t1,t2 := stack1.Pop().(*gotype.BNode),stack2.Pop().(*gotype.BNode); t1 != nil && t2 != nil && t1.Data.(int) == t2.Data.(int);{
		commonParent = t1
		t1,t2 = stack1.Pop().(*gotype.BNode),stack2.Pop().(*gotype.BNode)
	}
	return commonParent
}

func main() {
	data := []int{
		1,2,3,4,5,6,7,8,9,10,
	}
	fmt.Println("数组：", data)
	root := gotype.ArrayToTree(data,0,len(data)-1)
	node1 := root.LeftChild.LeftChild.LeftChild
	node2 := root.LeftChild.RightChild
	result := FindParentNode(root,node1,node2)
	if result != nil {
		fmt.Println(node1.Data,"与",node2.Data,"的最近公共父节点为:",result.Data)
	}else {
		fmt.Println("没有公共父节点")
	}
}
