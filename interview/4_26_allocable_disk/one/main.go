package main

import "fmt"

func isAllocable(d,p []int) bool {
	//磁盘分区下标
	dIndex := 0
	for _,v := range p {
		//先找到符合条件的磁盘
		for dIndex < len(d) && v > d[dIndex] {
			dIndex ++
		}
		//没有可用的磁盘
		if dIndex >= len(d) {
			return false
		}
		//给分区分配磁盘
		d[dIndex] -= v
	}
	return true
}

func main() {
	//磁盘
	d := []int{120,120,120}
	//分区
	p := []int{60,60,80,20,80}
	if isAllocable(d,p) {
		fmt.Println("分配成功")
	} else {
		fmt.Println("分配失败")
	}
}
