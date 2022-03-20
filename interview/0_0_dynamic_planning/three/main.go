package main

import (
	"fmt"
)

// • 有n块石头分别在x轴的0, 1, …, n-1位置
// • 一只青蛙在石头0，想跳到石头n-1
// • 如果青蛙在第i块石头上，它最多可以向右跳距离ai
// • 问青蛙能否跳到石头n-1

func CanJump(a []int) bool {
	length := len(a)
	if length == 0 {
		return false
	}
	f := make([]bool,length)
	f[0] = true

	//f(j) = f(i) && (j-i < a[i]) || f(k) && (j-i <= a[i]) ;0 =< k < j
	for j:=1;j<length;j++ {
		f[j] = false
		for i:=0; i < j; i ++ {
			if f[i] && j-i <= a[i] {
				f[j] = true
				break
			}
		}
	}

	return f[length-1]
}

func main() {
	fmt.Println(CanJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(CanJump([]int{3, 2, 1, 0, 4}))
}
