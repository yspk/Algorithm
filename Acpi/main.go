package main

import "fmt"

func main()  {
	var a float64 = 91000
	var I float64 = 0.17 //年利率
	i := I/12
	var n int = 24
	var count float64=1
	for k := 0 ; k < n ; k++ {
		count *= 1+i
	}

	pre := (a * i* count)/(count -1)
	fmt.Println(pre)
}
