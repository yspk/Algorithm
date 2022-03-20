package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func DFS(root *gotype.BNode)  {
	if root == nil {
		return
	}
	stack := gotype.NewSliceStack()
	already := make(map[*gotype.BNode]bool)
	//1.初始化
	//把初始节点放到栈里。如果有多个就都放进去
	//并标记初始节点被访问
	stack.Push(root)
	already[root] = true
	//2.不断访问栈，每次出队一个节点
	for !stack.IsEmpty() {
		node := stack.Pop().(*gotype.BNode)
		fmt.Println(node.Data)
		//3.拓展相邻节点
		//pop出的节点的相邻节点，加入栈并在already中标记
		if node.RightChild != nil && !already[node.RightChild] {
			stack.Push(node.RightChild)
			already[node.RightChild] = true
		}
		if node.LeftChild != nil && !already[node.LeftChild] {
			stack.Push(node.LeftChild)
			already[node.LeftChild] = true
		}
	}


}

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	root := gotype.ArrayToTree(arr,0,len(arr)-1)
	DFS(root)
}
