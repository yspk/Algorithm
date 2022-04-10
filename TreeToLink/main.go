package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

var pre *gotype.BNode
/**
* Definition for a binary tree node.

*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * @param root: root of a tree
 * @return: head node of a doubly linked list
 */
//func treeToDoublyList (root *TreeNode) *TreeNode {
//	// Write your code here.
//	if root == nil {
//		return nil
//	}
//
//
//
//
//}

func dfsTree(node *gotype.BNode) {
	if node == nil {
		return
	}

	if node.LeftChild != nil {
		dfsTree(node.LeftChild)
	}
	fmt.Println(node.Data)
	//TODO
	if pre != nil {
		pre.RightChild = node
		node.LeftChild = pre
	}
	pre = node
	if node.RightChild != nil {
		dfsTree(node.RightChild)
	}
}

func main()  {
	root := gotype.ArrayToTree([]int{1,2,3,4,5},0,4)
	dfsTree(root)
	for pre != nil {
		fmt.Println(pre.Data)
		pre = pre.LeftChild
	}
}
