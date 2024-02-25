package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	answer := rand.Intn(100)
	count := 0
	for {
		fmt.Printf("숫자값을 입력하세요!")
		n, err := InputIntValue()
		count += 1
		if err != nil {
			fmt.Println("숫자만을 입력하세요")
		} else {
			if n < answer {
				fmt.Println("입력하신 값이 더 작습니다!")
			} else if n > answer {
				fmt.Println("입력하신 값이 더 큽니다!")
			} else {
				fmt.Printf("%d번만에 정답! \n", count)
				count = 0
				answer = rand.Intn(100)
			}
		}
	}
}
