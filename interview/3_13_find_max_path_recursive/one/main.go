package main

import (
	"fmt"
	"github.com/isdamir/gotype"
	"math"
)

type IntRef struct {
	Val int
}

//求a,b,c的最大值
func Max(a,b,c int) int {
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}

	if max > c {
		return max
	} else {
		return c
	}
}

//寻找最长路径
func FindMaxPathRecursive(root *gotype.BNode, max *IntRef) int {
	if root == nil {
		return 0
	}
	//求左子树以root.left为起始节点的最大路径和
	sumLeft := FindMaxPathRecursive(root.LeftChild,max)
	//求右子树以root.right为起始节点的最大路径和
	sumRight := FindMaxPathRecursive(root.RightChild,max)

	//求以root为起始节点，叶子节点为结束节点的最大路径和
	allMax := root.Data.(int) + sumLeft + sumRight
	leftMax := root.Data.(int) + sumLeft
	rightMax := root.Data.(int) + sumRight
	tmpMax := Max(allMax,leftMax,rightMax)
	if tmpMax > max.Val {
		max.Val = tmpMax
	}
	var subMax int
	if sumLeft > sumRight {
		subMax = sumLeft
	} else {
		subMax = sumRight
	}

	//返回以root为起始节点，叶子节点为结束节点的最大路径和
	return root.Data.(int) + subMax
}

func FindMaxPath(root *gotype.BNode) int {
	max := &IntRef{Val: math.MinInt}
	FindMaxPathRecursive(root,max)
	return max.Val
}

func main() {
	data := []int{2,3,5}
	fmt.Println("数组：",data)
	root := gotype.ArrayToTree(data,0, len(data)-1)
	fmt.Println(FindMaxPath(root))
}
