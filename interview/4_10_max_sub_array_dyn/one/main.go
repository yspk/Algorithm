package main

import (
	"fmt"
	"math"
)

func minDynDistance(a,b,c []int) int {
	aLen := len(a)
	bLen := len(b)
	cLen := len(c)
	minDis := math.MaxInt
	i := 0 	//数组a的下标
	j := 0	//数组b的下标
	k := 0	//数组c的下标
	for true {
		curDist := Max3(Abs(a[i]-b[j]),Abs(a[i]-c[k]),Abs(b[j]-c[k]))
		if curDist < minDis {
			minDis = curDist
		}
		//找出当前遍历到三个数组中的最小值
		min := Min3(a[i],b[j],c[k])
		if min == a[i] {
			i ++
			if i>=aLen {
				break
			}
		} else if min == b[j] {
			j ++
			if j >= bLen {
				break
			}
		} else {
			k ++
			if k >= cLen {
				break
			}
		}

	}
	return minDis
}

func Min3(i int, i2 int, i3 int) int {
	tmp := i
	if i2 < tmp {
		tmp = i2
	}
	if i3 < tmp {
		tmp = i3
	}
	return tmp
}

func Max3(abs int, abs2 int, abs3 int) int {
	tmp := abs
	if abs2 > tmp {
		tmp = abs2
	}
	if abs3 > tmp {
		tmp = abs3
	}
	return tmp
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func min(dis int, i int) int {
	if dis > i {
		return i
	} else {
		return dis
	}
}

func main() {
	a := []int{3,4,5,7,15}
	b := []int{10,12,14,16,17}
	c := []int{20,21,23,24,30,37}
	fmt.Println("最小距离法")
	fmt.Println(minDynDistance(a,b,c))
}
