package dichotomization

func GetDuplication(idx []int) int {
	l := len(idx)
	if l == 0 {
		return -1
	}

	var start int = 1
	var end int = l -1
	for end >= start {
		var middle int = ((end-start)>>1) + start
		var count = countRange(idx,start,middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > middle - start + 1 {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}

func countRange(idx []int,start,end int) int {
	if len(idx) == 0 {
		return 0
	}

	var count = 0;
	for i := 0; i < len(idx);i++ {
		if idx[i] >= start && idx[i] <= end {
			 count ++
		}
	}
	return count
}