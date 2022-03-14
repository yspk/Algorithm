package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func FindParentNodeReverse(root,node1,node2 *gotype.BNode) *gotype.BNode {
	if root == nil || root.Data.(int) == node1.Data.(int) || root.Data.(int) == node2.Data.(int) {
		return root
	}
	lChild := FindParentNodeReverse(root.LeftChild,node1,node2)
	rChild := FindParentNodeReverse(root.RightChild,node1,node2)
	//root的左子树中没有节点node1和node2，那么一定在root的右子树上
	if lChild == nil {
		return rChild
	} else if rChild == nil {
		return lChild
	} else { //node1和node2分别位于root的左子树和右子树上，root就是它们最近的共同父节点
		return root
	}
}

func main() {
	data := []int{
		1,2,3,4,5,6,7,8,9,10,
	}
	fmt.Println("数组：", data)
	root := gotype.ArrayToTree(data,0,len(data)-1)
	node1 := root.LeftChild.LeftChild.LeftChild
	node2 := root.LeftChild.RightChild
	result := FindParentNodeReverse(root,node1,node2)
	if result != nil {
		fmt.Println(node1.Data,"与",node2.Data,"的最近公共父节点为:",result.Data)
	}else {
		fmt.Println("没有公共父节点")
	}
}
