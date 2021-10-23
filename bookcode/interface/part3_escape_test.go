package main

import "testing"

func BenchmarkDirect(b *testing.B) {
	adder := Sumer{id: 6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		adder.Add(10, 12)
	}
}

func BenchmarkInterface(b *testing.B) {
	adder := Sumer{id: 6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sumifier(adder).Add(10, 12)
	}
}