package main

import (
	"github.com/yspk/frame/src/common/logger"
	"coding.net/baoquan2017/dataqin-backend/common/util"
	"github.com/yspk/Algorithm/email/code"
	"strconv"
	"time"
)

var QueryUrl = "https://www.zt.com/api/v1/allticker"

func main()  {
	type DataTicker struct {
		Buy string `json:"buy"`
		Change string `json:"change"`
		High string `json:"high"`
		Last string `json:"last"`
		Low string `json:"low"`
		Sell string `json:"sell"`
		Symbol string `json:"symbol"`
		Vol string `json:"vol"`
	}

	type Res struct {
		Data string `json:"data"`
		Ticker []*DataTicker `json:"ticker"`
	}

	reporter := time.NewTicker(time.Hour)
	//reporter := time.NewTicker(time.Second *30)
	tester := time.NewTicker(time.Minute * 5)

	for {
		select {
		case <-reporter.C:
			var response Res
			var err error
			if err = util.Get(QueryUrl, &response, nil); err != nil {
				logger.Error(err)
			}
			var sipc *DataTicker
			if response.Ticker != nil {
				for _,v := range response.Ticker {
					if v.Symbol == "SIPC_CNT" {
						sipc = v
						var buy,sell float64
						buy, err = strconv.ParseFloat(sipc.Buy, 64)
						if err != nil {
							logger.Error(err)
						}
						sell, err = strconv.ParseFloat(sipc.Sell, 64)
						if err != nil {
							logger.Error(err)
						}
						code.SendCode("yus501501@163.com",0,sell,buy)
					}
				}
			}

		case <-tester.C:
			var response Res
			var err error
			if err = util.Get(QueryUrl, &response, nil); err != nil {
				logger.Error(err)
			}
			var sipc *DataTicker
			if response.Ticker != nil {
				for _,v := range response.Ticker {
					if v.Symbol == "SIPC_CNT" {
						sipc = v
						var buy,sell float64
						buy, err = strconv.ParseFloat(sipc.Buy, 64)
						if err != nil {
							logger.Error(err)
						}
						sell, err = strconv.ParseFloat(sipc.Sell, 64)
						if err != nil {
							logger.Error(err)
						}
						if buy >= 1.2 {
							code.SendCode("yus501501@163.com",1,sell,buy)
						} else if sell <= 0.95 {
							code.SendCode("yus501501@163.com",2,sell,buy)
						}
					}
				}
			}
		}
	}
}
