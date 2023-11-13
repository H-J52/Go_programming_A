package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5 ) - make 함수를 이용하여 슬라이스 생성 후 메모리 할당, 제로 값 적용

	// s := []int{0, 0, 0, 0, 0} //슬라이스 리터럴 이용해 슬라이스 생성 및 메모리 할당, 초기화 진행
	// for _, value := range s {
	// 	fmt.Println(value)
	// }
	s := []int{0, 0, 0, 0, 0}

	s[4] = 99
	s[2] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "student"} //배열 리터럴을 이용해서 test 배열 생성
	testS := test[0:2]
	testS2 := test[1:]
	testS2[0] = "python"
	fmt.Println(testS2)
	fmt.Println(testS, len(testS))
}
