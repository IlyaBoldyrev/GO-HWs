package main

import "fmt"

func fib(num int) int {
	switch num {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(num-1) + fib(num-2)
	}
}

func main() {
	fmt.Println("Введите число:")
	var num int
	fmt.Scanln(&num)
	if num < 1 {
		fmt.Println("Число должно быть больше нуля!")
		panic(num)
	}
	fmt.Println(fib(num))
}
