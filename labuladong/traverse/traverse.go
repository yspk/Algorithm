package traverse

//数组遍历
func arrTraverse(arr []int)  {
	for i:= 0;i < len(arr); i++ {
		//迭代访问 arr
	}
}

//链表遍历
type ListNode struct {
	val  int
	next *ListNode
}

func listTraverse(head *ListNode) {
	if head == nil {
		return
	}
	for p := head; p != nil; p = p.next {
		// 迭代访问 p.val
	}
}

func listTraverseRecursion(head *ListNode) {
	if head == nil {
		return
	}
	// 递归访问 head.val
	listTraverseRecursion(head.next)
}

//二叉树遍历
type TreeNode struct {
	val int
	left, right *TreeNode
}

func treeTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	//前序遍历
	treeTraverse(root.left)
	//中序遍历
	treeTraverse(root.right)
	//后序遍历
}

//N叉树遍历
type NTreeNode struct {
	val int
	child []*NTreeNode
}

func nTreeTraverse(root *NTreeNode) {
	if root == nil {
		return
	}
	for _,v := range root.child {
		nTreeTraverse(v)
	}
}



