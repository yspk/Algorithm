package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func PrintTreeLayer(root *gotype.BNode) {
	if root == nil {
		return
	}
	var p *gotype.BNode
	queue := gotype.NewSliceQueue()
	queue.EnQueue(root)
	for queue.Size() > 0 {
		p = queue.DeQueue().(*gotype.BNode)
		//访问当前节点
		fmt.Print(p.Data, " ")
		//如果这个节点的左孩子不为空则入队列
		if p.LeftChild != nil {
			queue.EnQueue(p.LeftChild)
		}
		if p.RightChild != nil {
			queue.EnQueue(p.RightChild)
		}
	}
}

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Println("数组:", arr)
	root := gotype.ArrayToTree(arr, 0, len(arr)-1)
	fmt.Println("转换成树的层序遍历为：")
	PrintTreeLayer(root)
	fmt.Println()
}
