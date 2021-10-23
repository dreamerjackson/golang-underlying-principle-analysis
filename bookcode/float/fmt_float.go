package main

import "fmt"

func main() {
	var f float32 = 0.085
	fmt.Println(f)
	fmt.Printf("%b\n",f) // 11408507p-27 底数为2的科学计数法
	fmt.Printf("%E\n",f) // 8.500000E-02  科学计数法
	fmt.Printf("%e\n",f) // 8.500000e-02  科学计数法
	fmt.Printf("%f\n",f) // 0.085000 	打印浮点数
	fmt.Printf("%F\n",f) // 0.085000 	与f相同
	fmt.Printf("%g\n",f) // 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%G\n",f) // 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%8.1f\n",123.456) // 123.5 保留长度为8不足以空格填充，并且小数位数为1
}