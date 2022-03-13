package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	rand.Seed(1)
	for i:= 0;i<=1000000;i++ {
		fmt.Println(rand.Float64())
	}
}
