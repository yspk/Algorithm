package main

import (
	"fmt"
)

func reverse(arr []int,start,end int)  {
	for start < end {
		//tmp := arr[start]
		//arr[start] = arr[end]
		//arr[end] = tmp
		arr[start] ^= arr[end]
		arr[end] ^= arr[start]
		arr[start] ^= arr[end]
		start ++
		end --
	}
}

func rightShift3(arr []int,k int)  {
	if arr == nil {
		fmt.Println("参数错误")
		return
	}
	ll := len(arr)
	k %= ll
	reverse(arr,0,ll-k-1)
	reverse(arr,ll-k,ll-1)
	reverse(arr,0,ll-1)
}




func main() {
	a := []int{1,2,3,4,5,6,7,8}
	rightShift3(a,4)
	fmt.Println("数组循环移位的结果",a)
}
