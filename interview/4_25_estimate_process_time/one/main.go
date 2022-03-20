package main

import (
	"fmt"
)

func CalculateProcessTime(t []int,n int) []int {
	if t == nil || n <= 0 {
		return nil
	}
	m := len(t)
	proTime := make([]int,m)
	for i:=0;i<n;i++ {
		//把任务给第j个机器上后这个机器的执行时间
		minTime := proTime[0] + t[0]
		//把任务给第minIndex个机器上
		minIndex := 0
		for j:=1;j<m;j++ {
			//分配到第j台机器上后执行时间更短
			if minTime > proTime[j] + t[j] {
				minTime = proTime[j] + t[j]
				minIndex = j
			}
		}
		proTime[minIndex] += t[minIndex]
	}
	return proTime
}

func main() {
	t := []int{7,10}
	proTime := CalculateProcessTime(t,6)
	if proTime == nil {
		fmt.Println("分配失败")
		return
	}
	totalTime := proTime[0]
	for i,v := range proTime {
		fmt.Printf("第%v台服务器有%v个任务，执行总时间为：%v",i+1,v/t[i],v)
		if v > totalTime {
			totalTime = v
		}
		fmt.Println()
	}
	fmt.Println("执行完成所有任务所需的时间为：",totalTime)
}
