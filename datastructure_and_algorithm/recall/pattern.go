package recall

import "fmt"

var matched = false

func Match(str, pattern string) bool {
	strs := []rune(str)
	patterns := []rune(pattern)
	rMatch(0, 0, strs, len(strs), patterns, len(patterns))
	return matched
}

func rMatch(sIndex int, pIndex int, strs []rune, lenOfS int, patterns []rune, lenOfP int) {
	fmt.Println("sIndex:", sIndex, "pIndex:", pIndex)
	if matched {
		return
	}
	if pIndex == lenOfP {
		if sIndex == lenOfS {
			matched = true
		}
		return
	}

	if patterns[pIndex] == '*' {
		for i := 0; i <= lenOfS-sIndex; i++ {
			rMatch(sIndex+i, pIndex+1, strs, lenOfS, patterns, lenOfP)
		}
	} else if patterns[pIndex] == '?' {
		rMatch(sIndex, pIndex+1, strs, lenOfS, patterns, lenOfP)
		rMatch(sIndex+1, pIndex+1, strs, lenOfS, patterns, lenOfP)
	} else if patterns[pIndex] == '.' {
		rMatch(sIndex+1, pIndex+1, strs, lenOfS, patterns, lenOfP)
	} else if sIndex < lenOfS && strs[sIndex] == patterns[pIndex] {
		rMatch(sIndex+1, pIndex+1, strs, lenOfS, patterns, lenOfP)
	}
}
