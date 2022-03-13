package dichotomization

import "testing"

func TestGetDuplication(t *testing.T) {
	t.Log(GetDuplication([]int{2,3,5,4,3,2,6,7}))
	t.Log(GetDuplication([]int{0,1,2,3,4,6,5}))
	t.Log(GetDuplication([]int{}))
	t.Log(GetDuplication([]int{2,3,5,4,3,2,6,7}))
}
