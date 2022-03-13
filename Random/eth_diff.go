package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

var (
	maxint256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
)

func main() {
	fmt.Println(hexutil.EncodeBig(maxint256))
	var diff = big.NewInt(2238.54e12)
	target_from_diff := new(big.Int).Div(maxint256, diff)
	fmt.Println(hexutil.EncodeBig(target_from_diff))

	fz := diff
	fz = fz.Div(fz,big.NewInt(13)) //平均出块时间
	fmt.Println(fz)
}
