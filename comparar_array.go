package main

import "fmt"

func Comparar() bool {

	a := []int{1, 2, 4}
	b := []int{1, 2, 4}

	if len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
}

func main() {
	fmt.Println(Comparar())
}
