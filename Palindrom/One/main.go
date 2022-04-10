package main

import (
	"bytes"
	"fmt"
)
//Manacher算法:马拉车算法
func longestPalindrome (s string) string {
	// write your code here
	length := len(s)
	if length == 0 || length == 1 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteByte(' ') //起始和结尾也需要添加
	for i :=0;i< length; i ++ {
		buf.WriteByte(s[i])
		buf.WriteByte(' ')
	}
	newS := buf.String()

	var count int
	var index int
	for i:=0;i < len(newS); i ++ {
		var j int
		for j=0;i-j >=0 &&i+j < len(newS);j++ {
			if newS[i-j] == newS[i+j] {
				//if count < 2*j+1 && newS[i-j] != ' ' { //排除' b '这样的记录
				if count < 2*j+1 {
					count = 2*j+1
					index = i
				}
			} else {
				break
			}
		}
	}
	var result bytes.Buffer
	for i := index -count/2;i<= index+count/2;i++ {
		if newS[i] != ' ' {
			result.WriteByte(newS[i]) //字符串中单个字符提取
		}
	}
	return result.String()
}

func main()  {
	fmt.Println(longestPalindrome("abb"))
}
