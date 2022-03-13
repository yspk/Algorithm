package main

import "fmt"

func main() {
	var (
		a int         = 0
		b int64       = 0
		c interface{} = int(0)
		d interface{} = int64(0)
	)

	fmt.Println(c == 0)
	fmt.Println(c == a)
	fmt.Println(c == b)
	fmt.Println(d == b)
	fmt.Println(d == 0)
}
