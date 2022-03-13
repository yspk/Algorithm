package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	target := 1.0 / 20000.0
	rand.Seed(time.Now().Unix())
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			var find bool
			var val float64
			for i:= 0;i<=20000;i++ {
				r := rand.Float64()
				if r < target {
					find = true
					val = r
					break
				}
			}

			if find {
				fmt.Println("find",val)
			} else {
				fmt.Println("not find")
			}
		}
	}
}
