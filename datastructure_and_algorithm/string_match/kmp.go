package stringmatch

//定义 “k-前缀” 为一个字符串的前 k 个字符； “k-后缀” 为一个字符串的后 k 个字符。k 必须小于字符串长度。
//next[x] 定义为： P[0]~P[x] 这一段字符串，使得k-前缀恰等于k-后缀的最大的k.
var next []int

func initNext(pattern []rune) {
	next = make([]int, len(pattern))
	next[0] = 0
	i := 1
	now := 0
	for i < len(pattern) {
		if pattern[now] == pattern[i] {
			now++
			next[i] = now
			i++
		} else if now > 0 {
			now = next[now-1]
		} else {
			next[i] = 0
			i++
		}
	}
}

func KMP(main string, pattern string) int {
	mainStr := []rune(main)
	patternStr := []rune(pattern)
	initNext(patternStr)
	tar := 0
	pos := 0
	for tar < len(mainStr) {
		if mainStr[tar] == patternStr[pos] {
			tar++
			pos++
		} else if pos > 0 {
			pos = next[pos-1]
		} else {
			tar++
		}
		if pos == len(patternStr) {
			return tar - pos //找到
			//pos=next[pos-1]
		}
	}

	return -1
}
