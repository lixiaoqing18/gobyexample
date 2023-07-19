package stringmatch

func BF(str string, sub string) int {
	mainStr := []rune(str)
	patternStr := []rune(sub)
	for i := 0; i < len(mainStr)-len(patternStr); i++ {
		if compareSubstr(mainStr[i:len(patternStr)+i], patternStr) {
			return i
		}
	}

	return -1 //not found
}
