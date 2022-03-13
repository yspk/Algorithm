package main

import (
	"fmt"
	"github.com/simplechain-org/go-simplechain/common"
	"time"
)

func main()  {
	fmt.Println("start")
	regesiter()
	var begin uint64 = 2
	var count uint64 = 10000
	var i uint64
	for i=0; begin< 11111111111; i++ { //跑到100000000跑不动了
		tstart := time.Now()
		testinviter(begin,count)
		fmt.Printf("Insert,begin=%d,count=%d,elapsed=%s \n", begin, count, common.PrettyDuration(time.Since(tstart)))
		begin += count
	}
}

type Inviter struct {
	Children []uint64
	Deep  [10]uint64
	Parent uint64
}



var relationShip = make(map[uint64]*Inviter)
var roots = make(map[uint64]uint64)
//注册函数
func regesiter()  {
	var deep [10]uint64
	roots[1] = 1
	relationShip[1] = &Inviter{
		Children:make([]uint64,0),
		Deep:deep,
		Parent:0,
	}
}

func inviter(people,_inviter uint64) uint64 {
	if _inviter != 0 {
		if _,ok := relationShip[people]; !ok {
			// 邀请人为根节点,people成为其直接孩子
			if _,ok := roots[_inviter]; ok && relationShip[_inviter].Deep[0] < 10 {
				var deep [10]uint64
				relationShip[people] = &Inviter{
					Children:make([]uint64,0),
					Deep:deep,
					Parent:_inviter,
				}
				relationShip[_inviter].Children  = append(relationShip[_inviter].Children,people)
				relationShip[_inviter].Deep[0] ++
				return _inviter
				// 邀请人为子节点或邀请人为根节点但第一层孩子满了
			} else {
				// 找到“大儿子”
				parent := wideTraversal(_inviter)
				if parent == 0 {
					return 0
				}
				var deep [10]uint64
				relationShip[people] = &Inviter{
					Children:make([]uint64,0),
					Deep:deep,
					Parent:parent,
				}
				relationShip[parent].Children = append(relationShip[parent].Children,people)
				relationShip[parent].Deep[0] ++

				var depth uint64 = 2
				for p := relationShip[parent].Parent; p != 0 && depth <= 10 ; p = relationShip[p].Parent {
					relationShip[p].Deep[depth-1] ++
					depth ++
				}
				return parent

			}
		}
	}
	return 0
}


func wideTraversal(inviterI uint64) uint64 {
	relation := relationShip[inviterI]
	var i uint64
	//找到第i层孩子没满
	for ;i < 10 && relation.Deep[i] ==  pow(10,i+1); {
		i ++
	}
	if i == 0 {
		return inviterI
	} else if i == 10 {
		return 0
	} else {
		return FindChildren(i,relation)
	}
}

func FindChildren(i uint64,rela *Inviter)  uint64 {
	var j uint64
	// 找到第一个孩子，其第i-1层孩子没满
	for ;j < rela.Deep[0] && relationShip[rela.Children[j]].Deep[i-1] == pow(10,i); j++ {

	}

	if i == 1 {
		return rela.Children[j]
	} else {
		return  FindChildren(i-1,relationShip[rela.Children[j]])
	}
}

func testinviter(begin,count uint64) {
	var tmp uint64
	var i uint64
	for i = 0; i < count;i++ {
		if tmp == 0 {
			tmp = inviter(begin+i,1)
		} else {
			if relationShip[tmp].Deep[0] != 10 {
				tmp = inviter(begin+i,tmp)
		} else {
			 	p := relationShip[tmp].Parent
			 	var depth uint64
				for depth = 2; p != 0 && depth <= 10 ; p = relationShip[p].Parent {
					if relationShip[p].Deep[depth-1] !=  pow(10,depth) {
						tmp = inviter(begin+i,p)
						break
					}
					depth ++
				}
				if p==0 {
					tmp = inviter(begin+i,1)
				}
			}
		}
	}
}

func pow(x, n uint64) uint64 {
	var ret uint64 = 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}