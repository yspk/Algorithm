package main

import (
	"bytes"
	"fmt"
)

var keys = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func main() {
	fmt.Println(*letterCombinations("289786573"))
}

func letterCombinations(digits string) *[]string {
	combinations := make([]string, 0)
	//特殊情况处理
	if len(digits) == 0 {
		return &combinations
	}
	//0代表dfs的起点为为digits的下标0位置
	//[]表示combinations的初始值空
	dfs(digits, 0, "", &combinations)
	return &combinations
}

//递归三要素之一：递归的定义
//digists代表输入的数字
//index当前代表输入数字的下标
//combination代表到目前为止得到的组合
//combinations代表到目前为止得到的完整组合
func dfs(digits string, index int, combination string, combinations *[]string) {
	//递归三要素之三：递归的出口
	//已经找到给定数字可以表示的一组combination.记录答案，并立即返回
	if index == len(digits) {
		*combinations = append(*combinations, combination)
		return
	}
	digit := digits[index] - '0'

	//递归三要素之二：递归的拆解
	for i := 0; i < len(keys[digit]); i++ {
		var buf bytes.Buffer
		buf.WriteByte(keys[digit][i])
		dfs(digits, index+1, combination+buf.String(), combinations)
		//这里为什么没有回溯，把之前加入combination的字母移除？
		//因为传入下层的是combination+keys[i]，combination并没有改变
	}
}
