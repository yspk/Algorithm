package main

import (
	"context"
	"fmt"
	"github.com/simplechain-org/go-simplechain/accounts"
	"github.com/simplechain-org/go-simplechain/accounts/keystore"
	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/core/types"
	"github.com/simplechain-org/go-simplechain/ethclient"
	"github.com/simplechain-org/go-simplechain/rlp"
	"github.com/simplechain-org/go-simplechain/rpc"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	//Host = "ws://192.168.3.124:8546"
	//Host = "http://192.168.3.124:8545"
	Host = "http://localhost:8501"
	//Host = "ws://localhost:8546"
)

var (
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	gasPrice  = big.NewInt(1000000000)
	chainId   = big.NewInt(1515)
	value     = big.NewInt(1)

	keystore1  = `{"address":"9c7461b44ae3023e9f16ad58e5f17c24685ec2a3","crypto":{"cipher":"aes-128-ctr","ciphertext":"b47e278b16ee99963c89dad2ee725f6cf57edfa7619ee6abdeb9b6c1ce1e2c0d","cipherparams":{"iv":"b4dd1dc32936ab5530082311bca77a60"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"dbe135df81708c45a5f96dfcba6d9fbd3e19463de5f24b5afd5551c5593f1661"},"mac":"5467225a040209e85b78b1c6350e3cffc1be9cd5aab87e967a25bd968786809c"},"id":"4146b8d2-e261-47d1-984b-12cdef27a410","version":3}`
	keystore2  = `{"address":"7317cff0248f25a822f4348c79e1645a5e8a3929","crypto":{"cipher":"aes-128-ctr","ciphertext":"57bccaaae99b9fe6ee9e15f855b89a168f236664fbccb42cdb3148c4437a7119","cipherparams":{"iv":"a92803374fedddd65fe885b19e27bf19"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"0d553a6a2de25cc786c803e518235676a64b1aad078d3aa68342b27dc0596c2e"},"mac":"bce8c73b5aba11650c0ee437517adf18c85a84b8545911a23ab62a07a482195e"},"id":"357341a6-8db3-41c5-9248-3cc7fc996937","version":3}`
	keystore3  = `{"address":"735fefe632acecd304a4377daa91a30fcb0a2af8","crypto":{"cipher":"aes-128-ctr","ciphertext":"efa2f150d69946824d89e0f258ddf7ebb8967587ece1c17cc9c21aa0f95df965","cipherparams":{"iv":"f89cdd748f028262273195aeb1a04903"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d6450aa93d6cc9db66b1395f9f143efc50b29c57859de9435ee6e22946f91d2b"},"mac":"d027f81e572fc6e8ee86f01b68755fc795892e5699b6e5052add47dd72cdc058"},"id":"124ab27f-3be6-467e-ad3b-74b06c148824","version":3}`
	keystore4  = `{"address":"4728e7ab6b1bdd5983cb357c7d8ac79c6931ec09","crypto":{"cipher":"aes-128-ctr","ciphertext":"7322509a86815d9980ad90466d04751822e27fba98422f0c54105dac0a2a0120","cipherparams":{"iv":"61ed3bb48b31306c2dee0e102d8c2a0a"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5cd4f4278578c89e20af3d7fd4f33874ed6ae54b6cc145cb1efd58ef124b7a7b"},"mac":"62f4f12589f331a7f8e2a3fe89a2821c5531717cfc87157704acac3db7d113aa"},"id":"52ff2c6d-6755-4698-9b02-ca4a40ea443a","version":3}`
	keystore5  = `{"address":"41e257a6980fdb590cc0a49c1de8b2ff7845241b","crypto":{"cipher":"aes-128-ctr","ciphertext":"70c3a64268ff8606ad9cbe456f1de21fa82a311e4ed24253b076f65ee258e012","cipherparams":{"iv":"a69b2c8d6a99c84dbad7031247aa887a"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"db51badce5317916458b3da54e1408957babc28582ca674ee3ba4b598199abb2"},"mac":"99ff055d551b0ba9720c43e80fc6e71d872f1d8ac4ff86f108fc1b8e500765f4"},"id":"7696a96e-034b-4d50-9b40-1ecc7605202c","version":3}`
	keystore6  = `{"address":"0408fee2ef1b15da5b8057546113fea6e299b613","crypto":{"cipher":"aes-128-ctr","ciphertext":"08d795b3539b8e866a90103c8c611ce6a2201633ac42f96b412023dcbff2051f","cipherparams":{"iv":"d75b4ba28a1190fe74256a82c684cc6e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"07798b7ea0110be2440f2fd354e83d51bead9570f0372d964f46114b3e24178c"},"mac":"ef15165642e16e7a06afed282ee51da228d2bf003f6e9e13f1fbf835e6f95760"},"id":"bfa27bf5-2a3a-4c7c-a6fe-cf3a9e8d4d46","version":3}`
	keystore7  = `{"address":"bcde2c0fcfc780e20679f20df71f54e0e9919d24","crypto":{"cipher":"aes-128-ctr","ciphertext":"3d6595c6d6cf7e2441351d1cd1ec0a4265543e1d9fbbd86b321b37e868c27059","cipherparams":{"iv":"ff97e2b443cc10f59d22ecb7f316bf2f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f403d3abbdef5ae376ce9c733016bc27b3256f6b4c9a1353e8cc7de75e1aebdd"},"mac":"30583a1076d68af0367b66a1ecf6bea5db82afc3d2ab62b84c156394755e617e"},"id":"7ceb9b06-16ae-4212-bc7a-00da2d3a0a7f","version":3}`
	keystore8  = `{"address":"afe3bc9c5b8f989bc5c6ddeb237f5917e765aa9e","crypto":{"cipher":"aes-128-ctr","ciphertext":"db640612b70aa10bee71032033702ea27f60ccb287b35c04b737e0aeb969bc6e","cipherparams":{"iv":"a39bcc6f8a47c8963d11812bf6d4d9ec"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"2e396aa5cfeda0e01399ccc722c6a8194f920b8080102883f8e5532537ba9916"},"mac":"3455d481cd6df886cee0bafcad16f7a4ca2bc931686e9328e14873e62dc46c0c"},"id":"a9d120a6-eea3-4a0d-bd0c-91ec617e611c","version":3}`
	keystore9  = `{"address":"2f2a38100163289776adb6c1e0075aaa64cc743e","crypto":{"cipher":"aes-128-ctr","ciphertext":"939a94b5a2946970937e87409bbb7b160ebfad98d877ce0490c735786b513ea2","cipherparams":{"iv":"447b32169196bec277678724eaed26f8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"064abbaabced9533575d73500d12d3736817e3a5bb16c294d176991f1e7c56d4"},"mac":"4ca7d991565abf7097e5afd12ed2db9be9de7d55e2e44e6979b506b9eb4f71f5"},"id":"540d2a2a-7cd0-47ef-9ecf-0a616f58b749","version":3}`
	keystore10 = `{"address":"4cd6fc92d6e564721074c0a1487c03a3879f31d3","crypto":{"cipher":"aes-128-ctr","ciphertext":"19f7ca5b74c5c9b2fe11fbb90f56be6e5440442ffaf8041083d926ae27fa4bc0","cipherparams":{"iv":"14f5a048823adaa090ff52696a456c91"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"05e7a445de7db05f2ac860655b261a0fec35fb5db0953e911aa427a194374d2f"},"mac":"5137e471cff26d7d02ebb7aa489a19acbca7acb30c7bbdf663a93b25e68c6591"},"id":"15fdd498-f44a-4e9e-a62a-c5ae13e4e5e5","version":3}`
	keystore11 = `{"address":"738b546958b83feba48a758172a1d946944c5581","crypto":{"cipher":"aes-128-ctr","ciphertext":"a4f438301d37143f4c866ead06f525d7ba5b12741a4a6ecebffd9070b9dee6ec","cipherparams":{"iv":"01626cc785cd5fe3ea8be7db6960f470"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"675b86987a2df842d3f5731ade637ffdb0ca6d63e619abe1f3434be4f13bb1c4"},"mac":"8db0e579d14dd634b8e2d39a16378e867dff62ce52c346702dc99c538f305740"},"id":"bbfe9cb9-66bb-49a6-b62a-7468be153ebb","version":3}`
	keystore12 = `{"address":"db4ed4b17db0721a39db115e07bc12007e793a19","crypto":{"cipher":"aes-128-ctr","ciphertext":"5dcfdf76d9284e85f8972a850f3845e17134b4bf4b1572cc3bb4d5699b590fb1","cipherparams":{"iv":"66c25887f2c077c889af5c56e483b398"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"dce0f92dcbeb515019c4ec5d072c3be728c992fd88c5c0f20b521c2019270375"},"mac":"89c0275ffa31baec788eeb1dc1c74e5cb31560d96c155c405d0e26919da98e30"},"id":"a81b648a-92de-41b2-889c-f4a73f869703","version":3}`
	keystore13 = `{"address":"34145ff275ca97e5e1f4f6208f2458383c13b4df","crypto":{"cipher":"aes-128-ctr","ciphertext":"7cf9f054711a12b9bd38a86b6a9681f2a50b15dcc5c5d23229b528c5cb6dc5fc","cipherparams":{"iv":"81d96530d16bd5f385adbb04b4b58120"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"dae74016488fa63a1d88703340eeb33f421ca743ab220396af28e13738091987"},"mac":"cc744e6eda7354a40f794e5d70549c86a0ac31aabcfcef8eb75a0f84347faa55"},"id":"a7ad0a54-e283-497a-ae8b-1bf325889059","version":3}`
	keystore14 = `{"address":"5dd5dee094f5cc8e7d3cd78e86b66584a22957da","crypto":{"cipher":"aes-128-ctr","ciphertext":"fb3614ac89c43cb3b795dbe3b25dd6b2aaab1fbc36962f4c17dcb50f7a546da3","cipherparams":{"iv":"1bae8f9edade2d3461b188c4b0d4cd20"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"93bd63ee6a33ab0a589743c56e3eaf84cd0920ccb84d882e3d3998f8a835f95f"},"mac":"213617edfcb61b7a507c2d0a7d1827db4a9f5f19f22c1a37c776642ea594f83c"},"id":"1814af88-af43-4403-bb48-209b37bb81d1","version":3}`


	to         = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ks         *keystore.KeyStore
	workDir, _ = ioutil.TempDir("", "keystore")
	acc        accounts.Account

	password  = "fuckyou2012"
	sendCount chan string
)

