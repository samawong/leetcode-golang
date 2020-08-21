package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 5, 6}

	//nums1 = append(nums1, nums2...)
	//sort.Ints(nums1)
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

}
