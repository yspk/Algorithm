package main

import (
	"fmt"
	"github.com/isdamir/gotype"
	"math"
)

//方法功能：找出节点在二叉树中的编号
func getNumber(root *gotype.BNode,node *gotype.BNode,number int) (bool,int) {
	if root == nil {
		return false,number
	}
	if root.Data.(int) == node.Data.(int) {
		return true, number
	}
	num := 2*number
	//如果node节点在root节点的左子树中，左子树编号为当前节点的2倍
	if b,num := getNumber(root.LeftChild,node,num);b{
		return true,num
	} else {
		//如果node节点在root节点的右子树中，右子树编号为当前节点的2倍+1
		num = 2*number+1
		return getNumber(root.RightChild,node,num)
	}
}

func getNodeFromNum(root *gotype.BNode,number int) *gotype.BNode {
	if root == nil || number < 0 {
		return nil
	}  else if number == 1 {
		return root
	}
	//节点编号对应二进制的位数(最高位一定为1，由于根节点代表1)
	ll := (uint)(math.Log(float64(number))/math.Log(2.0))
	//去掉根节点表示的1
	number -= 1<<ll
	for ;ll >0; ll -- {
		//如果这一位二进制的值为1，那么编号为number的节点必定在当前节点的右子树上
		if((1<<(ll-1))&number) == 1 {
			root = root.RightChild
		} else {
			root = root.LeftChild
		}
	}
	return root
}

//查找二叉树中两个节点最近的共同父节点
func FindParentNode(root,node1,node2 *gotype.BNode) *gotype.BNode {
	num1 := 1
	num2 := 1
	_,num1 = getNumber(root,node1,num1)
	_,num2 = getNumber(root,node2,num2)
	//找出编号为num1和num2的共同父节点
	for num1 != num2 {
		if num1 > num2 {
			num1 /= 2
		} else {
			num2 /= 2
		}
	}
	return getNodeFromNum(root,num1)
}

func main() {
	data := []int{
		1,2,3,4,5,6,7,8,9,10,
	}
	fmt.Println("数组：", data)
	root := gotype.ArrayToTree(data,0,len(data)-1)
	node1 := root.LeftChild.LeftChild.LeftChild
	node2 := root.LeftChild.RightChild
	result := FindParentNode(root,node1,node2)
	if result != nil {
		fmt.Println(node1.Data,"与",node2.Data,"的最近公共父节点为:",result.Data)
	}else {
		fmt.Println("没有公共父节点")
	}
}
