package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func IsPopSerial(push, pop []int) bool {
	if len(push) == 0 || len(pop) == 0 || len(push) != len(pop) {
		return false
	}

	stack := *gotype.NewSliceStack()

	tmp := -1
	i := 0
	j := 0
	for ; i < len(pop); i++ {
		if !stack.IsEmpty() && stack.Top() == pop[i] {
			stack.Pop()
			continue
		}

		for j = tmp + 1; j < len(push); j++ {
			if push[j] != pop[i] {
				stack.Push(push[j])
			} else {
				tmp = j
				break
			}
		}

		if tmp != j {
			return false
		}

	}

	return true
}

func main() {
	push := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//pop := []int{3, 2, 5, 4, 1}
	//pop := []int{5, 3, 4, 1, 2}
	pop := []int{3, 2, 1, 5, 4, 7, 6, 8, 9, 10}

	if IsPopSerial(push, pop) {
		fmt.Println(pop, "is the right pop Serial to", push)
	} else {
		fmt.Println(pop, "is not the right pop Serial to", push)
	}
}
