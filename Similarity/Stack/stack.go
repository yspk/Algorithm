package Stack

import (
	"fmt"
	"log"
)

//数据结构 栈  先进后出，像电梯

type Stack struct {
	size int64 //容量
	top  int64 //栈顶
	data []interface{}
}

func MakeStack(size int64) *Stack {
	q := &Stack{}
	q.size = size
	q.data = make([]interface{}, size)
	return q
}

//入栈，栈顶升高
func (t *Stack) Push(element interface{}) bool {
	if t.IsFull() {
		log.Printf("栈已满，无法完成入栈")
		return false
	}
	t.data[t.top] = element
	t.top++
	return true
}

//出栈，栈顶下降
func (t *Stack) Pop() (r interface{}, err error) {
	if t.IsEmpty() {
		err = fmt.Errorf("栈已空，无法完成出栈")
		log.Println("栈已空，无法完成出栈")
		return
	}
	t.top--
	r = t.data[t.top]
	return
}

//栈长度, 已有容量的长度
func (t *Stack) StackLength() int64 {
	return t.top
}

//清空, 不需要清空值 ，再入栈，覆盖即可
func (t *Stack) Clear() {
	t.top = 0
}

//判空
func (t *Stack) IsEmpty() bool {
	return t.top == 0
}

//判满
func (t *Stack) IsFull() bool {
	return t.top == t.size
}

//遍历
func (t *Stack) Traverse(fn func(node interface{}), isTop2Bottom bool) {
	if isTop2Bottom {
		var i int64 = 0
		for ; i < t.top; i++ {
			fn(t.data[i])
		}
	} else {
		for i := t.top - 1; i >= 0; i-- {
			fn(t.data[i])
		}
	}

}
