package main

import (
	"fmt"
	"net/netip"
	"net/url"
)

func main() {
	u, err := url.Parse("ws://127.0.0.1:9888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Host)
	p, _ := netip.ParseAddrPort("127.0.0.1:9888")
	fmt.Println(p.Addr())
}
