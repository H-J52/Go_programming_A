package main

import "fmt"

func main() {
	var a []string
	var b []bool

	// a = make([]string, 4, 5)
	// 슬라이스 - 변수 생성 후 make로 슬라이스 만들기, make로 단축 연산, 슬라이스 리터럴
	b = append(b, true)
	fmt.Printf("%#v %#v\n", a, b)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a, len(b), cap(b))
}
