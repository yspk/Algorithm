package main

import "fmt"

func main() {
	//sum := SumPeach(1)
	sum := 1
	for i :=  1; i< 10 ; i ++ {
		sum = sum *2 + 2
	}
	fmt.Println("第一天摘得桃子有：", sum)
}

func SumPeach(day int) int {
	if day == 10 {
		return 1
	}
	return 2*SumPeach(day+1) + 2
}
