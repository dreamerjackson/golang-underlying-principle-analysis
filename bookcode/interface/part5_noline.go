package main

type Sumifier interface{ Add(a, b int32) int32 }

type Sumer struct{ id int32 }

//go:noinline
func (math Sumer) Add(a, b int32) int32 { return a + b }

func main() {
	adder := Sumer{id: 6754}
	m := Sumifier(adder)
	m.Add(10,12)
}
