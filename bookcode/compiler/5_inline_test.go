package main

import "testing"

//go:noinline
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Result int

func BenchmarkMax(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, i)
	}
	Result = r
}
