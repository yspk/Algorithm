package main

import (
	"encoding/hex"
	"fmt"
	"github.com/W1llyu/ourjson"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//jsonStr := `{
	//	"call_module": "区块链证书",
	//	"call_module_function": "notary",
	//	"extrinsic_hash": "0x53a0a660e200ef32549a3f66093a7ee49a759b8cd71efa4b61ff26114a1b5e48",
	//	"block_time":" 2021-12-29 14:23:39",
	//	"info": {
	//		"token_id": "93911476072456014925824",
	//		"token_class_id":"",
	//		"owner": "",
	//		"content_hash": "0xb47e3cd837ddf8e4c57f05d70ab865de6e193baa",
	//		"content": {
	//			"林木编号": "469007009009JE00001L00019003",
	//			"GPS": "109.024986,19.013589",
	//			"地址": "海南省白沙黎族自治县细水乡白石岭",
	//			"育苗时间": "2015/12/07",
	//			"栽种时间": "2016/12/07",
	//			"参数": {
	//				"测量时间": "2016/12/07",
	//				"品种": "降香黄檀",
	//				"胸径": "2 cm",
	//				"高度": "100 cm",
	//				"面积": "16 平方",
	//				"株数": "1 株"
	//			}
	//		}
	//	}
	//}
    //`
	jsonStr := "7b22e69e97e69ca8e7bc96e58fb7223a20223436393030373030393030394a4530303030314c3030303139303035222c22475053223a20223130392e3032343938362c31392e303133353839222c22e59cb0e59d80223a2022e6b5b7e58d97e79c81e799bde6b299e9bb8ee6978fe887aae6b2bbe58ebfe7bb86e6b0b4e4b9a1e799bde79fb3e5b2ad222c22e882b2e88b97e697b6e997b4223a2022323031352f31322f3037222c22e6a0bde7a78de697b6e997b4223a2022323031362f31322f3037222c22e58f82e695b0223a207b22e6b58be9878fe697b6e997b4223a2022323031362f31322f3037222c22e59381e7a78d223a2022e9998de9a699e9bb84e6aa80222c22e883b8e5be84223a20223220636d222c22e9ab98e5baa6223a202231303020636d222c22e99da2e7a7af223a2022313620e5b9b3e696b9222c22e6a0aae695b0223a20223120e6a0aa22207d7d"
	b,err := hex.DecodeString(jsonStr)
	jsonObject, err := ourjson.ParseObject(string(b))
	fmt.Println(jsonObject, err)

	//user := jsonObject.GetJsonObject("参数")
	//fmt.Println(user)
	//
	//address, err := jsonObject.GetString("地址")
	//fmt.Println(address, err)
	//
	//jsonObject.Put("地址",user)
	//fmt.Println(jsonObject)
	//
	//type Test struct  {
	//	ExtrinsicIndex string `json:"extrinsic_index"`
	//	Hash           string `json:"hash"`
	//}
	//
	//
	//
	//v := Test{"1925","0x13456265395841jdalfjsda"}
	//
	//jsonObject.Put("学习",v)
	fmt.Println(jsonObject)
}
