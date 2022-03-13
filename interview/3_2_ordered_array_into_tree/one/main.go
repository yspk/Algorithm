package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func arrayToTree(arr []int, start, end int) *gotype.BNode {
	var root *gotype.BNode
	if end >= start {
		root = gotype.NewBNode()
		mid := (start + end + 1) / 2
		//树的根节点为数组中间的元素
		root.Data = arr[mid]
		//递归用左半部分数组构建root的左子树
		root.LeftChild = arrayToTree(arr, start, mid-1)
		//递归用右半部分数组构建root的右子树
		root.RightChild = arrayToTree(arr, mid+1, end)
	}
	return root
}

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Println("数组:", arr)
	root := arrayToTree(arr, 0, len(arr)-1)
	fmt.Println("转换成树的中序遍历为：")
	gotype.PrintTreeMidOrder(root)
	fmt.Println()
}