func main() {
	Init()
	wg := &sync.WaitGroup{}

	sendCount = make(chan string, 10000)

	a1, n1 := getKeystore(keystore1, password)
	//a2, n2 := getKeystore(keystore2, password)
	//a3, n3 := getKeystore(keystore3, password)
	//a11,n11 := getKeystore(keystore11,"123456")
	//a12,n12 := getKeystore(keystore12,"123456")
	//a13,n13 := getKeystore(keystore13,"123456")
	//a14,n14 := getKeystore(keystore14,"123456")

	start := time.Now()
	go func() {
		i := 0
		ticker := time.NewTicker(30 * time.Minute)
		for {
			select {
			case l := <-sendCount:
				log.Printf("%010d,%s", i, l)
			case <-ticker.C:
				log.Println("time after", time.Now().Sub(start).String())
				os.Exit(0)
			}
			i++
		}
	}()
	wg.Add(1)
	length := 1
	go test(a1, n1, length, wg)
	//go test(a2, n2, length, wg)
	//go test(a3, n3, length, wg)
	//go test(a11, n11, length, wg)
	//go test(a12, n12, length, wg)
	//go test(a13, n13, length, wg)
	//go test(a14, n14, length, wg)

	//go test(getKeystore(keystore4, "fuckyou2012"))
	//go test(getKeystore(keystore5, "fuckyou2012"))
	//go test(getKeystore(keystore6, "fuckyou2012"))
	//go test(getKeystore(keystore7, "fuckyou2012"))
	//go test(getKeystore(keystore8, "fuckyou2012"))
	//go test(getKeystore(keystore9, "fuckyou2012"))
	//go test(getKeystore(keystore10, "fuckyou2012"))
	wg.Wait()
	log.Println("time after", time.Now().Sub(start).String())
}

