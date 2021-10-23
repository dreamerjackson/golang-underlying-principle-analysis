package main

var j uint32
var Eface interface{}

func typeSwitch() {
	i := uint32(88)
	Eface = i
	switch Eface.(type) {
	case uint16:
		j = 99
	case uint32:
		j = 66
	}
}
