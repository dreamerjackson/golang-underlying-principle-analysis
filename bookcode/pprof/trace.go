package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func traceDemo(){
	go func(){
		c:=0
		for {
			time.Sleep(100*time.Millisecond)
			for i:=0;i<5000000;i++{
				c++
			}
		}
	}()
}

//func main(){
//	go func() {
//		if err := http.ListenAndServe(":6060", nil); err != nil {
//			log.Fatal(err)
//		}
//		os.Exit(0)
//	}()
//	var Ball int
//
//	table:= make(chan int)
//
//	go player(table)
//	go player(table)
//	/*
//		for i:=0;i<100;i++{
//			go player(table)
//		}
//	*/
//	for
//	table<-Ball
//
//	time.Sleep(5*time.Minute)
//	// <-table
//
//}
//func player(table chan int) {
//	for{
//		ball:=<-table
//		ball++
//	}
//}

//
func main() {
	runtime.GOMAXPROCS(4)
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for i:=0;i<3;i++{
		go func(num int){
			i:=0
			for {
				i++
			}
		}(i)
	}
	//traceDemo()
	c:=0
	for {
		time.Sleep(100*time.Millisecond)
		for i:=0;i<5000000;i++{
			c++
		}
	}
}
