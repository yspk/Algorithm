package kmp

import "fmt"

func KMP(bigstr,smallstr string) int {
	i := 0
	j := 0 //模式串下标
	lb:= len(bigstr)
	ls := len(smallstr)
	next := GetNextVal(smallstr)
	fmt.Println("next数组为：")
	for i,v := range next {
		fmt.Println(i,"-->",v)
	}

	for ;i<lb && j < ls;{
		if j == 0  || bigstr[i] == smallstr[j] {
			i ++
			j ++
		} else {
			j = next[j-1]
		}
	}

	if j == ls {
		return i-ls
	}
	return -1
}

func GetNextVal(smallstr string) []int { //计算前缀表
	//1.初始化
	var i int //前缀表下标
	var j int = 0 //前缀末尾
	l := len(smallstr)
	next := make([]int,l)
	next[0] = 0
	for i = 1;i < l; i ++ {
		//2.处理前后缀不同
		for j > 0 && smallstr[i] != smallstr[j] {
			j = next[j-1]
		}
		//3.处理前后缀相同
		if smallstr[i] == smallstr[j] {
			j ++
		}
		//4.next赋值
		next[i] =j
	}




	//k := -1
	//j := 0
	//l := len(smallstr)
	//next := make([]int,l)
	//next[j] = -1 	//next[i]=m 表示的意思是 p0p1...pm-1=pi-m...pi-2pi-1
	//
	//for ;j < l-1; {
	//	if k == -1 || smallstr[k] == smallstr[j] {
	//		j++
	//		k++
	//		next[j] = k
	//	}else {
	//		k = next[k]  //递归
	//	}
	//}
	return next
}
