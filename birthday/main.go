package main

import "fmt"

func main()  {
	//var day int = 365
	var day int = 4
	for j:=2;;j++ {
		if p:= BirthDay(j,day);p<1 {
			fmt.Println(j,"->",p)
		}else {
			break
		}
	}
}

func BirthDay(n int,day int) float64 {
	var k,m float64 = 1,1
	for i:=0;i<n;i++ {
		k *= float64(day-i)
		m *= float64(day)
	}
	return 1-k/m
}

// 23人，概率大于0.5
// 121人，概率大于1

