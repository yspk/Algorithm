package main

import (
	"./trie"
	"./hashset"
	"strconv"
	"fmt"
	"strings"
)

func main()  {
	tr := trie.NewTrieNode()
	tr.AddTrieNode("say",1)
	tr.AddTrieNode("she",2)
	tr.AddTrieNode("shr",3)
	tr.AddTrieNode("her",4)
	tr.AddTrieNode("he",5)

	tr.BuildFailNodeBFS()
	s := "yasherhs"
	h := hashset.NewHashSet()
	tr.SearchAC(s,h)

	fmt.Printf("在主串%s中存在模式串的编号为:%v",s,strings.Join(h,","))
}





















