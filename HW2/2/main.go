package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scan(&s)
	temp := 2 * int(math.Round(math.SqrtPi*1000000)) * int(math.Round(math.Sqrt(s)*1000000))
	l := float64(temp) / math.Pow10(12)
	d := l / math.Pi
	fmt.Print("Длина окружности равна: ")
	fmt.Println(l)
	fmt.Print("Диаметр окружности равен: ")
	fmt.Println(d)
}
