package main

import "fmt"

func main() {
	//var a float64 = 91000
	//var I float64 = 0.165 //年利率
	var a float64 = 7116
	var I float64 = 0.12 //年利率
	i := I / 12
	var n int = 36 //分期
	var count float64 = 1
	for k := 0; k < n; k++ {
		count *= 1 + i
	}

	pre := (a * i * count) / (count - 1)
	fmt.Println(pre)
	fmt.Println(pre*float64(n) - a)
}
