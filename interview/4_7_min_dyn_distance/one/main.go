package main

import (
	"fmt"
	"math"
)

func minDynDistance(arr []int,num1,num2 int) int {
	if arr == nil || len(arr) <= 0 {
		return math.MaxInt
	}
	lastPos1 := -1	//上次遍历到num1的位置
	lastPos2 := -1	//上次遍历到num2的位置
	minDis := math.MaxInt
	for i,v := range arr {
		if v == num1 {
			lastPos1 = i
			if lastPos2 >= 0 {
				minDis = min(minDis,lastPos1-lastPos2)
			}
		}

		if v == num2 {
			lastPos2 = i
			if lastPos1 >= 0 {
				minDis = min(minDis,lastPos2-lastPos1)
			}
		}
	}
	return minDis
}

func min(dis int, i int) int {
	if dis > i {
		return i
	} else {
		return dis
	}
}

func main() {
	arr := []int{4,5,6,4,7,4,6,4,7,8,5,6,4,3,10,8}
	num1 := 4
	num2 := 8
	fmt.Println("动态规划法")
	fmt.Println(minDynDistance(arr,num1,num2))
}
