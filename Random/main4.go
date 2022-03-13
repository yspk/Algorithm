package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
)

func main()  {
	hashRate := flag.Int("hashRate", 2000, "Hash Rate")
	diff := flag.Float64("diff", 20000, "Difficulty")
	blockTime := flag.Int("blockTime", 1, "Block Time")
	flag.Parse()
	fmt.Println(*hashRate,*diff,*blockTime)
	target := 1.0 / *diff
	rand.Seed(time.Now().Unix())
	ticker := time.NewTicker(time.Duration(*blockTime) * time.Second)
	for {
		select {
		case <-ticker.C:
			var find bool
			var val float64
			for i:= 0;i<=*hashRate*(*blockTime);i++ {
				r := rand.Float64() // 求hash
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
