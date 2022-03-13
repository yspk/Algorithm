package main

import (
	//"encoding/hex"
	"fmt"
	//"github.com/simplechain-org/go-simplechain/crypto/sha3"
)

func main() {
	//transferFnSignature := []byte("getTakerTx(bytes32,uint256)")
	//hash := sha3.NewKeccak256()
	//hash.Write(transferFnSignature)
	//methodID := hash.Sum(nil)[:4]
	//fmt.Println(hex.EncodeToString(methodID))
	a := []int{8,16,8,16,8,8,8,8,256,8,64,8,8,16,16,256,8,16,8,8,64,8,8,8,16,16,8,16,8,16,8,1024,8,16,8,8,16,8,8,16,8,16,8,8,64,8,16,16,8,16,16,16,8,8,8,16,8,64,16,8,
		16,16,8,16,8,256,8,8,16,16,8,16,8,16,16,8,16,8,8,64,8,16,64,8,256,16,16,8,8,
		16,8,8,8,16,16,16,8,8,16,64,8,8,8,8,16,8,16,8,8,16,8,16,8,8,8,64,8,16,16,64,
		16,8,16,8,8,8,8,8,16,8,64,8,8,16,8,16,8,64,64,8,8,8,8,8,16,8,16,8,16,16,8,8,
		16,8,8,16,16,64,64,8,8,8,64,16,8,8,8,16,8,16,8,16,8,8,16,8,8,8,8,16,16,16,8,
		8,8,8,8,8,8,8,8,8,16,64,8,8,8,8}
	var total,count8,count16,count64,count256,count1024 int
	for _, i:= range a {
		total += i
		if i == 8 {
			count8 ++
		} else if i == 16 {
			count16 ++
		} else if i == 64 {
			count64 ++
		} else if i == 256 {
			count256 ++
		} else if i == 1024 {
			count1024 ++
		}
	}
	fmt.Println(len(a),total,count8,count16,count64,count256,count1024,count8+count16+count64+count256+count1024)

}
