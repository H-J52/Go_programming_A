package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	// var primes [3]int = [3]int{2, 3, 5}
	// fmt.Println(primes, primes[1])

	primes := [3]int{2, 3, 5}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) //bool의 제로값은 false

	// fmt.Println(primes[4]) //invalid argument: index 4 out of bounds [0:3]

	i := 0

	for i < 4 {
		fmt.Println(primes[i]) //panic: runtime error: index out of range [3] with length 3
		i++
	}
}
