package main

import "fmt"

var resMap = make(map[int]int, 20)

func fib(num int) int {
	_, flag := resMap[num]
	if flag == true {
		return resMap[num]
	}
	switch num {
	case 0:
		return 0
	case 1:
		return 1
	default:
		resMap[num] = fib(num-1) + fib(num-2)
		return resMap[num]
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
