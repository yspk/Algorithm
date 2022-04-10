package main

import (
	"bytes"
	"fmt"
)

func longestPalindrome(s string) string {
	// write your code here
	length := len(s)
	if length == 0 || length == 1 {
		return s
	}

	var longest int = 1
	var start int = 0

	isPaliindrom := make([][]bool, length)
	for i := 0; i < length; i++ {
		isPaliindrom[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		isPaliindrom[i][i] = true
	}

	for i := 0; i < length-1; i++ {
		isPaliindrom[i][i+1] = s[i] == s[i+1]
		if isPaliindrom[i][i+1] {
			longest = 2
			start = i
		}
	}

	for i := length - 1; i >= 0; i-- {
		for j := i + 2; j < length; j++ {
			isPaliindrom[i][j] = isPaliindrom[i+1][j-1] && s[i] == s[j]
			if isPaliindrom[i][j] && j-i+1 > longest {
				longest = j - i + 1
				start = i
			}
		}
	}

	var result bytes.Buffer
	for i := start; i < start+longest; i++ {
		result.WriteByte(s[i]) //字符串中单个字符提取
	}
	return result.String()
}

func main() {
	fmt.Println(longestPalindrome("abb"))
}
