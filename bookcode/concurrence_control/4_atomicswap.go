package main

import "sync/atomic"


var flag  int64  = 0
var count  int64  = 0
func add() {
	for {
		if atomic.CompareAndSwapInt64(&flag, 0, 1) {
			count++
			atomic.StoreInt64(&flag, 0)
			return
		}
	}
}

func main() {
	go add()
	go add()
}
