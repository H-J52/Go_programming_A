package main

import (
	"fmt"
	"log"
)

// func main() {
// 	var c rune = '가' //가에 대한 유니코드? 44032 (rune은 4바이트?)

// 	fmt.Println(c)        //유니코드(UTF-8) 값 출력
// 	fmt.Printf("%c\n", c) //한 글자 출력
// 	fmt.Printf("%T\n", c) //rune은 결국 int32의 별명

// 	//var a int16 = 7
// 	//var a = 7
// 	//a를 두번 선언하면 값은 바뀔 수 있으나 타입은 변하지 않는다(먼저 int가 선언된 후 float형으로 나중에 값을 선언하면 오류남)
// 	a := 7
// 	var b float64 = 5.3
// 	a = int(b) //Type Conversion, Casting
// 	d := false

// 	fmt.Println(d)
// 	fmt.Printf("%T\n", d)

// 	fmt.Println(a)s
// 	fmt.Printf("%T\n", a) //a의 타입을 출력

// 	fmt.Println(math.Round(2.71)) //round는 반올림 ceil은 올림 floor은 내림
// 	fmt.Println("Hello Go!")
// 	fmt.Println(strings.Title("go git github java"))
// }

//===========================

// func main() {
// 	var e string
// 	var d bool
// 	var b float64
// 	var c rune = '가'
// 	a := 7

// 	fmt.Println(b) //제로 값 0
// 	fmt.Println(e) //제로 값 ""
// 	fmt.Println(d) //제로 값 false

// 	fmt.Printf("%T\n", c)
// 	fmt.Printf("%T\n", a)

// 	fmt.Println(reflect.TypeOf(c))
// 	fmt.Println(reflect.TypeOf(a))
// }

//===========================

// func main() {
// 	texts := "GO M@ney"
// 	fmt.Println(texts)
// 	a := strings.NewReplacer("@", "o")
// 	newTexts := a.Replace(texts)

// 	fmt.Println(newTexts)
// }

//===========================

// func main() {
// 	fmt.Println("단 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)

// 	in, err := rd.ReadString('\n')
// 	if err != nil { //에러가 발생하지 않는다면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	dan, err := strconv.ParseInt(in, 10, 32) //dan, err := strconv.Atoi(in)
// 	if err != nil { //에러가 발생하지 않는다면
// 		log.Fatal(err)
// 	}

// 	for i := 1; i < 10; i++ {
// 		fmt.Println(dan, " * ", i, " = ", (int(dan) * i)) //Atoi를 쓰면 fmt.Println(dan, " * ", i, " = ", (dan) * i))
// 	}

// }

//===========================

// func main() {
// 	fmt.Println("단 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)

// 	in, err := rd.ReadString('\n')
// 	if err != nil { //에러가 발생하지 않는다면
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)
// 	dan, err := strconv.ParseInt(in, 10, 32) //dan, err := strconv.Atoi(in)
// 	if err != nil { //에러가 발생하지 않는다면
// 		log.Fatal(err)
// 	}

// 	//go에서 while처럼 쓰기
// 	i:=1
// 	for i < 10;{
// 		fmt.Println(dan, " * ", i, " = ", (int(dan) * i)) //Atoi를 쓰면 fmt.Println(dan, " * ", i, " = ", (dan) * i))
// 		i++
// 	}

// }

//===========================

// func main() {
// 	//seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// }

//===========================

// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	number := rand.Intn(150) + 2 //0. 1제외 2~151 사이 값
// 	fmt.Println("임의 수 : ", number)

// 	count := 0
// 	for i := 2; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

//===========================

// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	number := rand.Intn(150) + 2 //0. 1제외 2~151 사이 값
// 	fmt.Println("임의 수 : ", number)

// 	count := 0
// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count = count + 1
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

//===========================

// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2 //0. 1제외 2~151 사이 값
// 	fmt.Println("임의 수 : ", number)

// 	for i := 1; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 		fmt.Print(i, " ")
// 	}

// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

func main() {
	var number int
	fmt.Println("단 입력 : ")
	n, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal(err)
	}

	i := 1
	for i < 10 {
		fmt.Println(n, " * ", i, " = ", int(n)*i)
		i++
	}
}
