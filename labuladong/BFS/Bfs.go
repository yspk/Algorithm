package BFS

import (
	"container/list"
)

type Node struct {
	val int
	left,right *Node
}

func Create(arr []int,i int) *Node{
	if i >= len(arr){
		return nil
	}
	var t *Node
	if arr[i] != 0 {
		t = &Node{arr[i],nil,nil}
		t.left = Create(arr,2*i+1)
		t.right = Create(arr,2*i+2)
	}else {
		return nil
	}
	return t
}

// 计算从起点 start 到终点 target 的最近距离
func bfs(start *Node,target int) int {
	q := list.New()                 // 核心数据结构
	visited := make(map[*Node]bool) // 避免走回头路

	q.PushBack(start) // 将起点加入队列
	visited[start] = true
	var step = 0 // 记录扩散的步数

	for q.Len() != 0 {
		var sz = q.Len()
		/* 将当前队列中的所有节点向四周扩散 */
		for i := 0; i < sz; i++ {
			var cur = q.Front()
			curNode := cur.Value.(*Node)
			if curNode == nil {
				return 0
			}
			q.Remove(cur)
			/* 划重点：这里判断是否到达终点 */
			if curNode.val == target {
				return step
			}

			/* 将 cur 的相邻节点加入队列 */
			left := curNode.left
			if !visited[left] {
				q.PushBack(left)
				visited[left] = true
			}
			right := curNode.right
			if !visited[right] {
				q.PushBack(right)
				visited[right] = true
			}

		}
		/* 划重点：更新步数在这里 */
		step++
	}
	return step
}
