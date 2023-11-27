package main

import (
	"fmt"
)

func main() {
	// var games map[int]string
	// games = make(map[int]string)

	games := map[int]string{ //map literal

		456: "정기훈",
		218: "박해수",
		67:  "강새벽",
		1:   "오일남",
		199: "알리",
		101: "아이오아이",
	}

	// fmt.Println(games[100])

	name, ok := games[101] //존재하면 True 아니면 False
	fmt.Println(name, ok)

	for _, v := range games {
		fmt.Println(v)
	}

	games[101] = "장덕수" //update
	delete(games, 199)

	for k, v := range games {
		fmt.Println(k, v)
	}
}
