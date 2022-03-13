package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/simplechain-org/go-simplechain/accounts/abi"
	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/common/hexutil"
	"github.com/simplechain-org/go-simplechain/rpc"
	"io/ioutil"
	"log"
	"math/big"
)

var rawurlVar *string =flag.String("rawurl", "http://127.0.0.1:8555", "rpc url")

var abiPath *string =flag.String("abi", "./main_ctx_abi.json", "abi文件路径")

var contract *string =flag.String("contract", "0xbCd7d2069B81D93932dC8b57E135EDc7B5594996", "合约地址")

var value *uint64 = flag.Uint64("value", 1e+18, "转入合约的数量")

var chainId *uint64=flag.Uint64("chainId", 1, "目的链id")

var destValue *uint64=flag.Uint64("destValue", 1e+18, "兑换数量")

var fromVar *string=flag.String("from", "0x99baf6f31b8fae1715e0cc4f5d27bd464aa9d8ca", "发起人地址")

var gaslimitVar *uint64=flag.Uint64("gaslimit", 500000, "gas最大值")

type SendTxArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`
}

func Maker(client *rpc.Client) {

	data,err:=ioutil.ReadFile(*abiPath)

	if err!=nil{
		fmt.Println(err)
		return
	}

	from := common.HexToAddress(*fromVar)

	to := common.HexToAddress(*contract)

	gas := hexutil.Uint64(*gaslimitVar)

	value := hexutil.Big(*big.NewInt(0).SetUint64(*value))

	abi, err := abi.JSON(bytes.NewReader(data))

	if err != nil {
		log.Fatalln(err)
	}
	////想要随机，请设置随机种子
	//rand.Seed(time.Now().UnixNano())
	//
	////这个值必须改变，因为已经限定合约中，计算出来的txid不能相同
	//nonce:=big.NewInt(0).SetUint64(rand.Uint64())
	//nonce:=big.NewInt(0).SetUint64(9536605289005490782)

	id:=big.NewInt(0).SetUint64(*chainId)

	des:=big.NewInt(0).SetUint64(*destValue)

	out, err := abi.Pack("makerStart",id ,des)

	input := hexutil.Bytes(out)

	fmt.Println("input=",input,"id=",id,"des=",des)

	args := &SendTxArgs{
		From:  from,
		To:    &to,
		Gas:   &gas,
		Value: &value,
		Input: &input,
	}

	var result common.Hash

	err = client.CallContext(context.Background(), &result, "eth_sendTransaction", args)

	if err != nil {
		fmt.Println("CallContext", "err", err)
		return
	}

	fmt.Println("result=", result.Hex())
}


//跨链交易发起人
func main() {
	flag.Parse()
	client, err := rpc.Dial(*rawurlVar)
	if err != nil {
		fmt.Println("dial", "err", err)
		return
	}
	for i:=0; i< 5000; i++ {
		Maker(client)
	}
}


