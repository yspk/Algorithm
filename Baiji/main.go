package main

import "fmt"

func main() {
	var k, x, y, z int
	for k = 1; k <= 3; k++ {
		x, y, z = 4*k, 25-7*k, 75+3*k
		fmt.Printf("公鸡：%d只，母鸡：%d只，小鸡：%d只。\n", x, y, z)
	}
}
