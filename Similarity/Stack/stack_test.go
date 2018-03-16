package Stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := int64(4)
	q := MakeStack(s)
	if q.size != s {
		t.Error("初始长度%d与创建的栈长度%d不一致", s, q.size)
	}
	if q.top != 0 {
		t.Error("栈顶值不正确，衩始化时，栈顶应为0，实际为%d", q.top)
	}
	if !q.IsEmpty() {
		t.Error("栈应为空，实际为非空")
	}
	if q.IsFull() {
		t.Error("新创建的栈，不应为满")
	}
	if q.StackLength() != 0 {
		t.Error("初始长度应为0，但现在是%d", q.StackLength())
	}
}
func TestStack_Push(t *testing.T) {
	s := int64(4)
	q := MakeStack(s)
	q.Push(10)
	if q.StackLength() != 1 {
		t.Error("长度应为1，但现在是%d", q.StackLength())
	}
	q.Push(12)
	q.Push(14)
	q.Push(16)
	v := q.Push(18)
	if v == true {
		t.Error("满的栈不能再加入，应返回假")
	}
	if q.StackLength() != s {
		t.Error("长度应为%s，但现在是%d", s, q.StackLength())
	}
	q.Clear()
	if q.StackLength() != 0 {
		t.Error("长度应为0，但现在是%d", q.StackLength())
	}
}
func TestStack_Pop(t *testing.T) {
	s := int64(4)
	q := MakeStack(s)
	q.Push(10)
	q.Push(12)
	q.Push(14)
	q.Push(16)
	v, err := q.Pop()
	if err != nil {
		t.Errorf("执行Stack_Pop报错, %s", err)
	}
	if v != 16 {
		t.Errorf("此时栈顶应为16, 但取到的值为%d", v)
	}
	v, err = q.Pop()
	if v != 14 {
		t.Errorf("此时栈顶应为14, 但取到的值为%d", v)
	}
	if q.StackLength() != s-2 {
		t.Errorf("长度应为%d，但现在是%d", s-2, q.StackLength())
	}
}

//测试遍历
func ExampleStack_Traverse() {
	s := int64(4)
	q := MakeStack(s)
	q.Push(10)
	q.Push(12)
	q.Push(14)
	q.Push(16)
	q.Traverse(func(node interface{}) {
		fmt.Println(node)
	}, false)

	q.Pop()
	q.Traverse(func(node interface{}) {
		fmt.Println(node)
	}, true)
	//output:
	//16
	//14
	//12
	//10
	//10
	//12
	//14
}

func ExampleStackDemo1() {

	//栈应用一，十进制转换
	/**
	*   短除法进制转换规则
		N = (N div d)  * d + N mod d    (div ，除， Mod ,求余）
		十进制转8进制
		N       N div 8     N mod 8
		1348    168         4
		168     21          0
		21      2           5
		2       0           2

		N       N div 16    N mod 16
		1348    84          4
		84      5           4
		5       0           5
	*
	**/
	const BINARY int = 2
	const OCTONARY int = 8
	const HEXADECIMAL int = 16

	var fn_c = func(n int, target int) {
		var stack = MakeStack(40)
		for {
			if n == 0 {
				break
			}
			stack.Push(n % target)
			n = n / target
		}
		stack.Traverse(func(node interface{}) {
			fmt.Print(node)
		}, false)
	}

	fn_c(1348, OCTONARY)
	fmt.Println()
	fn_c(1348, HEXADECIMAL)
	fmt.Println()
	fn_c(1348, BINARY)

	//output:
	// 2504
	// 544
	// 10101000100
}
