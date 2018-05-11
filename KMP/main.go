package main

import (
	"fmt"
	"./kmp"
)

func main()  {
	//zstr := "ababcabababdc"
	//mstr := "babdc"
	zstr := "ababfjl310cababfjl310abdc"
	mstr := "fjl310"
	index := kmp.KMP(zstr,mstr)
	if index == -1 {
		fmt.Println("没有匹配的字符串！")
	}else {
		fmt.Println("哈哈，找到字符啦，位置为：" , index)
	}
}

