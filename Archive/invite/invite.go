package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/simplechain-org/go-simplechain"
	"github.com/simplechain-org/go-simplechain/accounts/abi"
	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/common/hexutil"
	"github.com/simplechain-org/go-simplechain/rpc"
	"io/ioutil"
	"log"
	"math/big"
)

var rawurlVar *string = flag.String("rawurl", "http://192.168.2.222:9555", "rpc url")

var abiPath *string = flag.String("abi", "./invite.json", "abi文件路径")

var contract *string = flag.String("contract", "0xc00d8b15363ca1f609a37b70e4a06d81ad84d85e", "合约地址")

var fromVar *string = flag.String("from", "0x0678c0f4ff27cdbe9f6f08dcf768f63a87e2f096", "接单人地址")

var gaslimitVar *uint64 = flag.Uint64("gaslimit", 18000000, "gas最大值")

var begin *int64 = flag.Int64("begin",2,"开始")


var count *int64 = flag.Int64("count",1,"计数")

type SendTxArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	Input    *hexutil.Bytes  `json:"input"`
}

func main() {
	flag.Parse()
	Match()
}

func Match() {

	data, err := ioutil.ReadFile(*abiPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	//账户地址
	from := common.HexToAddress(*fromVar)
	//合约地址
	//在子链上接单就要填写子链上的合约地址
	//在主链上接单就要填写主链上的合约地址
	to := common.HexToAddress(*contract)

	gas := hexutil.Uint64(*gaslimitVar)

	value := new(hexutil.Big)

	abi, err := abi.JSON(bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
		return
	}

	client, err := rpc.Dial(*rawurlVar)
	if err != nil {
		fmt.Println("dial", "err", err)
		return
	}
	var i int64
	var step int64 = 100
	for i = 0; i < *count; i ++ {
		start := *begin + i * step

		out, err := abi.Pack("testinviter", big.NewInt(start), big.NewInt(step))

		if err != nil {
			fmt.Println("abi.Pack err=", err)
			return
		}

		input := hexutil.Bytes(out)
		//var hex hexutil.Uint64
		//
		//err = client.CallContext(context.Background(), &hex, "eth_estimateGas", toCallArg(simplechain.CallMsg{
		//	From: from, //若无From字段，方法中的require(msg.sender)会报错
		//	To:   &to,
		//	Data: out,
		//}))
		//if err != nil {
		//	fmt.Println("eth_estimateGas", "err", err)
		//	return
		//}



		var result common.Hash
		args := &SendTxArgs{
			From:  from,
			To:    &to,
			Gas:   &gas,
			Value: value,
			Input: &input,
		}

		err = client.CallContext(context.Background(), &result, "eth_sendTransaction", args)

		if err != nil {
			fmt.Println("CallContext", "err", err,"last begin",start)
		}
		fmt.Println("eth_sendTransaction result=", result.Hex(),"count",i)

	}
	client.Close()

}

func toCallArg(msg simplechain.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}