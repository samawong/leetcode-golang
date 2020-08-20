package main

import "fmt"

func main() {
	A := []int{0, 0, 1, 1, 1, 2, 3, 3, 4}
	lenA := removeDuplicates(A)
	fmt.Println(lenA)
	for m := 0; m < lenA; m++ {
		fmt.Println(A[m])
	}
}

//0, 0, 1, 1, 1, 2, 3, 3, 4
//i  j
//nums[i] == nums[j]时什么都不做，如第1个0 和第2个0
//nums[i] ！= nums[j] 如第1个0 和 第3个1不相等，则i+1，然后到第2个0的位置，再让第3个的1赋值到第2个0的位置，重复该动作
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
