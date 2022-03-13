package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func printAtLevel(root *gotype.BNode, level int) int {
	if root == nil || level < 0 {
		return 0
	} else if level == 0 {
		fmt.Print(root.Data, " ")
		return 1
	} else {
		return printAtLevel(root.LeftChild, level-1) + printAtLevel(root.RightChild, level-1)
	}
}

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Println("数组:", arr)
	root := gotype.ArrayToTree(arr, 0, len(arr)-1)
	fmt.Println("转换成树的层序遍历为：")
	printAtLevel(root, 0)
	printAtLevel(root, 1)
	printAtLevel(root, 2)
	printAtLevel(root, 3)
	fmt.Println()
}
