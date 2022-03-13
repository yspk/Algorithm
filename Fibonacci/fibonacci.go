package main

import (
	"fmt"
	"time"
)

var temp = make(map[int]uint64)

func main() {
	begin := time.Now().Unix()
	fmt.Println(fibonacci(1000000))
	fmt.Printf("cost time:%ds", time.Now().Unix()-begin)
}

func fibonacci(n int) uint64 {
	if k, ok := temp[n]; ok {
		return k
	} else if n == 0 {
		temp[0] = 0
		return temp[0]
	} else if n == 1 {
		temp[1] = 1
		return temp[1]
	} else {
		temp[n] = fibonacci(n-1) + fibonacci(n-2)
		return temp[n]
	}
}
