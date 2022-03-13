package main

import (
	"fmt"
	"github.com/simplechain-org/go-simplechain/common/hexutil"
)

func main()  {
	var data []byte
	for i:=0;i<1;i++ {
		data = append(data,0)
	}

	fmt.Println(hexutil.Encode(data))

}
