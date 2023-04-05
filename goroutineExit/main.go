package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s) //无限循环，除非收到close命令才会退出
			}
		}()
		return completed
	}
	strings := make(chan string)
	result := doWork(strings)
	strings <- "Hello"
	strings <- "World"
	close(strings)
	<-result
	fmt.Println("Done.")
}
