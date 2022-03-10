package main

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	arr       []int
	stackSize int
}

func (p *SliceStack) IsEmpty() bool {
	return p.stackSize == 0
}

func (p *SliceStack) Size() int {
	return p.stackSize
}

func (p *SliceStack) Top() int {
	if p.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	return p.arr[p.stackSize-1]
}

func (p *SliceStack) Pop() int {
	if p.stackSize > 0 {
		p.stackSize--
		ret := p.arr[p.stackSize]
		p.arr = p.arr[:p.stackSize]
		return ret
	}
	panic(errors.New("stack is empty"))
}

func (p *SliceStack) Push(t int) {
	p.arr = append(p.arr, t)
	p.stackSize++
}

func SliceMode() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误：", err)
		}
	}()

	fmt.Println("Slice 构建栈结构")
	sliceStack := &SliceStack{arr: make([]int, 0), stackSize: 0}
	sliceStack.Push(1)
	fmt.Println("栈顶元素为：", sliceStack.Top())
	fmt.Println("栈大小为：", sliceStack.Size())
	sliceStack.Pop()
	fmt.Println("弹栈成功：", sliceStack.Size())
	sliceStack.Pop()
}

func main() {
	SliceMode()
}
