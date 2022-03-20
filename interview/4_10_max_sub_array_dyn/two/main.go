package main

import (
	"fmt"
)

func maxSubArrayDyn(arr []int) int {
	if arr == nil || len(arr) <+ 0 {
		return -1
	}
	n := len(arr)
	//End := make([]int,n) //包含arr[n-1]的最大子数组和
	//All := make([]int,n) //最大子数组和
	//End[n-1] = arr[n-1]
	//All[n-1] = arr[n-1]
	//包含arr[n-1]的最大子数组和
	nEnd := arr[0]
	//最大子数组和
	nAll := arr[0]
	// All[i-1] = max(End[i-1],arr[i-1],All[i-2])
	for i:=1;i<n;i++ {
		nEnd = Max(nEnd+arr[i],arr[i])
		nAll = Max(nEnd,nAll)
	}
	return nAll
}

func Max(abs int, abs2 int) int {
	if abs > abs2 {
		return abs
	}
	return abs2
}

func main() {
	a := []int{1,-2,4,8,-4,7,-1,-5}
	fmt.Println("优化存储复杂度的动态规划法")
	fmt.Println(maxSubArrayDyn(a))
}
