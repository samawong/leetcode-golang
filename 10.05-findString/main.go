package main

import "fmt"

func main() {
	w := [](string){"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s := "ball"
	fmt.Println(findString(w, s))
}

func findString(words []string, s string) int {
	m := -1
	for i := 0; i < len(words); i++ {
		if s == words[i] {
			m = i
		}
	}
	return m
}
