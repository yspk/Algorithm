package main

import "fmt"

func BinarySearch(nums []int,target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums)
	for ;start + 1 < end; {
		mid := start + (end - start)/2
		if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

func main()  {
	fmt.Println(BinarySearch([]int{1,3,5,7,9,10,12,14,17,18},11))
}
