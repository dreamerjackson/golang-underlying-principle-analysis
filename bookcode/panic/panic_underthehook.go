package main


func a(){
	for i:=0;i<1;i++{
		defer fa()
	}
	b()
}

func b(){
	defer fb1()
	if cond{
		defer fb2()
	}
	c()
}

func c(){
	for i:=0;i<1;i++{
		defer fb()
	}
	panic("this is panic")
}

func main(){
	a()
}