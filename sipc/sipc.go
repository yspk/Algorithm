package main

import (
	"coding.net/baoquan2017/wallet-go/lib/web3"
	"coding.net/baoquan2017/wallet-go/lib/web3/providers"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:toor@tcp(192.168.2.222:3306)/simple_explorer?charset=utf8mb4&parseTime=true&loc=Local")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func main() {
	pool()
}

//func startHttpServer() {
//	http.HandleFunc("/pool", pool)
//	err := http.ListenAndServe(":9090", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

func pool() {
	rows, err := db.Query("select distinct miner as address,concat(year(from_unixtime(`timestamp`)),'-',LPAD(month(from_unixtime(`timestamp`)),2,0),'-',LPAD(day(from_unixtime(`timestamp`)),2,0)) as date from blocks  union select distinct t.`from` as address,concat(year(from_unixtime(b.`timestamp`)),'-',LPAD(month(from_unixtime(b.`timestamp`)),2,0),'-',LPAD(day(from_unixtime(b.`timestamp`)),2,0)) as date from transactions as t left join blocks b on t.blockNumber = b.number union select distinct t.`to` as address,concat(year(from_unixtime(b.`timestamp`)),'-',LPAD(month(from_unixtime(b.`timestamp`)),2,0),'-',LPAD(day(from_unixtime(b.`timestamp`)),2,0)) as date from transactions  as t left join blocks b on t.blockNumber = b.number")
	defer rows.Close()
	checkErr(err)

	httpP, _ := providers.NewHttpProvider("http://192.168.2.222:8545", 100)
	//httpP, _ := providers.NewHttpProvider("http://127.0.0.1:8545", 100)
	w3 := web3.NewWeb3(httpP)
	t := time.Now().Format("2006-01-02 15:04:05")
	for rows.Next() {
		var address string
		var date string
		rows.Scan(&address,&date)
		if count, err := w3.Eth.GetBalance(address); err != nil {
			fmt.Println(err.Error())
		} else {
			var d string
			err := db.QueryRow("select date from balance where address = ? and time = ?",address,t).Scan(&d)
			//err := db.QueryRow("select time from balance where address = ? limit 1",address).Scan(&d)
			if err == sql.ErrNoRows {
				ro,err := db.Query("INSERT INTO balance (address,date,balance,time) value (?,?,?,?)",address,date,count,t)
				checkErr(err)
				ro.Close()
				fmt.Println(address,count,date)
			} else {
				//fmt.Println(d,t)
				checkErr(err)
				fmt.Println("already record:",address)
				continue
			}

		}
	}


	fmt.Println("finish")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
