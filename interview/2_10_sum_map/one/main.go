package main

type Pairs struct {
	first  int
	second int
}

func FindPairs(arr []int) bool {
	//健为数对的和，值为数对
	sumPair := make(map[int]*Pairs)
	n := len(arr)
	//遍历数组中可能的所有数对
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			//如果这个数对的和在map中没有,则放入map
			sum := arr[i] + arr[j]
			if v, ok := sumPair[sum]; ok {
				println("发现数对", arr[i], "+", arr[j], "=", v.first, "+", v.second)
				return true
			} else {
				sumPair[sum] = &Pairs{arr[i], arr[j]}
			}
		}
	}
	return false
}

func main() {
	arr := []int{
		3, 4, 7, 10, 20, 9, 8,
	}
	FindPairs(arr)
}
