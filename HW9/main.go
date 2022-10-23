package main

import (
	"fmt"
	"github.com/IlyaBoldyrev/GO-HWs/HW9/another"
	"github.com/IlyaBoldyrev/GO-HWs/HW9/file"
)

func main() {
	s := another.GetConfig()
	fmt.Println(s)

	f, err := file.GetConfigFromFile("/home/ilya/GO-Git/HW9/file/file.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
