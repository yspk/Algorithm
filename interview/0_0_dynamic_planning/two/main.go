package main

import (
	"fmt"
	"math"
)

//三种硬盘：分别面值2元、5元、7元，每种硬盘都有足够多
//买一本书需要27元
//如何用最少的硬币组合付款

func MinCoinToPay(arr []int,value int) int {
	if value == 0 || len(arr) != 3 {
		return 0
	}
	minValue := make(map[int]int)
	for i := -7; i < 0; i ++ {
		minValue[i] = math.MaxInt
	}
	minValue[0] = 0

	for i:=1;i<= value;i++ {
		tmp := math.MaxInt
		if minValue[i-arr[0]] <= tmp {
			tmp = minValue[i-arr[0]]
		}
		if minValue[i-arr[1]] <= tmp {
			tmp = minValue[i-arr[1]]
		}
		if minValue[i-arr[2]] <= tmp {
			tmp = minValue[i-arr[2]]
		}
		if tmp != math.MaxInt {
			minValue[i] = tmp + 1
		} else {
			minValue[i] = math.MaxInt
		}

	}
	if minValue[value] != math.MaxInt {
		return minValue[value]
	}
	return 0
}

func main() {
	fmt.Println(MinCoinToPay([]int{2,5,7},27))
}
