package main

import (
	"fmt"
)

func PrintResult(input map[string]string) {
	//用来存储把input的健与值调换后的信息
	reverseInput := make(map[string]string)
	for k, v := range input {
		reverseInput[v] = k
	}
	//找到起点
	start := ""
	for k, _ := range input {
		if _, ok := reverseInput[k]; !ok {
			start = k
			break
		}
	}

	if start == "" {
		fmt.Println("输入不合理")
	} else {
		//从起点出发按照顺序遍历路径
		to := input[start]
		fmt.Print(start, "->", to)
		start = to
		to = input[to]
		for to != "" {
			fmt.Print(",", start, "->", to)
			start = to
			to = input[to]
		}
		fmt.Println()
	}

}

func main() {
	input := map[string]string{
		"西安": "成都",
		"北京": "上海",
		"大连": "西安",
		"上海": "大连",
	}
	PrintResult(input)

}
