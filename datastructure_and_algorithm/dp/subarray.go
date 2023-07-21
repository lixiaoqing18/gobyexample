package dp

import (
	"fmt"
	"math"
)

func MaxSumArray(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
	}

	max := math.MinInt

	for x := len(nums) - 1; x >= 0; x-- {
		for y := x; y <= len(nums)-1; y++ {
			if x == y {
				dp[x][y] = nums[x]

			} else {
				dp[x][y] = nums[x] + dp[x+1][y]
			}

			fmt.Printf("dp[%d][%d]=%d\n", x, y, dp[x][y])
			if max < dp[x][y] {
				max = dp[x][y]

			}
		}
	}

	return max
}

func MaxSumArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func MaxSumArray3(nums []int) int {
	maxSum := nums[0]
	p := nums[0]
	for i := 1; i < len(nums); i++ {
		p = max(nums[i], p+nums[i])
		maxSum = max(maxSum, p)
	}
	return maxSum
}

func max(nums ...int) int {
	maxNum := math.MinInt
	for _, v := range nums {
		if maxNum < v {
			maxNum = v
		}
	}
	return maxNum
}
