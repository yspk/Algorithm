package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func moveBottomToTop(s *gotype.SliceStack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		moveBottomToTop(s)
		top2 := s.Pop()
		s.Push(top1)
		s.Push(top2)
	} else {
		s.Push(top1)
	}
}

//翻转栈顺序
func ReverseStack(s *gotype.SliceStack) {
	if s.IsEmpty() {
		return
	}
	//把栈底元素移到栈顶
	moveBottomToTop(s)
	top := s.Pop()
	//递归处理子栈
	ReverseStack(s)
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
	stack := CreateStack([]int{1, 2, 3, 4, 5, 6})
	//PrintStack("创建的栈为：", stack) //出栈的方式打印
	ReverseStack(stack)
	fmt.Println(stack.Size())
	PrintStack("翻转后的栈为：", stack)
}
