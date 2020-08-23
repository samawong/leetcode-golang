package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4}
	target := 6
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] == target-numbers[i] {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}
