package main

import (
	"sync"
	"testing"
)

func f1(){
	var m sync.Mutex
	m.Lock()
	m.Unlock()
}

func f2(){
	var m sync.Mutex
	defer m.Unlock()
	m.Lock()
}

func BenchmarkDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2()
	}
}