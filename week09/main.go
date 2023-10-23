package main

import "fmt"

func double(n *int) {
	*n = *n * 2
}

func main() {
	var a int = 0
	double(&a)
	fmt.Println(a)
}
