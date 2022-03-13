package Point

type ListNode struct {
	val int
	next *ListNode
}

//判断单链表中是否有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast,slow := head,head //初始化快慢节点指向头节点
	for fast != nil && fast.next != nil {
		fast = fast.next.next //快指针每次前进两步
		slow = slow.next //慢指针每次前进一步
		if fast == slow {
			return true
		}
	}
	return false
}

//已知单链表中含有环，返回这个环的起始位置
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast,slow := head,head //初始化快慢节点指向头节点
	for fast != nil && fast.next != nil {
		fast = fast.next.next //快指针每次前进两步
		slow = slow.next //慢指针每次前进一步
		if fast == slow {
			break
		}
	}
	slow = head
	for slow != fast {
		//两指针以相同的速度前进
		fast = fast.next
		slow = slow.next
	}
	//两个指针相遇的那个单链表节点就是环的起点
	return slow
}

//寻找无环单链表的中点
func middleCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast,slow := head,head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	//slow就在中间位置
	return slow
}

//二分搜索
func binarySearch(nums []int, target int) int {
	//左、右指针在数组的两端初始化
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return  -1
}

