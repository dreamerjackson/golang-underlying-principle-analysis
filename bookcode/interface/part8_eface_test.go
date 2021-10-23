package main

import (
	"testing"
)

func BenchmarkEfaceScalar(b *testing.B) {
	var Uint uint32
	b.Run("uint32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Uint = uint32(i)
		}
	})
	var Eface interface{}
	b.Run("eface32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Eface = uint32(i)
		}
	})
	b.Run("eface8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Eface = uint8(i)
		}
	})
	b.Run("eface-zeroval", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Eface = uint32(i - i) // outsmart the compiler
		}
	})
	b.Run("eface-static", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Eface = uint64(42)
		}
	})

	_ = Uint
	_ = Eface
}

func main() {
	BenchmarkEfaceScalar(&testing.B{})
}