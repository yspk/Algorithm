package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func moveBottomToTopSort(s *gotype.SliceStack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		moveBottomToTopSort(s)
		top2 := s.Top()
		if top1.(int) > top2.(int) {
			s.Pop()
			s.Push(top1)
			s.Push(top2)
		}
	}
	s.Push(top1)
}

//对栈排序
func SortStack(s *gotype.SliceStack) {
	if s.IsEmpty() {
		return
	}

	moveBottomToTopSort(s)
	top := s.Pop()
	SortStack(s)
	s.Push(top)
}

//快速创建一个栈
func CreateStack(list []int) *gotype.SliceStack {
	stack := gotype.NewSliceStack()
	for _, v := range list {
		stack.Push(v)
	}
	return stack
}

func PrintStack(str string, s *gotype.SliceStack) {
	fmt.Print(str)
	for !s.IsEmpty() {
		fmt.Print(s.Pop(), " ")
	}
	fmt.Println()
}

func main() {
	stack := CreateStack([]int{6, 5, 4, 3, 2, 1, 7, 8, 10, 9})
	//PrintStack("创建的栈为：", stack) //出栈的方式打印
	SortStack(stack)
	fmt.Println(stack.Size())
	PrintStack("排序后的栈为：", stack)
}
