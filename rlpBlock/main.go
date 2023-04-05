package main

import (
	"fmt"
	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	var block evmcore.EvmBlock
	v, err := hexutil.Decode("0xf8b1f8ae80a00000000000000000c20dbfb2ec18ae20037c716f3ba2d9e1da768a9deca17cb4a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421881652e52d61b17fff94000000000000000000000000000000000000000088ffffffffffffffff8080c0")
	err = rlp.DecodeBytes(v, &block)
	if err != nil {
		fmt.Println("decode block", "err", err)
		return
	}
	//if msg.ChainID == 56 {
	fmt.Println("receive block", "number", block.NumberU64(), "hash", block.Hash, "blockTime", block.Time, "block", block)
	//}
}
