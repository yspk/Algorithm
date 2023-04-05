package main

import (
	"fmt"
	"github.com/ethereumpow/go-ethereum/p2p/enode"
)

func main() {
	fmt.Println(enode.MustParse("enode://00ae0524916d62e387cca4217f66e06373e042cd3f9f853d7c66acc1b1a29ccd59c9c9582eda76f99e67a0692318f090328d24e43cd910cb00c7063326fe08ce@172.31.112.69:30311").ID().String())
}
