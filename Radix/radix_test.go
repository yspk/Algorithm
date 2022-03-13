package Radix

import "testing"

func TestRadix(t *testing.T) {
	//duplication1,duplication2,duplication3,duplication4 := new(int),new(int),new(int),new(int)
	// var duplication1,duplication2,duplication3,duplication4 *int //空指针
	var duplication1,duplication2,duplication3,duplication4 int //提防空指针，建议初始化为实例（申请内存空间，不会出现空指针），传参时传地址就好
	idx1 := []int{2,3,1,0,2,5,3}
	idx2 := []int{2,3,1,0,6,5,4}
	idx3 := []int{}
	idx4 := []int{2,7,1,0,2,5,3}

	if Radix(idx1,&duplication1) {
		t.Log(idx1,duplication1)
	}

	if Radix(idx2,&duplication2) {
		t.Log(idx2,duplication2)
	}

	if Radix(idx3,&duplication3) {
		t.Log(idx3,duplication3)
	}

	if Radix(idx4,&duplication4) {
		t.Log(idx4,duplication4)
	}
}
