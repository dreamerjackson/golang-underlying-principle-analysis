package main

type Sumifier interface{ Add(a, b int32) int32 }

type Sumer struct{ id int32 }

//go:noinline
func (math *Sumer) Add(a, b int32) int32 { return a + b }