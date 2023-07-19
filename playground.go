package main

import (
	"fmt"

	"github.com/lixiaoqing18/gobyexample/datastructure_and_algorithm/dp"
)

func main() {
	//str := "竹杖芒鞋轻胜马，谁怕？"
	//fmt.Printf("%c\n", []rune(str))

	//recall.EightQueen()

	//w := recall.Knapsack(10, 4, [][]int{{1, 3}, {10, 4}, {2, 5}, {7, 6}})
	//fmt.Println("max weight: ", w)

	//fmt.Println(recall.Match("qingqing2jiajiab", "q*i*2?j*a."))

	//triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 2, 8, 3}}
	//fmt.Println(dp.MinTrace(triangle))

	//fmt.Println(dp.CoinChange([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(dp.CoinChange_DP([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(dp.CoinChange([]int{83, 186, 408, 419}, 6249))
	//fmt.Println(dp.CoinChange([]int{2}, 3))

	//fmt.Println(dp.MaxSumArray([]int{-2, 1, -3, 1, -1, 6, 2, -5, 5, 1}))
	//fmt.Println(dp.MaxSumArray2([]int{-2, 1, -3, 1, -1, 6, 2, -5, 5, 1}))
	fmt.Println(dp.MaxSumArray3([]int{-2, 1, -3, 1, -1, 6, 2, -5, 5, 1}))
}
