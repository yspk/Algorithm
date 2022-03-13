package main

import (
	"crypto/ecdsa"
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/simplechain-org/go-simplechain/common"
	"github.com/simplechain-org/go-simplechain/crypto"
	"github.com/simplechain-org/go-simplechain/crypto/sha3"
	"math"
)

var begin *uint64=flag.Uint64("begin", 1, "开始值")

var db *sql.DB
func init() {
	db, _ = sql.Open("mysql", "root:toor@tcp(192.168.2.222:3306)/simple_explorer?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func main()  {
	flag.Parse()
	var i uint64
	for i=*begin;i<= math.MaxUint64;i++ {
		if key,err := StringToPrivateKey(Int64ToBytes(i));err == nil {
			Int64ToBytes(i)
			address := Hex(crypto.PubkeyToAddress(key.PublicKey))
			var d float64
			err := db.QueryRow("select balance from balance where address = ? limit 1",address).Scan(&d)
			if err != sql.ErrNoRows {
				fmt.Println(address,i,d,"find address key!!")
			} else {
				fmt.Println(address,i)
			}
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("finish")
}

func StringToPrivateKey(privateKeyByte []byte) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func Int64ToBytes(i uint64) []byte {
	var buf = make([]byte, 32)
	PutUint64(buf, i)
	return buf
}

func Hex(address common.Address) string {
	unchecksummed := hex.EncodeToString(address[:])
	sha := sha3.NewKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		//if result[i] > '9' && hashByte > 7 {
		//	result[i] -= 32
		//}
	}
	return "0x" + string(result)
}

func PutUint64(b []byte, v uint64) {
	_ = b[31] // early bounds check to guarantee safety of writes below
	b[24] = byte(v >> 56)
	b[25] = byte(v >> 48)
	b[26] = byte(v >> 40)
	b[27] = byte(v >> 32)
	b[28] = byte(v >> 24)
	b[29] = byte(v >> 16)
	b[30] = byte(v >> 8)
	b[31] = byte(v)
}