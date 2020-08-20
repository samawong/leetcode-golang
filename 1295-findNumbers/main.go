package main

import (
	"fmt"
	"strconv"
)

func main() {
	A := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(A)) //12，7896两个
}

//
//strconv.Itoa是将int转为string，数字12就转成了“12”，变成了两位数，然后再求这个string的长度，用长度对2求余即可
func findNumbers(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if len(strconv.Itoa(nums[i]))%2 == 0 {
			result++
		}
	}
	return result
}
