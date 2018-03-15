package main

import (
	"./matrix"
	"fmt"
)

func main() {
	//str1, str2 := "cnblogs", "belong"
	//str1, str2 := "acgbfhk", "cegefkh"
	str1, str2 := "abcbdab", "bdcaba"
	m := matrix.NewMatrix(len(str1)+1, len(str2)+1)
	flag := matrix.NewMatrix2(len(str1)+1, len(str2)+1)
	MostSameStringLength(str1, str2, m, flag)
	fmt.Println("当前最大公共子序列的长度为:", m.Get(len(str1)+1, len(str2)+1))
	SubSequence(len(str1), len(str2), str1, str2, flag)
}

func MostSameStringLength(str1, str2 string, m *matrix.Matrix, flag *matrix.Matrix2) {
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				m.Set(i+1, j+1, m.Get(i, j)+1)
				flag.Set(i+1, j+1, "left_up")
			} else {
				if m.Get(i, j+1) >= m.Get(i+1, j) {
					m.Set(i+1, j+1, m.Get(i, j+1))
					flag.Set(i+1, j+1, "left")
				} else {
					m.Set(i+1, j+1, m.Get(i+1, j))
					flag.Set(i+1, j+1, "up")
				}
			}
		}
	}
}

func SubSequence(i, j int, str1, str2 string, flag *matrix.Matrix2) {
	if i == 0 || j == 0 {
		return
	}
	if flag.Get(i+1, j+1) == "left_up" {
		fmt.Println(string(str2[j-1]), ": 当前坐标:", i, ",", j)
		SubSequence(i-1, j-1, str1, str2, flag)
	} else if flag.Get(i+1, j+1) == "left" {
		SubSequence(i-1, j, str1, str2, flag)
	} else {
		SubSequence(i, j-1, str1, str2, flag)
	}
}
