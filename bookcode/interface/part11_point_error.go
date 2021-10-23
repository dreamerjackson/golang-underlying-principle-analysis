package main

type Binary struct {
	uint64
}

type Stringer interface {
	String() string
}

func (i *Binary) String() string {
	return "hello world"
}

func main(){
	a:= Binary{54}
	b := Stringer(a)
	b.String()
}
