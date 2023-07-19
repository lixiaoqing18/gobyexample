package dp

import (
	"fmt"
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param triangle int整型二维数组
 * @return int整型
 */
func MinTrace(triangle [][]int) int {
	// write code here
	n := len(triangle)
	state := make([][]int, n)
	for i := 0; i < n; i++ {
		state[i] = make([]int, n)
	}

	state[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				state[i][0] = state[i-1][0] + triangle[i][0]
			} else if j > 0 && j < i {
				state[i][j] = min(state[i-1][j-1], state[i-1][j]) + triangle[i][j]
			} else if j == i {
				state[i][j] = state[i-1][j-1] + triangle[i][j]
			}
		}
	}

	fmt.Println(state)

	minPath := math.MaxInt
	for i := 0; i < len(state); i++ {
		if state[len(state)-1][i] < minPath {
			minPath = state[len(state)-1][i]
		}
	}
	return minPath
}

func min(nums ...int) int {
	minNum := math.MaxInt
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}
