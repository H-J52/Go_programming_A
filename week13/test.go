package test

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:]) //args = arguments
	fmt.Println(os.Args[2])
}
