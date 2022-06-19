package main

import (
	"fmt"
)

func main() {
	var s int
	fmt.Print("Введите число:")
	fmt.Scan(&s)
	fmt.Print("Количество сотен равно ")
	fmt.Println(int(s / 100))
	fmt.Print("Количество десятков равно ")
	fmt.Println(int((s % 100) / 10))
	fmt.Print("Количество единиц равно ")
	fmt.Println(s % 10)
}
