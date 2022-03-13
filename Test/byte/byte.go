package main

import (
	"fmt"
	"github.com/fusion/go-fusion/common/hexutil"
)

func main()  {
	data,err := hexutil.Decode("0x6b24a08dde7f967f5d9e552e2ff87fc9cb8a3431069c579753598996ac2c955c33343735343138")

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		var nonZeroGas uint64 = 68
		gas := nz * nonZeroGas

		z := uint64(len(data)) - nz
		gas += z * 4
		fmt.Println(len(data),nz,z,gas)
	}
}
