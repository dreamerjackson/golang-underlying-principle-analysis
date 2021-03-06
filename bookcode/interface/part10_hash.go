package main

import (
	"fmt"
	"unsafe"
)

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type _type struct {
	size    uintptr
	ptrdata uintptr
	hash    uint32
	//...
}

var Eface interface{}
func main() {
	Eface = uint32(42)
	fmt.Printf("eface<uint32>._type.hash = %d\n",
		int32((*eface)(unsafe.Pointer(&Eface))._type.hash))

	Eface = uint16(42)
	fmt.Printf("eface<uint16>._type.hash = %d\n",
		int32((*eface)(unsafe.Pointer(&Eface))._type.hash))
}