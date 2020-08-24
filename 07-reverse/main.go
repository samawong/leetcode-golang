package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	runes := []rune(strconv.Itoa(x))
	fmt.Println(runes)
	from := 0
	if runes[0] == 45 {
		from = 1
	}
	for to := len(runes) - 1; from < to; from, to = from+1, to-1 {

		runes[from], runes[to] = runes[to], runes[from]
	}
	fmt.Println(runes)
	ans, err := strconv.Atoi(string(runes))
	if err != nil {
		fmt.Println(err)
	}
	if ans > (2.147483648e+09-1) || ans < (-2.147483648e+09) {
		return 0
	}
	return ans
}

func main() {

	A := 1534236469
	fmt.Println(reverse(A))
}
