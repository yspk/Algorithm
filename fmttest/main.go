package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(fmt.Sprintf("%d", 67))
	fmt.Println(time.Now().UnixMilli())
	t := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-t.C:
			fmt.Println(time.Now().UnixMilli())
		}
	}
}
