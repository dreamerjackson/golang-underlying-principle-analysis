package main

import (
	"fmt"
	"time"
)

func main(){
	var c  = make(chan int)
	go func() {
		data,ok:= <-c
		fmt.Println("goroutine one: ",data,ok)
	}()
	go func() {
		data,ok:= <-c
		fmt.Println("goroutine two: ",data,ok)
	}()
	close(c)
	time.Sleep(1*time.Second)

}
