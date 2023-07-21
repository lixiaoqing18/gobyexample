package dp

func LastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	//背包容量
	w := sum / 2
	//物品数量
	n := len(stones)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			curStore := stones[i-1]
			if j < curStore {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-curStore]+curStore)
			}
		}
	}
	return sum - 2*dp[n][w]
}
