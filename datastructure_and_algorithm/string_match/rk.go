package stringmatch

import "math"

const Radix = 128

var data []int

func init() {
	for i := 0; i < 100; i++ {
		data = append(data, int(math.Pow(Radix, float64(i))))
	}
}

func RK(str, match string) int {
	mainStr := []rune(str)
	matchStr := []rune(match)
	matchHash := hash(matchStr)
	lenOfMatch := len(matchStr)
	hashCount := len(mainStr) - len(matchStr)
	hashcode := 0
	for i := 0; i < hashCount; i++ {
		if i == 0 {
			hashcode = hash(mainStr[0:lenOfMatch])
		} else {
			hashcode = hashForNextSubStr(hashcode, mainStr[i-1], mainStr[i+lenOfMatch-1], lenOfMatch)
		}
		if hashcode == matchHash {
			//为防止hash冲突，可以逐字符进行比较
			if compareSubstr(mainStr[i:i+lenOfMatch], matchStr) {
				return i
			}
		}
	}
	return -1
}

func hash(str []rune) int {
	var hashcode int
	for i, v := range str {
		hashcode += int(v-'a') * data[len(str)-1-i]
	}
	return hashcode
}

func hashForNextSubStr(curHash int, firstCharOfCurSubStr rune, lastCharOfNextSubStr rune, lenOfMatch int) int {
	return (curHash-int(firstCharOfCurSubStr-'a')*data[lenOfMatch-1])*Radix + int(lastCharOfNextSubStr-'a')
}
