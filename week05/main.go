package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)        //메모리 번지 수를 reader가 가진다?
	inputScore, err := reader.ReadString('\n') //option 2
	log.Fatal(err)
	fmt.Println(inputScore)
}
