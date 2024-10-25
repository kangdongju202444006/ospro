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
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // 1 ~ 3
	fmt.Printf("%d\n", answer)

	win := false
	for guesses := 3; guesses > 0; guesses-- {
		fmt.Printf("%d번의 기회가 남아있습니다 숫자입력:", guesses)
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}

		if answer == guess {
			fmt.Println("정답입니다!")
			win = true
			break
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
		}
	}
	if win {
		fmt.Print("당신이 이겼습니다.")
	} else {
		fmt.Printf("당신이 졌습니다.정답은 %d 입니다.", answer)
	}
}
