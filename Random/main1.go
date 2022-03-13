package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	target := 1.0 / 20000.0
	rand.Seed(time.Now().Unix())
	for i:= 0;i<=100000;i++ {
		r := rand.Float64()
		if r < target {
			fmt.Println(r)
		}
	}
}
