package main

import "fmt"

func main() {
	A := []int{9, 9}
	fmt.Println(plusOne(A))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 { //从最后一位向前都是9的话，每一位改成0
			digits[i] = 0
		} else { //都不是9的话，每一位+1
			digits[i]++
			return digits
		}
	}
	digits = make([]int, len(digits)+1) //进位，数组长度+1
	digits[0] = 1                       //第一位设置为1
	return digits
}
