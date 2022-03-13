package main

import (
	"github.com/FactomProject/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"fmt"
)

func main(){
	// Generate a mnemonic for memorization or user-friendly seeds
	//entropy, _ := bip39.NewEntropy(256)
	//mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed("cable harbor icon crime holiday bullet box happy throw puzzle wolf stand", "")

	masterKey, _ := bip32.NewMasterKey(seed)
	childKey,_ := masterKey.NewChildKey(0)
	publicKey := childKey.PublicKey()

	// Display mnemonic and keys
	//fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey.Key)
}