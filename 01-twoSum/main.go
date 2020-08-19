package main

func main() {
	nums := []int{1, 3, 5, 6, 7}
	target := 8
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	//m := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}

	}
	return nil
}

/*
  func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		other := target - nums[i]
		if _,ok := m[other];ok {
			return []int{m[other],i}
		}
		m[nums[i]] = i
	}
	return nil
  }
*/
