package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 올바른 시간 시드 설정

	fmt.Println("Guess Number Game!")

	answer := rand.Intn(100) + 1 // (1 ~ 100) 사이의 난수 생성
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")
	fmt.Println(answer)

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Input your guess: ")
		inputNumberString, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)

		if err != nil {
			log.Fatal(err)
		}

		if inputNumber < answer {
			fmt.Println("답보다 숫자가 낮습니다.")
		} else if inputNumber > answer {
			fmt.Println("답보다 숫자가 높습니다.")
		} else {
			fmt.Println("Congratulations! You guessed it!")
			break // 정답을 맞췄으므로 반복문 종료
		}
	}
}
