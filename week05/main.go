package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings" // 필요한 패키지 추가
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 올바른 시간 시드 설정
	answer := rand.Intn(100) + 1 // (1 ~ 100) 사이의 난수 생성
	fmt.Println("Guess Number Game!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input guess number : ")
	inputNumberString, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	inputNumberString = strings.TrimSpace(inputNumberString) // 올바른 변수 이름 사용
	inputNumber, err := strconv.Atoi(inputNumberString)

	if err != nil {
		log.Fatal(err)
	}

	if inputNumber < answer {
		fmt.Println("Guess number is lower than answer") // 정답보다 입력한 숫자가 작음
	} else if inputNumber > answer {
		fmt.Println("Guess number is higher than answer") // 정답보다 입력한 숫자가 큼
	} else {
		fmt.Println("Congratulations! You guessed it!") // 정답을 맞춤
	}
}