func Init() {
	var err error
	rpcClient, err = rpc.Dial(Host)
	if err != nil {
		log.Fatal(err)
	}

	ethClient, err = ethclient.Dial(Host)
	if err != nil {
		log.Fatal(err)
	}

	ks = keystore.NewKeyStore(workDir, keystore.StandardScryptN, keystore.StandardScryptP)
}

func test(account accounts.Account, nonce uint64, length int, wg *sync.WaitGroup) {
	//batches := make([]rpc.BatchElem, 0, 1000)
	//inputs := make([][]byte, 0, 1000)
	i := 0
	var txs []*types.Transaction
	for {
		if i >= length {
			break
		}
		input := getHash()
		tx := types.NewTransaction(nonce, to, value, uint64(50000), gasPrice, input)

		signedTx, err := ks.SignTx(account, tx, chainId)
		if err != nil {
			log.Fatal(err)
		}

		//data, err := rlp.EncodeToBytes(signedTx)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//var result common.Hash
		//elem := rpc.BatchElem{
		//	Method: "eth_sendRawTransaction", Args: []interface{}{common.ToHex(data)}, Result: &result}
		//
		//batches = append(batches, elem)
		//inputs = append(inputs, input)
		//if len(batches) > 300 {
		//	err = rpcClient.BatchCall(batches)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	for k, v := range batches {
		//		sendCount <- fmt.Sprintf("i%x,t:%s\n", inputs[k], v.Result.(*common.Hash).String())
		//	}
		//	batches = make([]rpc.BatchElem, 0, 1000)
		//	inputs = make([][]byte, 0, 1000)
		//}
		txs = append(txs,signedTx)
		sendCount <- fmt.Sprintf("i:%x,t:%s\n", input, signedTx.Hash().String())
		if len(txs) >= 1 {
			err = ethClient.SendTransactions(context.Background(), txs)
			if err != nil {
				fmt.Println("SendTransaction", "err", err)
				return
			}
			txs = txs[:0]
		}


		nonce++
		i++
	}

	wg.Done()
}

