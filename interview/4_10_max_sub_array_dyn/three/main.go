package main

import (
	"fmt"
	"math"
)

func maxSubArrayEx(arr []int) (maxSum,begin,end int){
	if arr == nil || len(arr) <+ 0 {
		return -1,-1,-1
	}
	//子数组最大值
	maxSum = math.MinInt
	//包含最后一位的最大子数组和
	nSum := 0
	nStart := 0
	for i,v := range arr {
		if nSum < 0 {
			nSum = v
			nStart = i //如果某个值使得nSum < 0,那么要从arr[i]重新开始求和
		} else {
			nSum += v
		}
		if nSum > maxSum {
			maxSum = nSum
			begin = nStart
			end = i
		}
	}
	return
}

func Max(abs int, abs2 int) int {
	if abs > abs2 {
		return abs
	}
	return abs2
}

func main() {
	a := []int{1,-2,4,8,-4,7,-1,-5}
	fmt.Println("拓展题")
	fmt.Println(maxSubArrayEx(a))
}
