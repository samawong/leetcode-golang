package main

import "fmt"

func main() {
	A := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(A))

}

//
/*
func maxProfit(prices []int) int {
	profit := 0

	for i, v := range prices {
		for m, n := range prices {
			if v > n && i > m {
				cha := v - n
				if cha > profit {
					profit = cha
				}
			}
		}
	}

	return profit
}
*/

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return profit
}
