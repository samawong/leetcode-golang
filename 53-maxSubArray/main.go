package main

import "fmt"

func main() {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(A))
}

//扫描法
//
func maxSubArray(nums []int) int {
	sum, cur := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if cur < 0 { //当前数小于0，舍去
			cur = nums[i]
		} else { //当前数大于0，则相加
			cur += nums[i]
		}
		if cur > sum { //当前数大于和，则返回
			sum = cur
		}
	}
	return sum
}
