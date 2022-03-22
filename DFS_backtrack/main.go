package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := gotype.ArrayToTree(arr, 0, len(arr)-1)
	DFS(root)
}

func DFS(root *gotype.BNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	DFS(root.LeftChild)
	DFS(root.RightChild)
}
