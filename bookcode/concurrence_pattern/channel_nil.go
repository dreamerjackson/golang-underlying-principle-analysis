package main

import (
	"fmt"
	"time"
)

func main() {
	enable := make(chan bool)
	close(enable)
	signal := make(chan bool, 100)
	go func() {
	LOOP:
		for {
			select {
			case <-enable:
			case e := <-signal:
				// we want to disable the analysis
				if e == false {
					enable = nil
				} else {
					// we want to enable the analysis
					enable = make(chan bool)
					close(enable)
				}
				continue LOOP
			}
			fmt.Println("11111111")
		}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		signal <- false

		time.Sleep(2 * time.Second)
		signal <- true

		time.Sleep(2 * time.Second)
		signal <- true

		time.Sleep(2 * time.Second)
		signal <- false

	}()
	time.Sleep(10 * time.Minute)
}