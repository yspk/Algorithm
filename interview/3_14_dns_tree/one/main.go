package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

var CharCount = 11

type DNSCache struct {
	root *gotype.TrieNode
}

func (p *DNSCache) getIndexFromRune(char rune) int {
	if char == '.' {
		return 10
	} else {
		return int(char)-'0'
	}
}

func (p *DNSCache) getRuneFromIndex(i int) rune {
	if i == 10 {
		return '.'
	} else {
		return rune('0'+i)
	}
}

//把一个IP地址和相应的URL添加到Trie树种，最后一个节点是URL
func (p *DNSCache) Insert(ip,url string) {
	pCrawl := p.root
	for _,v := range []rune(ip) {
		//根据当前遍历到的IP中的字符，找出子节点的索引
		index := p.getIndexFromRune(v)
		//如果子节点不存在，则创建一个
		if pCrawl.Child[index] == nil {
			pCrawl.Child[index] = gotype.NewTrieNode(CharCount)
		}
		//移动到子节点
		pCrawl = pCrawl.Child[index]
	}
	//在叶子节点中存储IP地址对应的URL
	pCrawl.IsLeaf = true
	pCrawl.Url = url
}

func (p *DNSCache) SearchDNSCache(ip string) string {
	pCrawl := p.root
	for _,v := range []rune(ip) {
		index := p.getIndexFromRune(v)
		if pCrawl.Child[index] == nil {
			return ""
		}
		pCrawl = pCrawl.Child[index]
	}
	//返回找到的URL
	if pCrawl != nil && pCrawl.IsLeaf {
		return pCrawl.Url
	}
	return ""
}

func NewDNSCache() *DNSCache {
	return &DNSCache{root: gotype.NewTrieNode(CharCount)}
}

func main() {
	ipAdds := []string{"10.57.11.127","121.57.61.129","66.125.100.103"}
	urls := []string{
		"www.samsung.com",
		"www.samsung.net",
		"www.samsung.org",
	}
	dnsCache := NewDNSCache()
	//把IP地址和对应的URL插入到Trie中
	for i,v := range ipAdds {
		dnsCache.Insert(v,urls[i])
	}
	ip := ipAdds[1]
	result := dnsCache.SearchDNSCache(ip)
	if result != "" {
		fmt.Println("找到了IP对应的URL：",ip,"---->",result)
	} else {
		fmt.Println("没有找到对应的URL")
	}
}
