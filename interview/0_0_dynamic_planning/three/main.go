package main

import (
	"fmt"
)

//给定m行n列的网格，有一个机器人从左上角（0，0）出发，没补可以向下或者向右走一步
//问有多少种不同的方式走到右下角

func CountPath(m,n int) int {
	if m == 0 || n == 0 {
		return -1
	}
	minValue := make(map[[2]int]int)
	for i := 0; i < m; i ++ {
		minValue[[2]int{i,0}] = 1
	}
	for j := 0; j<n; j++ {
		minValue[[2]int{0,j}] = 1
	}

	//f(x,y) = f(x-1,y) + f(x,y-1)
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {
			minValue[[2]int{i,j}] = minValue[[2]int{i-1,j}] + minValue[[2]int{i,j-1}]
		}
	}

	return minValue[[2]int{m-1,n-1}]
}

func main() {
	fmt.Println(CountPath(3,1))
}
