package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/common/hexutil"
	"github.com/simplechain-org/go-simplechain/rpc"
	"math/big"
	"time"
)

//var rawurlVar *string =flag.String("rawurl", "http://127.0.0.1:8501", "rpc url")

//var contract *string =flag.String("contract", "0x738b546958b83feba48a758172a1d946944c5581", "合约地址")

var value *uint64 = flag.Uint64("value", 1e+18, "转入合约的数量")

//var fromVar *string=flag.String("from", "0x738b546958b83feba48a758172a1d946944c5581", "发起人地址")

var gaslimitVar *uint64=flag.Uint64("gaslimit", 22000, "gas最大值")

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

func Maker(client *rpc.Client,reqs []rpc.BatchElem) {
	if err := client.BatchCallContext(context.Background(), reqs); err != nil {
		fmt.Println("BatchCallContext","err",err)
		return
	}
	for i := range reqs {
		if reqs[i].Error != nil {
			fmt.Println(i,reqs[i].Error)
			return
		}
		if reqs[i].Result == nil {
			fmt.Println("result",nil)
			return
		}
	}

	//fmt.Println("result=", result.Hex())
}


//跨链交易发起人
func main() {
	flag.Parse()
	client1, err := rpc.Dial("http://127.0.0.1:8501")
	if err != nil {
		fmt.Println("dial", "err", err)
		return
	}

	client2, err := rpc.Dial("http://127.0.0.1:8501")
	if err != nil {
		fmt.Println("dial", "err", err)
		return
	}

	client3, err := rpc.Dial("http://127.0.0.1:8501")
	if err != nil {
		fmt.Println("dial", "err", err)
		return
	}

	//client4, err := rpc.Dial("http://172.16.160.63:8504")
	//if err != nil {
	//	fmt.Println("dial", "err", err)
	//	return
	//}

	go Sender(client1,"0x738b546958b83feba48a758172a1d946944c5581")
	go Sender(client2,"0x34145ff275ca97e5e1f4f6208f2458383c13b4df")
	go Sender(client3,"0xdb4ed4b17db0721a39db115e07bc12007e793a19")
	//go Sender(client4,"0x5dd5dee094f5cc8e7d3cd78e86b66584a22957da")
	time.Sleep(time.Second*1000)

}

func Int64ToBytes(i,j uint64) []byte {
	var buf = make([]byte, 64)
	PutUint64(buf, i,j)
	return buf
}

func PutUint64(b []byte, v,j uint64) {
	_ = b[63] // early bounds check to guarantee safety of writes below
	b[48] = byte(j >> 56)
	b[49] = byte(j >> 48)
	b[50] = byte(j >> 40)
	b[51] = byte(j >> 32)
	b[52] = byte(j >> 24)
	b[53] = byte(j >> 16)
	b[54] = byte(j >> 8)
	b[55] = byte(j)

	b[56] = byte(v >> 56)
	b[57] = byte(v >> 48)
	b[58] = byte(v >> 40)
	b[59] = byte(v >> 32)
	b[60] = byte(v >> 24)
	b[61] = byte(v >> 16)
	b[62] = byte(v >> 8)
	b[63] = byte(v)
}

func Sender(client *rpc.Client,address string)  {
	from := common.HexToAddress(address)

	to := common.HexToAddress(address)

	gas := hexutil.Uint64(*gaslimitVar)

	value := hexutil.Big(*big.NewInt(0).SetUint64(*value))



	var i uint64
	begin := time.Now()
	for i=0; i< 500; i++ {
		var j uint64
		results := make([]common.Hash,200)
		reqs := make([]rpc.BatchElem,200)
		for j=0; j < 200; j++ {
			input := hexutil.Bytes(Int64ToBytes(i,j))

			args := &SendTxArgs{
				From:  from,
				To:    &to,
				Gas:   &gas,
				Value: &value,
				Input: &input,
			}
			reqs[j] = rpc.BatchElem{
				Method: "eth_sendTransaction",
				Args:   []interface{}{args},
				Result: &results[j],
			}
		}
		Maker(client,reqs)
	}
	fmt.Println("finish",time.Now().Sub(begin))
}