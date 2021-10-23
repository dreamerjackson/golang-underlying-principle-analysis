package main

import "sync/atomic"

var count int64  = 0
func add() {
	atomic.AddInt64(&count,1)
}

func main() {
	go add()
	go add()
}
