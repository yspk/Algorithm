package main

import (
	"errors"
	"fmt"
)

type SliceQueue struct {
	arr   []int
	front int
	rear  int
}

func (p *SliceQueue) IsEmpty() bool {
	return p.front == p.rear
}

func (p *SliceQueue) Size() int {
	return p.rear - p.front
}

func (p *SliceQueue) GetFront() int {
	if !p.IsEmpty() {
		return p.arr[p.front]
	}
	panic(errors.New("queue is empty"))
}

func (p *SliceQueue) GetBack() int {
	if !p.IsEmpty() {
		return p.arr[p.rear-1]
	}
	panic(errors.New("queue is empty"))
}

func (p *SliceQueue) DeQueue() {
	if !p.IsEmpty() {
		p.rear--
		p.arr = p.arr[1:]
	}
	panic(errors.New("queue is empty"))
}

func (p *SliceQueue) EnQueue(t int) {
	p.arr = append(p.arr, t)
	p.rear++
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误：", err)
		}
	}()

	fmt.Println("Slice 构建队列结构")
	sliceStack := &SliceQueue{arr: make([]int, 0)}
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
	SliceMode()
}
