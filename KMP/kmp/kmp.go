package kmp

func KMP(bigstr,smallstr string) int {
	i := 0
	j := 0
	lb:= len(bigstr)
	ls := len(smallstr)
	next := GetNextVal(smallstr)

	for ;i<lb && j < ls;{
		if j == -1 || bigstr[i] == smallstr[j] {
			i ++
			j ++
		}else {
			j = next[j]
		}
	}

	if j ==ls {
		return i-ls
	}
	return -1
}

func GetNextVal(smallstr string) []int {
	k := -1
	j := 0
	l := len(smallstr)
	next := make([]int,l)
	next[j] = -1

	for ;j < l-1; {
		if k == -1 || smallstr[k] == smallstr[j] {
			j++
			k++
			next[j] = k
		}else {
			k = next[k]  //递归
		}
	}
	return next
}
