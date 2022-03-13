package main

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/xml"
			"io"
	"golang.org/x/net/html/charset"
)

func main()  {
	char := "abcdefghijklmnopqrstuvwxyz"
	for _,v := range char {
		domain := string(v) + ".com"
		go WhoIs(domain)
		for _,v1 := range char {
			domain1 := string(v1) + domain
			go WhoIs(domain1)
			for _,v2 := range char {
				domain2 := string(v2) + domain1
				go WhoIs(domain2)
				for _,v3 := range char {
					domain3 := string(v3) + domain2
					go WhoIs(domain3)
					for _,v4 := range char {
						domain4 := string(v4) + domain3
						go WhoIs(domain4)
						//for _,v5 := range char {
						//	domain5 := string(v5) + domain4
						//	go WhoIs(domain5)
						//	for _,v6 := range char {
						//		domain6 := string(v6) + domain5
						//		go WhoIs(domain6)
						//		for _,v7 := range char {
						//			domain7 := string(v7) + domain6
						//			go WhoIs(domain7)
						//		}
						//	}
						//}
					}
				}
			}
		}
	}


}


func WhoIs(domain string) {
	values := url.Values{}
	values.Add("area_domain", domain)
	apiUrl := "http://panda.www.net.cn/cgi-bin/check.cgi?" + values.Encode()
	req, err := http.NewRequest("GET", apiUrl, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	type Response struct {
		XMLName    xml.Name `xml:"property"`  // 指定最外层的标签为property
		Returncode int  `xml:"returncode"` // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
		SmtpPort int `xml:"smtpPort"`
		Key string `xml:"key"`
		Original string `xml:"original"`
	}
	var who Response

	if resp.StatusCode != http.StatusOK {
		//fmt.Println(resp.StatusCode)
		return
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		return charset.NewReaderLabel(c, i)
	}

	decoder.Decode(&who)
	//fmt.Println(who.Original)

	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// CAN BE REMOVED
	//// logger.Debug(string(b))
	//if err = xml.Unmarshal(b, &who); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	if who.Original != "211 : Domain name is not available" {
		fmt.Println("可用:", who.Key)
	} else {
		//fmt.Println("不可用:", who.Key)
	}
}