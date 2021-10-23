package main

import "fmt"

func main(){
	a := 1
	defer func(b int) {
		fmt.Println("defer b",b)
	}(a+1)
	a = 99
}