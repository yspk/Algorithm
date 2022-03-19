package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func findSingle(arr []int) int {
	len := len(arr)
	for i:=0;i<32;i++ {
		count1 := 0
		count0 := 0
		result1 := 0
		result0 := 0
		for j:=0;j<len;j++ {
			//判断数字n的二进制数从右往左数第i位是否为1
			if gotype.IsOne(arr[j],i) {
				//第i位为1的值异或操作
				result1 ^= arr[j]
				//第i位为1的数字个数
				count1 ++
			} else {
				//第i位为1的值异或操作
				result0 ^= arr[j]
				//第i位为1的数字个数
				count0 ++
			}
		}
		//bit值为1的子数组元素个数为奇数，且出现1次的数字被分配到bit值为0
		//的子数组说明只有一个出现1次的数字被分配到bit值为1的子数组中
		//异或结果就是这个出现一次的数字
		if count1%2 == 1 && result0 != 0 {
			return result1
		}
		if count0%2 == 1 && result1 != 0 {
			return result0
		}

	}
	return -1

}

func main() {
	a := []int{6,3,4,5,9,4,3}
	fmt.Println(findSingle(a))
}
