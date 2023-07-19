package stringmatch

import "sync"

// 存放坏字符最后在模式串中的位置
type BmBc map[rune]int

var bmBc BmBc

// 移动距离为坏字符的位置i减去getBC(c),如果getBC(c)是0，就是i+1
func (bmBc BmBc) getBc(c rune) int {
	if v, ok := bmBc[c]; ok {
		return v
	}
	return -1 //
}

func initBmBC(pattern []rune) {
	bmBc = make(map[rune]int)
	for i, v := range pattern {
		bmBc[v] = i
	}
}

// suff[i]=v表示以第i个字符为后缀的字串与模式串后缀可最长匹配的字符数
var suff []int

func initSuff(pattern []rune) {
	m := len(pattern)
	suff = make([]int, m)

	suff[m-1] = m

	for i := m - 2; i >= 0; i-- {
		j := i
		for j >= 0 && pattern[j] == pattern[m-1-(i-j)] {
			j--
		}

		suff[i] = i - j
	}
}

// bmGS[j]=pos表示好后缀的前一个位置对应的模式串右移的位数
var bmGs map[int]int

func initBmGs(pattern []rune) {
	m := len(pattern)
	bmGs = make(map[int]int)

	// 计算后缀数组
	initSuff(pattern)

	// 先全部赋值为m，包含Case3
	for i := 0; i < m; i++ {
		bmGs[i] = m
	}

	// Case2
	j := 0
	for i := m - 1; i >= 0; i-- {
		if suff[i] == i+1 {
			for ; j < m-1-i; j++ {
				if bmGs[j] == m {
					bmGs[j] = m - 1 - i
				}
			}
		}
	}

	// Case1
	for i := 0; i <= m-2; i++ {
		bmGs[m-1-suff[i]] = m - 1 - i
	}
}

func BM(str string, sub string) int {
	main := []rune(str)
	pattern := []rune(sub)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		initBmBC(pattern)
		wg.Done()
	}()
	go func() {
		initBmGs(pattern)
		wg.Done()
	}()
	wg.Wait()
	lenOfPattern := len(pattern)

	i := 0
	for i < len(main)-len(pattern) {
		j := 0
		for j = lenOfPattern - 1; j >= 0 && main[i+j] == pattern[j]; j-- {

		}
		if j < 0 {
			return i
		} else {
			//计算坏字符移动距离
			bcDis := j - bmBc.getBc(main[i+j])

			//计算好后缀移动距离
			gsDis := bmGs[j]

			max := bcDis
			if gsDis > max {
				max = gsDis
			}
			i += max
		}
	}

	return -1 //not found
}
