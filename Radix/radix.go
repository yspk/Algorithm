package Radix

// 利用基数排序查找数组中的的一个重复元素
func Radix(idx []int, duplication *int) bool {
	l := len(idx)
	if l == 0 {
		return false
	}

	for i := 0;i< l ; i++ {
		if idx[i] < 0 || idx[i] > l - 1 {
			return false
		}
	}

	for i:= 0;i< l ; i++  {
		if idx[i] != i {
			if idx[i] == idx[idx[i]] {
				*duplication = idx[i]
				return true
			}

			idx[i],idx[idx[i]] = idx[idx[i]],idx[i]
		}
	}
	return false
}


