package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Println("Введите число N: ")
	fmt.Scan(&n)
	n = math.Abs(n)
	for i := 2; i <= int(n); i++ {
		k := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == 2 {
			fmt.Println(i)
		}
	}
}
