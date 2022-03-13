package main

import (
	"time"
	"fmt"
)

func main(){

	queue := make(chan int)
	qu := queue
	fmt.Println(qu,queue)
	go func(chan int) {
		fmt.Println("程序等待")
		time.Sleep(time.Second*1)

		fmt.Println("结束命令")
		//close(queue)
		queue <- 0
	}(queue)

	<- qu
	fmt.Println("结束程序")
}
