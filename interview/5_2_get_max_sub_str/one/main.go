package main

import (
	"bytes"
	"fmt"
)

//方法功能：获取两个字符串的最长公共子串
func getMaxSubStr(str1,str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var buf bytes.Buffer
	maxI := 0 //记录最长公共子串最后一个字符的位置
	max := 0
	//申请新的空间来记录公共子串长度信息
	M := make([][]int,len1+1)
	for i := range M {
		M[i] = make([]int, len2 + 1)
	}
	//通过利用递归公式填写新建的二维数组（公共子串的长度信息）
	for i:=1;i<len1+1;i++ {
		for j:=1;j<len2+1;j++ {
			if str1[i-1] == str2[j-1] {
				M[i][j] = M[i-1][j-1] + 1
				if M[i][j] > max {
					max = M[i][j]
					maxI = i
				}
			} else {
				M[i][j] = 0
			}
		}
	}
	for i := maxI -max;i< maxI;i++ {
		buf.WriteByte(str1[i]) //字符串中单个字符提取
	}
	return buf.String()
}

func main() {
	str1 := "abccade"
	str2 := "dgcadde"

	fmt.Println("动态规划法")
	fmt.Println(getMaxSubStr(str1,str2))
}
