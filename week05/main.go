package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)      //메모리 번지 수를 reader가 가진다?
	inputScore, _ := reader.ReadString('\n') //option 1, Igonore the error return value with the..
	fmt.Println(inputScore)
}
