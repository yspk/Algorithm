package main

import "fmt"

func main()  {
	prove(5)
}

func prove(n int)  {
	var k int
	fmt.Printf("现在开始证明P(%d)成立。\n",n)
	fmt.Printf("根据步骤1得出P(%d)成立。\n",k)
	for ;k<n;k++{
		fmt.Printf("根据步骤2可以说'若P(%d)成立，则P(%d)也成立。'\n",k,k+1)
		fmt.Printf("因此可以说'P(%d)是成立的'。\n",k+1)
	}
	fmt.Printf("证明结束。\n")
}