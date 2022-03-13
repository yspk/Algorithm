package dplan

import "testing"

func TestCoinChange(T *testing.T)  {
	T.Log(coinChange([]int{1,2,5,10},199))
}
