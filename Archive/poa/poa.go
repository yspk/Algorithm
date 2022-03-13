package main

import (
	//"bytes"
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/simplechain-org/go-simplechain/ethclient"
	//"github.com/simplechain-org/go-simplechain/accounts/abi"
	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/core/types"
	"github.com/simplechain-org/go-simplechain/crypto"
	//"io/ioutil"
	//"log"
	"math/big"
	"time"
)

//var rawurlVar *string =flag.String("rawurl", "http://127.0.0.1:8501", "rpc url")

var contract *string =flag.String("contract", "0x738b546958b83feba48a758172a1d946944c5581", "合约地址")

var value *uint64 = flag.Uint64("value", 1e+18, "转入合约的数量")

var fromVar *string=flag.String("from", "0x738b546958b83feba48a758172a1d946944c5581", "发起人地址")

var gaslimitVar *uint64=flag.Uint64("gaslimit", 22000, "gas最大值")

var i uint64

func Maker(key *ecdsa.PrivateKey,to common.Address,out []byte,nonce,gasLimit uint64,amount,gasPrice *big.Int) *types.Transaction {
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, out)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1515)), key)
	if err != nil {
		fmt.Println("signedTx", "err", err)
		return nil
	}

	return signedTx
}


//跨链交易发起人
func main() {
	flag.Parse()

	//rpc.Client{}.
	//	BatchCall()

	client, err := ethclient.Dial("http://127.0.0.1:8501")
	if err != nil {
		fmt.Println("dial", "err", err)
		return
	}

	privateKey1, err := crypto.HexToECDSA("a4a5b53fe15132362e4dab8245df2329bca24edd0dbe4a1f436dd6b16fd34a5f")
	if err != nil {
		fmt.Println("privateKey", "err", err)
		return
	}

	//client2, err := ethclient.Dial("http://172.16.160.62:8502")
	//client2, err := ethclient.Dial("http://127.0.0.1:8501")
	//if err != nil {
	//	fmt.Println("dial", "err", err)
	//	return
	//}

	privateKey2, err := crypto.HexToECDSA("f6ff8db0e372b0d94468a7ff09ff57c6a112793aa79bf5221baa02979b9085ef")
	if err != nil {
		fmt.Println("privateKey", "err", err)
		return
	}
	//
	//client3, err := ethclient.Dial("http://172.16.160.64:8503")
	//client3, err := ethclient.Dial("http://127.0.0.1:8501")
	//if err != nil {
	//	fmt.Println("dial", "err", err)
	//	return
	//}

	privateKey3, err := crypto.HexToECDSA("ae3b6fb1662d143c3ec78030de4aa29661bc0fb70a6a2895a8daa23ddc2ae59a")
	if err != nil {
		fmt.Println("privateKey", "err", err)
		return
	}
	//
	//client4, err := ethclient.Dial("http://172.16.160.63:8504")
	//client4, err := ethclient.Dial("http://127.0.0.1:8501")
	//if err != nil {
	//	fmt.Println("dial", "err", err)
	//	return
	//}

	privateKey4, err := crypto.HexToECDSA("856234c591d083ef01da7cfdb620c48e59c94e1a453b6258beb50696d2445de0")
	if err != nil {
		fmt.Println("privateKey", "err", err)
		return
	}

	go Sender(client,privateKey1)
	go Sender(client,privateKey2)
	go Sender(client,privateKey3)
	go Sender(client,privateKey4)
	time.Sleep(time.Second*1000)
}

func Sender(client *ethclient.Client,key *ecdsa.PrivateKey)  {
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("publicKey", "err", ok)
		return
	}

	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		fmt.Println("nonce", "err", err)
		return
	}

	gas := *gaslimitVar

	amount :=  new(big.Int).SetUint64(*value)


	begin := time.Now()
	var txs []*types.Transaction
	for ; i< 10000000; i++ {
		hash := Int64ToBytes(i)
		tx := Maker(key,from,hash,nonce+i,gas,amount,big.NewInt(1e9))
		fmt.Println(i,tx.Hash().String(),"0x"+common.Bytes2Hex(hash))
		txs = append(txs,tx)
		if len(txs) >= 50 {
			err = client.SendTransactions(context.Background(), txs)
			if err != nil {
				fmt.Println("SendTransaction", "err", err)
				return
			}
			txs = txs[:0]
		}
	}
	fmt.Println("finish",time.Now().Sub(begin))
}


//func getHash() []byte {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//
//	a := r.Uint64()
//	b := r.Uint64()
//	c := r.Uint64()
//	d := r.Uint64()
//
//	str := fmt.Sprintf("%016x%016x%016x%016x", a, b, c, d)
//	data := common.Hex2Bytes(str)
//	return data
//}

func Int64ToBytes(i uint64) []byte {
	var buf = make([]byte, 64)
	PutUint64(buf, i)
	return buf
}

func PutUint64(b []byte, v uint64) {
	_ = b[63] // early bounds check to guarantee safety of writes below

	b[56] = byte(v >> 56)
	b[57] = byte(v >> 48)
	b[58] = byte(v >> 40)
	b[59] = byte(v >> 32)
	b[60] = byte(v >> 24)
	b[61] = byte(v >> 16)
	b[62] = byte(v >> 8)
	b[63] = byte(v)
}