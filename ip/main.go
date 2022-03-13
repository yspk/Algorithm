package main

import (
	"bufio"
	"fmt"
	"coding.net/baoquan2017/dataqin-backend/common/util"
	"os"
	"time"
)

func main()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %n\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	var cn,aboard int
	for line, n := range counts {
		time.Sleep(time.Second)
		if getIpInfo(line) {
			cn += n
		} else {
			aboard += n
		}
		fmt.Println(line,n)
	}
	fmt.Println(cn,aboard)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

type IpInfo struct {
	Status string `json:"status"`
	Country string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region string `json:"region"`
	RegionName string `json:"regionName"`
	City string `json:"city"`
	Zip string `json:"zip"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Timezone string `json:"timezone"`
	Isp string `json:"isp"`
	Org string `json:"org"`
	As string `json:"as"`
	Query string `json:"query"`
}

func getIpInfo(ip string) bool {
	apiUrl := "http://ip-api.com/json/" + ip
	var response IpInfo
	if err := util.Get(apiUrl, &response, nil); err != nil {
		fmt.Println(err)
		return false
	}
	if response.CountryCode == "CN" {
		return true
	}
	return false
}
