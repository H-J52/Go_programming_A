package main

import (
	"fmt"
	"reflect"
)

func main() {
	//변수명은 대문자로 시작
	// 영문 대문자의 경우 다른 패키지에서 접근할 수 있다. 소문자로 시작하는 변수는 동일 패키지에서만 접근 가능하다
	var e = string
	var d = bool
	var c rune
	var b float64
	var a int

	//naming conversation
	var studedentId string
	var i int32

	//a := 7
	fmt.Println(studedentId)
	fmt.Println(i)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)

	fmt.Println(reflect.Typeof(d))
	fmt.Println(reflect.Typeof(e))
}
