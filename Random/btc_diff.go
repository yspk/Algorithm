package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

var (
	maxint256_32 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256-32), big.NewInt(0))
	maxUint256_32 = new(big.Int).Sub(maxint256_32,big.NewInt(1))
)

func main() {
	fmt.Println(hexutil.EncodeBig(maxUint256_32))
	var bits,_ = hexutil.DecodeBig("0x371ef4") //底
	l,_ := hexutil.DecodeBig("0x17") //次方数
	q:= 8*(l.Sub(l,big.NewInt(3)).Int64())
	target_from_bits := new(big.Int).Lsh(bits,uint(q))
	fmt.Println(hexutil.EncodeBig(target_from_bits))
	//diff_from_bits := new(big.Int).Div(maxUint256_32,target_from_bits)

	//var diff = diff_from_bits
	var diff = big.NewInt(5106422924659)
	target_from_diff := new(big.Int).Div(maxUint256_32, diff)
	fmt.Println(hexutil.EncodeBig(target_from_diff))

	fz := diff.Mul(big.NewInt(1<<32),diff)
	fz = fz.Div(fz,big.NewInt(600)) //平均出块时间
	fmt.Println(fz)
}
