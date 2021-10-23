package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
	_ "net/http/pprof"
)

func main() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	var mu sync.Mutex
	var items = make(map[int]struct{})

	runtime.SetMutexProfileFraction(5)
	for i := 0; i < 1000*1000; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			items[i] = struct{}{}
		}(i)
	}

	time.Sleep(1 * time.Minute)
}