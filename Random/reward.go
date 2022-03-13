package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func main()  {
	miner,_ :=  hexutil.DecodeBig("0x203a3c86c6d26b380000")
	//fund,_ :=  hexutil.DecodeBig("0x1b2391491337ee80000")
	//								0x1b2391491337ee80000
	miner.Div(miner,big.NewInt(19))

	fmt.Println(hexutil.EncodeBig(miner))
}
