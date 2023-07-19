package dp

import (
	"math"
	"sort"
)

func CoinChange(coins []int, amount int) int {
	mem := make(map[int]int)
	sort.Ints(coins)
	return CoinChangeImpl(coins, amount, mem)
}

func CoinChangeImpl(coins []int, amount int, mem map[int]int) int {
	if v, ok := mem[amount]; ok {
		return v
	}
	//sort.Ints(coins)
	if amount == 0 {
		return 0
	}
	min := math.MaxInt
	for i := len(coins) - 1; i >= 0; i-- {
		coin := coins[i]
		if amount < coin {
			continue
		}
		rest := amount - coin
		//fmt.Println("amout ", amount, " - ", coin, "=", rest)
		restCount := CoinChangeImpl(coins, rest, mem)
		if restCount == -1 {
			continue
		} else {

			//return 1 + restCount
			if 1+restCount < min {
				min = 1 + restCount
				//fmt.Println("min=", min)
			}
		}
	}

	if min < math.MaxInt {
		mem[amount] = min
		return min
	} else {
		mem[amount] = -1
		return -1
	}
}

func CoinChange_DP(coins []int, amount int) int {
	mem := map[int]int{}
	mem[0] = 0
	for i := 1; i <= amount; i++ {
		min := math.MaxInt
		for j := len(coins) - 1; j >= 0; j-- {
			coin := coins[j]
			if i < coin {
				continue
			}
			rest := i - coin
			restCount := mem[rest]
			if restCount == -1 {
				continue
			} else {
				if 1+restCount < min {
					min = 1 + restCount
				}
			}
		}
		if min < math.MaxInt {
			mem[i] = min
		} else {
			mem[i] = -1
		}
	}

	return mem[amount]
}
