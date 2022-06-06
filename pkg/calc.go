package calc

import "fmt"

type calculator struct {
	x, y float64
	sign string
}

func NewNumber(x, y float64, sign string) calculator {
	return calculator{x: x, y: y, sign: sign}
}

func (b *calculator) summ(x, y float64) float64 {
	return x + y
}

func (b *calculator) minus(x, y float64) float64 {
	return x - y
}

func (b *calculator) multiplication(x, y float64) float64 {
	return x * y
}

func (b *calculator) division(x, y float64) float64 {
	if b.y != 0 {
		return x / y
	} else {
		fmt.Println("На ноль делить нельзя!")
		return 0
	}
}

func (с *calculator) Calculate() float64 {
	fmt.Print("Результат: ")
	switch с.sign {
	case "+":
		return с.summ(с.x, с.y)
	case "-":
		return с.minus(с.x, с.y)
	case "*":
		return с.multiplication(с.x, с.y)
	case "/":
		return с.division(с.x, с.y)
	default:
		fmt.Println("Некорректный ввод")
		return 0
	}
}
