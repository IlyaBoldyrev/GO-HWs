package main

import "fmt"

func main() {
	fmt.Println("Введите стороны прямоугольника через пробел: ")
	var a float64
	var b float64
	fmt.Scanln(&a, &b)
	s := float64(int(a*1000)*int(b*1000)) / 1000000
	// как я понял, арифметические операции лучше делать в интах
	fmt.Println(s)
}
