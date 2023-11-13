package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5 ) - 단축연산자

	s := []int{0, 0, 0, 0, 0}
	for _, value := range s {
		fmt.Println(value)
	}
}
