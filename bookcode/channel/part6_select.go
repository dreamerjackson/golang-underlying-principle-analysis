package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	tick := time.Tick(time.Second)
	for{
		select {
		case <-c:
			fmt.Println("random 01")
		case <-tick:
			fmt.Println("tick")
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}


