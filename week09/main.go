package main

import (
	"fmt"
)

func main() {
	var result string
	result = fmt.Sprintf("%0.1f 나누기 %0.1f는 %0.2f입니다. \n", 1.0, 3.0, 1.0/3.0)
	fmt.Print(result)
}
