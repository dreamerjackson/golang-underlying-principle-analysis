package main

import (
	"fmt"
	"log"
)

func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

func a(){
	defer fmt.Println("defer a")
	b()
	fmt.Println("after a")
}

func b(){
	defer func() {
		fmt.Println("after b")
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v\n", x)
		}
	}()
	c()
	fmt.Println("after b")
}

func c(){
	defer fmt.Println("defer c")
	panic("this is panic\n")
	fmt.Println("after c")
}

func main(){
	a()
}