package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 6, 6, 5, 7, 7}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
