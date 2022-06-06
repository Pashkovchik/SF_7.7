package main

import (
	calc "SF_7.7/pkg"
	"fmt"
)

func main() {
	var a, b float64
	var s string
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите действие:")
	fmt.Scan(&s)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	A := calc.NewNumber(a, b, s)
	fmt.Println(A.Calculate())
}
