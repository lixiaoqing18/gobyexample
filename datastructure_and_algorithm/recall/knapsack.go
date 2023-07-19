package recall

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算01背包问题的结果
 * @param V int整型 背包的体积
 * @param n int整型 物品的个数
 * @param vw int整型二维数组 第一维度为n,第二维度为2的二维数组,vw[i][0],vw[i][1]分别描述i+1个物品的vi,wi
 * @return int整型
 */
var maxWeight = 0
var choose []int

func Knapsack(V int, n int, vw [][]int) int {
	choose = make([]int, n)
	// write code here
	f(0, V, n, vw, 0, 0)

	return maxWeight
}

func f(i int, V int, n int, vw [][]int, cv int, cw int) {
	if cv == V || i == n {
		fmt.Println("all weight is ", cw)
		fmt.Println("the choose is ", choose)
		if cw > maxWeight {
			maxWeight = cw
		}
		return
	}
	choose[i] = 0
	f(i+1, V, n, vw, cv, cw)
	if cv+vw[i][0] <= V {
		choose[i] = 1
		f(i+1, V, n, vw, cv+vw[i][0], cw+vw[i][1])
	}
}
