package main

func main() {
	A := []int{1, 3, 4, 5}
	target := 4
	searchInsert(A, target)
}

//解体思路
//target与nums内对比，要是相等就返回该数值的索引，
//同时，因为数组从小到大排列，所以target只要比左边的大就行
func searchInsert(nums []int, target int) int {
	index := 0

	for i, v := range nums {
		if target == v {
			index = i
		} else if target > v {
			index = i + 1
		}
	}
	return index
}

/*
//暴力类比
func searchInsert(nums []int, target int) int {
    index := 0
    for i, v := range nums {
        if target > nums[len(nums)-1] {
            index = len(nums)
        }
        if target < nums[0] {
            index = 0
        }
        if target == v {
            index = i
        }
        for _, j := range nums {
            if target > v && target < j {
                index = i + 1
            }
        }
    }
    return index
}
*/
