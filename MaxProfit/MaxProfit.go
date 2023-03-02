package main

import (
	"fmt"
)

func MaxProfit(prices []int, i int) int {
	if len(prices) <= 1 {
		return -1
	}
	// profit[j][k] stores the maximum profit gained by doing
	// at most `i` transactions till k'th day
	profit := make([][]int, i+1)
	for j := range profit {
		profit[j] = make([]int, len(prices))
	}

	for j := 0; j <= i; j++ {
		for k := 0; k < len(prices); k++ {
			if j == 0 || k == 0 {
				profit[j][k] = 0
			} else {
				maxSoFar := 0
				for l := 0; l < k; l++ {
					currPrice := prices[k] - prices[l] + profit[j-1][l]
					if maxSoFar < currPrice {
						maxSoFar = currPrice
					}
				}
				profit[j][k] = max(profit[j][k-1], maxSoFar)
			}
		}
	}
	return profit[i][len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{4, 11, 2, 20, 59, 80}
	fmt.Println(MaxProfit(arr, 2))
}
