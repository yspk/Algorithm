package main

import (
	"fmt"
)

var pos = 0
//以arr[low]为基准把数组分成两部分
func partition(arr []int,low,high int)  {
	key := arr[low]
	for low < high {
		for low < high && arr[high] > key {
			high --
		}
		arr[low] = arr[high]
		for low < high && arr[low] < key {
			low ++
		}
		arr[high] = arr[low]
	}
	arr[low] = key
	pos = low
}


func getMid(arr []int) int {
	low := 0
	n := len(arr)
	high := n-1
	mid := (low + high)/2
	for true {
		//以arr[low]为基准把数组分成两部分
		partition(arr,low,high)
		if pos == mid { //找到中位数
			break
		} else if pos > mid { //继续在右半部分查找
			high = pos -1
		} else {//继续在左半部分查找
			low = pos + 1
		}
	}
	//如果数组长度是奇数，中位数为中间的元素，否则就是中间两个数的平均值
	if n%2 != 0 {
		return arr[mid]
	} else {
		return (arr[mid] + arr[mid+1])/2
	}

}

func main() {
	a := []int{7,5,3,1,11,9}
	fmt.Println(getMid(a))
}
