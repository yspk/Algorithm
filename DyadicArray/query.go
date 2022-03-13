package DyadicArray

func Query(ids [][]int,number int) bool {
	rows,columns := len(ids),len(ids[0])
	if rows == 0 || columns == 0 {
		return false
	}

	var row,column int = 0, columns -1
	for row < rows && column >= 0 {
		if ids[row][column] == number {
			return true
		} else if ids[row][column] > number {
			column --
		} else {
			row ++
		}
	}
	return false
}
