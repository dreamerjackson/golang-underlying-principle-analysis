package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func demo(){
	a := make(chan int)
	for {
		time.Sleep(time.Second)
		go func() {
			<-a
		}()
	}
}
func demo2(){
	for {
		time.Sleep(1*time.Second)
		go loop()
	}
}

func loop(){
	i:=0
	for {
		i++
	}
}

func main(){
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	demo2()
	time.Sleep(5 * time.Minute)
}