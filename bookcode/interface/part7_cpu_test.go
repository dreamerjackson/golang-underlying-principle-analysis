package main

import "testing"

func BenchmarkCPUfast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fast()
	}
}

func BenchmarkCPUslow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slow()
	}
}
