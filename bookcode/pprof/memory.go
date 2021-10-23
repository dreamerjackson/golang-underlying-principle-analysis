package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

const (
	Ki = 1024
	Mi = Ki * Ki
)

type stu struct {
	slice  []byte
	slice2 []byte
}

func (s *stu) Run() {
	s.slice = make([]byte, 8*Mi)
	s.Run2()
}

func (s *stu) Run2() {
	s.slice = make([]byte, 8*Mi)
}

func main() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	s := &stu{}
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			s.Run()
		}
	}()
	time.Sleep(100 * time.Minute)
}
