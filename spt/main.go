package main

import (
	"fmt"
	"github.com/Cocoon-break/speedtest"
)

func main() {
	report, err := speedtest.ByDistance("en0", 60)
	if err != nil {
		fmt.Printf("failed:%s", err.Error())
		return
	}
	fmt.Printf("%+v", report)
}
