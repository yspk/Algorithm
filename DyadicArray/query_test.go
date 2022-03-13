package DyadicArray

import "testing"

func TestQuery(t *testing.T) {
	ids := [][]int{
		{1,2,8,9},
		{2,4,9,12},
		{4,7,10,13},
		{6,8,11,15},
	}

	//ids1 := [][]int{
	//	{},
	//	{},
	//	{},
	//	{},
	//}

	t.Log(Query(ids,13))
}
