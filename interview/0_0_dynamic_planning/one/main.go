package main

import (
	"fmt"
)

//三种硬币：分别面值2元、5元、7元，每种硬盘都有足够多
//买一本书需要27元
//如何用最少的硬币组合付款

func MinCoinToPay(coins []int, amount int) int {
	if amount == 0 || len(coins) == 0 {
		return -1
	}
	minValue := make(map[int]int)
	//for i := 0 - coins[len(coins) -1]; i < 0; i ++ {
	//	minValue[i] = 1000000
	//}
	minValue[0] = 0

	//f(x) = min(f(x-arr[0])+1,f(x-arr[1])+1,f(x-arr[2])+1)
	for i:=1;i<= amount;i++ {
		tmp := 1000000
		for j:=0;j<len(coins);j ++ {
			if i - coins[j] < 0 {
				minValue[i-coins[j]] = 1000000
			}
			if coins[j] !=0 && minValue[i-coins[j]] <= tmp {
				tmp = minValue[i-coins[j]]
			}
		}

		//if minValue[i-coins[1]] <= tmp {
		//	tmp = minValue[i-coins[1]]
		//}
		//if minValue[i-coins[2]] <= tmp {
		//	tmp = minValue[i-coins[2]]
		//}
		if tmp != 1000000 {
			minValue[i] = tmp + 1
		} else {
			minValue[i] = 1000000
		}

	}
	if minValue[amount] != 1000000 {
		return minValue[amount]
	}
	return -1
}

func main() {
	fmt.Println(MinCoinToPay([]int{2,5,7},27))
	fmt.Println(MinCoinToPay([]int{1,2,5},11))
	fmt.Println(MinCoinToPay([]int{0,1,1,1,8},9))
	fmt.Println(MinCoinToPay([]int{21,31,51},91))
}
