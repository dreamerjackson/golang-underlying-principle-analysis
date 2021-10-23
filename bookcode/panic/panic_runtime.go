package main

import "fmt"

func main(){
	b()
}

func a(){
	arr := make([]int,5)
	fmt.Println(arr[5])
}

func b(){
	mm := make(map[int]int,100)
	go func() {
		for {
			mm[1] = 1
		}
	}()
	for{
		fmt.Println(mm[1])
	}
}