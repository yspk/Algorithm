package main

import (
	"fmt"
	"math"
)

func bestMatrixChain(p []int,n int) int {
	//申请数组来保存中间结果，为了简单不使用m[0][0]
	//cost[i,j] = 计算A[i]*A[i+1]*...*A[j]
	//所需的标量乘法的最小数量
	//其中A[i]的维度数是p[i-1]*p[i]
	cost := make([][]int,n)
	for i:= range cost {
		cost[i] = make([]int,n)
	}
	//Len表示矩阵链的长度
	for cLen := 2;cLen < n ; cLen ++ {
		for i:= 1;i<n-cLen+1;i++ {
			j:=i+cLen -1
			cost[i][j] = math.MaxInt
			for k:=i;k<=j-1;k++ {
				//计算乘法运算的代价
				q := cost[i][k] + cost[k+1][j] + p[i-1]*p[k]*p[j]
				if q < cost[i][j] {
					cost[i][j] = q
				}
			}
		}
	}
	return cost[1][n-1]
}

func main() {
	a := []int{1,5,2,4,6}
	fmt.Println("动态规划方法：")
	fmt.Println("最少的乘法次数为：",bestMatrixChain(a,len(a)))
}
