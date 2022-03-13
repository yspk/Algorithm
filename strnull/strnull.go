package strnull

func StringNull(str string) string {
	//ctr := []byte(str)
	if len(str) == 0 {
		return ""
	}
	var originalLength,numberOfBlank int = 0,0
	var i int = 0
	for ;i<len(str);i++ {
		originalLength ++
		if str[i] == ' ' {
			numberOfBlank ++
		}
	}

	var newLength = originalLength + numberOfBlank * 2
	var indexOfOriginal = originalLength-1
	var indexOfNew = newLength
	ctr := make([]byte,newLength)
	copy(ctr,[]byte(str))
	for indexOfOriginal >= 0 && indexOfNew > indexOfOriginal {
		if ctr[indexOfOriginal] == ' ' {
			indexOfNew--
			ctr[indexOfNew] = '0'
			indexOfNew--
			ctr[indexOfNew] = '2'
			indexOfNew--
			ctr[indexOfNew] = '%'
		} else {
			indexOfNew--
			ctr[indexOfNew] = ctr[indexOfOriginal]
		}
		indexOfOriginal --
	}
	return string(ctr)
}
