package main

type Sumifier interface{ Add(a, b int32) int32 }

type Sumer struct{
	id int32
	name string
}

func (math Sumer) Add(a, b int32) int32 { return a + b }

func main() {
	adder := Sumer{id: 6754,name:"jonson"}
	m := Sumifier(adder)
	m.Add(10,12)
}
