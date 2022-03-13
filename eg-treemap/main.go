package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/zencoder/go-smile/smile"

)

func main() {
	jsonStr := `{
		"call_module": "区块链证书",
		"call_module_function": "notary",
		"extrinsic_hash": "0x53a0a660e200ef32549a3f66093a7ee49a759b8cd71efa4b61ff26114a1b5e48",
		"block_time":" 2021-12-29 14:23:39",
		"info": {
			"token_id": "93911476072456014925824",
			"token_class_id":"",
			"owner": "",
			"content_hash": "0xb47e3cd837ddf8e4c57f05d70ab865de6e193baa",
			"content": {
				"林木编号": "469007009009JE00001L00019003",
				"GPS": "109.024986,19.013589",
				"地址": "海南省白沙黎族自治县细水乡白石岭",
				"育苗时间": "2015/12/07",
				"栽种时间": "2016/12/07",
				"参数": {
					"测量时间": "2016/12/07",
					"品种": "降香黄檀",
					"胸径": "2 cm",
					"高度": "100 cm",
					"面积": "16 平方",
					"株数": "1 株"
				}
			}
		}
	}
    `

	object, err := smile.DecodeToObject([]byte(jsonStr))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(object)
	}

	deserialized := linkedhashmap.New()
	err = deserialized.FromJSON([]byte(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	v,_ := deserialized.Get("info")
	fmt.Println(v)

	serialized, err := deserialized.ToJSON()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(serialized))

}
