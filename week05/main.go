package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)        //메모리 번지 수를 reader가 가진다?
	inputScore, err := reader.ReadString('\n') //err declared and not used
	fmt.Println(inputScore)
}
