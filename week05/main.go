package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScoreString = strings.TrimSpace(inputScoreString)      //remove space bar
	inputScore, err := strconv.ParseFloat(inputScoreString, 64) //casting

	var grade string
	if inputScore >= 90 {
		grade = "A Grade!"
	} else {
		grade = "B, C, D, F Grade"
	}

	fmt.Println("You got " + grade)
}
