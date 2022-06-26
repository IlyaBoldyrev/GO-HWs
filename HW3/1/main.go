package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Println("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)
	if op != "!" {
		fmt.Println("Введите второе число: ")
		fmt.Scanln(&b)
	}
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя")
			os.Exit(2)
		} else {
			res = a / b
		}
	case "^":
		res = math.Pow(a, b)
	case "!":
		res = float64(factorial(int(math.Abs(a))))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	res = math.Round(res*100) / 100
	fmt.Print("Результат равен: ")
	fmt.Println(res)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		num *= factorial(num - 1)
		return num
	}
}
