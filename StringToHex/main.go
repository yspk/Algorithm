package main

import (
	"fmt"
	"github.com/fusion/go-fusion/common/hexutil"
)

func main() {
	s := `{"林木编号":"469007100024JE00012L032804100051","GPS":"109.091843,18.978084","地址":"海南省东方市东河镇","参数":{"育苗时间":"2010-05-31 19:19:20","栽种时间":"2016-07-01 19:19:35","品种":"降香黄檀","面积(m²)":20,"株数":"单株树"}`
	fmt.Println(hexutil.Encode([]byte(s)))
}