func sendTx() {
	tx := types.NewTransaction(0, to, value, uint64(50000), gasPrice, getHash())

	signedTx, err := ks.SignTx(acc, tx, chainId)
	if err != nil {
		log.Fatal(err)
	}

	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
	}

	var result common.Hash
	err = rpcClient.Call(&result, "eth_sendRawTransaction", common.ToHex(data))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ok txHash", result.String())
	receipt, err := ethClient.TransactionReceipt(context.Background(), result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("receipt", receipt)
}

func getKeystore(keystore, password string) (accounts.Account, uint64) {
	account, err := ks.Import([]byte(keystore), password, password)
	if err != nil {
		log.Fatal("Failed to import faucet signer account", "err", err)
	}

	err = ks.Unlock(account, password)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), account.Address)
	if err != nil {
		log.Fatal(err)
	}

	return account, nonce
}

func generateAddress() {
	pass := "fuckyou2012"
	acc, err := ks.NewAccount(pass)
	if err != nil {
		log.Println(err)
	}

	if err := ks.Unlock(acc, pass); err != nil {
		log.Fatal(err)
	}
	log.Println(acc.Address.String())

}

func getHash() []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := r.Uint64()
	//b := r.Uint64()
	//c := r.Uint64()
	//d := r.Uint64()

	str := fmt.Sprintf("%016x%016x%016x%016x", a, a, a, a)
	data := common.Hex2Bytes(str)
	return data
}
