
package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func BFS(root *gotype.BNode)  {
	if root == nil {
		return
	}
	queue := gotype.NewSliceQueue()
	already := make(map[*gotype.BNode]bool)
	//1.初始化
	//把初始节点放到queue里。如果有多个就都放进去
	//并标记初始节点被访问
	queue.EnQueue(root)
	already[root] = true
	//2.不断访问队列，每次出队一个节点
	for !queue.IsEmpty() {
		node := queue.DeQueue().(*gotype.BNode)
		fmt.Println(node.Data)
		//if node.Data.(int) == 5 {
		//	fmt.Println("5的左孩子是：",node.LeftChild.Data)
		//}
		//3.拓展相邻节点
		//pop出的节点的相邻节点，加入队列并在already中标记
		if node.LeftChild != nil && !already[node.LeftChild] {
			queue.EnQueue(node.LeftChild)
			already[node.LeftChild] = true
		}
		if node.RightChild != nil && !already[node.RightChild] {
			queue.EnQueue(node.RightChild)
			already[node.RightChild] = true
		}
	}
}

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	root := gotype.ArrayToTree(arr,0,len(arr)-1)
	gotype.PrintTreeLayer(root)
	BFS(root)
}
