package main

import (
	"fmt"
	"time"
)

var temp = make(map[string]string)

func main()  {
	begin := time.Now().Unix()
	fmt.Println(hanoi(6,"A","B","C"))  // n>=18 内存会出问题
	fmt.Printf("cost time:%ds", time.Now().Unix()-begin)
}

func hanoi(n int,x,y,z string) string {
	ke := tempkey(n,x,y,z)
	if k,ok:=temp[ke];ok{
		//fmt.Println(k)
		return k
	} else if(n==0){
		//TODO
		return ""
	}else {
		hanoi(n-1,x,z,y)
		temp[ke] = hanoi(n-1,x,z,y) + fmt.Sprintf("%s->%s,",x,y) + hanoi(n-1,z,y,x)
		return temp[ke]
	}
}

func tempkey(n int,x,y,z string) string {
	return fmt.Sprintf("%d_%s_%s_%s",n,x,y,z)
}

