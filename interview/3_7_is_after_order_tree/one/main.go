package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

var pHead *gotype.BNode
var pEnd *gotype.BNode

//方法功能：判断一个数组是否是二元查找树的后续遍历序列
func IsAfterOrdeer(arr []int, start, end int) bool {
	if arr == nil {
		return false
	}
	//数组最后一个节点必定是根节点
	root := arr[end]
	var i, j int
	//找到第一个大于root的值，那么前面所有的节点都位于root的左子树上
	for i = start; i < end; i++ {
		if arr[i] > root {
			break
		}
	}
	//如果序列是后续遍历的序列，那么从i开始的所有值都应该大于根节点root的值
	for j = i; j < end; j++ {
		if arr[j] < root {
			return false
		}
	}
	leftIsAfterOrder := true
	rightIsAfterOrder := true
	//判断小于root值的序列是否是某一二元查找树的后序遍历
	if i > start {
		leftIsAfterOrder = IsAfterOrdeer(arr, start, i-1)
	}
	//判断大于root值的序列是否是某一二元查找树的后序遍历
	if j < end {
		rightIsAfterOrder = IsAfterOrdeer(arr, i, end-1)
	}
	return leftIsAfterOrder && rightIsAfterOrder
}

func main() {
	data := []int{
		1, 3, 2, 5, 7, 6, 4,
	}
	fmt.Print("数组：", data)
	result := IsAfterOrdeer(data, 0, len(data)-1)
	if result {
		fmt.Println("是某一二元查找树的后序遍历序列")
	}

}
