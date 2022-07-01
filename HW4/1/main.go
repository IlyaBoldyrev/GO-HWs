package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//"fmt"
//"math"

func main() {
	fmt.Println("Введите значения:")
	var inputSlice []int64 = make([]int64, 0, 25)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputSlice = append(inputSlice, num)
	}
	for i := 1; i < len(inputSlice); i++ {
		temp := inputSlice[i]
		j := i
		for j > 0 && inputSlice[j-1] > temp {
			inputSlice[j] = inputSlice[j-1]
			j--
		}
		inputSlice[j] = temp
	}
	fmt.Println(inputSlice)
}
