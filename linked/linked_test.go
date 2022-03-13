package linked

import "testing"

func TestPrintLinked(t *testing.T) {
	//var tempNode ListNode
	var temp *ListNode
	var max int = 2048000
	for i:=max;i>=1;i-- {
		var Node ListNode
		if i == max {
			Node = ListNode{i,nil}
			temp = &Node
		} else {
			Node = ListNode{i,temp}
			temp = &Node
		}
	}
	PrintLinked(temp)
	//PrintLinked2(temp)

}
