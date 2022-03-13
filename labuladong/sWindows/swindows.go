package sWindows

import (
	"fmt"
	"math"
)

func slidingWindow(s,t string) string {
	var need,window map[byte]int
	need,window = make(map[byte]int),make(map[byte]int)
	for i:=0;i<len(t);i++ {
		need[t[i]]++
	}
	var left,right,valid int
	//记录最小覆盖子串的起始索引及长度
	var start int
	var length int = math.MaxInt64
	for right < len(s) {
		//c是将移入窗口的字符
		c := s[right]
		//右移窗口
		right ++
		//进行窗口内数据的一系列更新
		//...
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid ++
			}
		}
		fmt.Println(left,right)

		//判断左侧窗口是否要搜索
		for valid == len(need) {
			//在这里更新最小覆盖子串
			if right - left < length {
				start = left
				length = right - left
			}
			//d是将移除窗口的字符
			d := s[left]
			//左移窗口
			left ++
			//进行窗口内数据的一系列更新
			//...
			if need[d] > 0 {
				if window[d] == need [d] {
					valid --
				}
				window[d] --
			}
		}
	}
	if length == math.MaxInt64 {
		return "None"
	}  else {
		return s[start:start+length]
	}
}


