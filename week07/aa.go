package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 변수 선언
	var x int

	// reflect.TypeOf()로 변수의 타입 확인
	t := reflect.TypeOf(x)
	fmt.Println("Type:", t)

	// reflect.ValueOf()로 변수의 값 확인
	v := reflect.ValueOf(x)
	fmt.Println("Value:", v)

	// 값의 타입과 인터페이스 변환
	fmt.Println("Type of v:", v.Type())
	fmt.Println("Is v a float64?", v.Kind() == reflect.Float64)
}
