package main

func main() {
	B := []int{0, 1, 1}
	prefixesDivBy5(B)
}

/*
func prefixesDivBy5(A []int) []bool {
	ret := make([]bool, len(A))
	tep := 0
	for i, v := range A {
		fmt.Println(i, v)
		tep = tep*2 + v
		tep %= 5
		fmt.Println(tep)
		ret[i] = tep == 0
	}
	return ret
}
*/

func prefixesDivBy5(A []int) []bool {
	asw := []bool{}
	num := 0
	for i := 0; i < len(A); i++ {
		num <<= 1
		num += A[i]
		num %= 10
		asw = append(asw, num%5 == 0)
	}
	return asw
}
