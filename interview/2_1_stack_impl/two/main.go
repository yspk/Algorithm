package main

import (
	"errors"
	"fmt"
	"github.com/isdamir/gotype"
)

type LinkStack struct {
	head *gotype.LNode
}

func (p *LinkStack) IsEmpty() bool {
	return p.head.Next == nil
}

func (p *LinkStack) Size() int {
	size := 0
	node := p.head.Next
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

func (p *LinkStack) Top() int {
	if p.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	return p.head.Next.Data.(int)
}

func (p *LinkStack) Pop() int {
	tmp := p.head.Next
	if tmp != nil {
		p.head.Next = tmp.Next
		return tmp.Data.(int)
	}
	panic(errors.New("stack is empty"))
}

func (p *LinkStack) Push(t int) {
	node := &gotype.LNode{Data: t, Next: p.head.Next}
	p.head.Next = node
}

func LinkMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误：", err)
		}
	}()

	fmt.Println("链表构建栈结构")
	sliceStack := &LinkStack{head: &gotype.LNode{}}
	sliceStack.Push(1)
	fmt.Println("栈顶元素为：", sliceStack.Top())
	fmt.Println("栈大小为：", sliceStack.Size())
	sliceStack.Pop()
	fmt.Println("弹栈成功：", sliceStack.Size())
	sliceStack.Pop()
}

func main() {
	LinkMode()
}
