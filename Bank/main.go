package main

import (
	"math"
	"github.com/yspk/frame/src/common/logger"
)

func main()  {
	m := 13000.00 + 5272.00
	logger.Infof("15号发工资，每月%.2f",m)
	logger.Infof("中信银行贷30万，20号还款")
	k := Acpi(300000,0.047,36)
	logger.Infof("每月20号还款:%.2f",k)
	//50万存投哪网，等额本息，19号回款
	//l := Acpi(500000,0.118)
	//loggepr.Infof("每月19号回款:%.2f",l)
	//logger.Infof("利息收益:%.2f",l-k)
	logger.Infof("用60万买房，贷款140万，21号换房贷")
	p := Acpi(1400000,0.049,360)
	logger.Infof("房贷月供%.2f",p)
	//logger.Infof("房子抵押贷款60万")
	//s := Acpi(600000,0.047,36)
	//logger.Infof("房子抵押贷款每月22号还款:%.2f",s)
	//logger.Infof("60万存投哪网，等额本息，19号回款")
	//l := Acpi(600000,0.118,36)
	//logger.Infof("每月19号回款:%.2f",l)
	l,s := float64(0),float64(0)
	logger.Infof("每月结余:%.2f",l+m-k-p-s)

}

func Acpi(a float64,I float64,n int) float64 {
	//var a float64 = 91000
	//var I float64 = 0.165 //年利率
	i := I/12
	//var n int = 24
	var count float64=1
	for k := 0 ; k < n ; k++ {
		count *= 1+i
	}

	pre := (a * i* count)/(count -1)
	return FloatCount(pre,2)
}

func FloatCount(s float64,n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((s+0.5/pow10_n)*pow10_n) / pow10_n
}
