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

	maxLen :=0
	maxLenEnd1 :=0
	for i:=0;i<len1+len2;i++ {
		s1Begin := 0
		s2Begin := 0
		tmpMaxlen := 0
		if i<len1 {
			s1Begin = len1 -i
		} else {
			s2Begin = i-len1
		}
		j :=0
		for ;s1Begin + j < len1 && s2Begin +j < len2;j ++ {
			if str1[s1Begin +j] == str2[s2Begin + j] {
				tmpMaxlen ++
			} else {
				if tmpMaxlen > maxLen {
					maxLen = tmpMaxlen
					maxLenEnd1 = s1Begin +j
				} else {
					tmpMaxlen =0
				}
			}
		}
		if tmpMaxlen > maxLen {
			maxLen = tmpMaxlen
			maxLenEnd1 = s1Begin + j
		}
	}

	for i := maxLenEnd1 -maxLen;i< maxLenEnd1;i++ {
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
