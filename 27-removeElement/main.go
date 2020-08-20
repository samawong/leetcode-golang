package main

import "fmt"

func main() {
	A := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	lenA := removeElement(A, target)

	for i := 0; i < lenA; i++ {
		fmt.Println(A[i])
	}
}

//
//拿Val跟nums的每个数比较，如果第一个数不相等，就把这数放第一位置，第二个数不相等，放第二个，以此重新排列一下
//不相等的次数也就是要返回的新长度
func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if val != nums[j] {

			nums[i] = nums[j]
			i++
		}
	}
	return i
}
