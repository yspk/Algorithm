package main

import (
	"errors"
	"fmt"
	"github.com/isdamir/gotype"
)

type LinkQueue struct {
	head *gotype.LNode
	end  *gotype.LNode
}

func (p *LinkQueue) IsEmpty() bool {
	return p.head == nil
}

func (p *LinkQueue) Size() int {
	size := 0
	node := p.head
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

func (p *LinkQueue) GetFront() int {
	if !p.IsEmpty() {
		return p.head.Data.(int)
	}
	panic(errors.New("queue is empty"))
}

func (p *LinkQueue) GetBack() int {
	if !p.IsEmpty() {
		return p.end.Data.(int)
	}
	panic(errors.New("queue is empty"))
}

func (p *LinkQueue) DeQueue() {
	if !p.IsEmpty() {
		p.head = p.head.Next
		if p.head == nil {
			p.end = nil
		}
	}
	panic(errors.New("queue is empty"))
}

func (p *LinkQueue) EnQueue(t int) {
	node := &gotype.LNode{Data: t}
	if p.IsEmpty() {
		p.head = node
		p.end = node
	} else {
		p.end.Next = node
		p.end = node
	}
}

func LinkMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误：", err)
		}
	}()

	fmt.Println("链表构建队列结构")
	sliceStack := &LinkQueue{}
	sliceStack.EnQueue(1)
	sliceStack.EnQueue(2)
	fmt.Println("队列头元素为：", sliceStack.GetFront())
	fmt.Println("队列尾元素为：", sliceStack.GetBack())
	fmt.Println("队列大小为：", sliceStack.Size())
	sliceStack.DeQueue()
	fmt.Println("出列成功：", sliceStack.Size())
	sliceStack.DeQueue()
}

func main() {
	LinkMode()
}
