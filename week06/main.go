package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := 12.9
	i := 0
	fmt.Printf("value i %d, value f: %f\n", i, f)
	//fmt.Printf("%d %f %f\n", i, f, float(i)*f)
	fmt.Println("%d * %f =%f\n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(i))
}
