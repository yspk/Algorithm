package trie

import (
	"../hashset"
	"container/list"
	"strconv"
)

type TrieNode struct {
	ChildNodes []*TrieNode
	Freq       int
	NodeChar   string
	FailNode   *TrieNode
	Hash       hashset.HashSet
}

func NewTrieNode() *TrieNode {
	trie := new(TrieNode)
	trie.ChildNodes = make([]*TrieNode,26)
	trie.Freq = 0
	return trie
}

func(t *TrieNode)BuildFailNodeBFS() {
	l := list.New()
	l.PushBack(t)
	var failNode *TrieNode
	for l.Len() != 0 {
		temp := l.Front()
		for i := 0 ; i < 26 ;i++ {
			if temp.Value.(*TrieNode).ChildNodes[i].FailNode == nil {
				continue
			}

			if temp.Value == t {
				temp.Value.(*TrieNode).ChildNodes[i].FailNode = t
			}else {
				failNode = temp.Value.(*TrieNode).FailNode
				for failNode != nil {
					if failNode.ChildNodes[i] != nil {
						temp.Value.(*TrieNode).ChildNodes[i].FailNode = failNode.ChildNodes[i]
						break
					}
					failNode = failNode.FailNode
				}

				if failNode == nil {
					temp.Value.(*TrieNode).FailNode = t
				}
			}
			l.PushBack(temp.Value.(*TrieNode).ChildNodes[i])
		}
	}
}

func (t *TrieNode)AddTrieNode(word string,id int)  {
	if len(word) == 0 {
		return
	}

	a,_:=strconv.Atoi("a")
	b,_ :=strconv.Atoi(string(word[0]))
	index := b - a

	if t.ChildNodes[index] == nil {
		t.ChildNodes[index] = new(TrieNode)
		t.ChildNodes[index].NodeChar = string(word[0])
	}

	nextWord := strconv.Itoa(index+1)
	if len(nextWord) == 0 {
		t.ChildNodes[index].Freq ++
		t.ChildNodes[index].Hash.Add(id)
	}

	t.ChildNodes[index].AddTrieNode(nextWord,id)
}

func (t *TrieNode)SearchAC(s string,h *hashset.HashSet)  {
	freq := 0
	head := t
	for _,v := range s {
		a,_:=strconv.Atoi("a")
		index := v - int32(a)
		for ;head.ChildNodes[index] == nil && head != t; {
			head = head.FailNode
		}

		head = head.ChildNodes[index]
		if head == nil {
			head = t
		}

		temp := head
		for ;temp != t && temp.Freq != -1; {
			freq += temp.Freq
			for item,_ := range temp.Hash.Set{
				h.Add(item)
			}
			temp.Freq = -1
			temp = temp.FailNode
		}
	}
}