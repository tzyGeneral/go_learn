package main

import "fmt"

//定义 有两个返回值的函数
func foo()(string, int){
	return "alex",900
}


func main() {
	//var name string
	//var age int
	//// 声明的变量必须使用
	//fmt.Println(name)
	//fmt.Println(age)
	//
	//var (
	//	a string
	//	b int
	//	c bool
	//	d string
	//)
	//a = "test"
	//b = 100
	//c = true
	//d = "100"
	//fmt.Println(a,b,c,d)
	//
	//var x string = "zhenyuan"
	//fmt.Printf("%s测试\n",x)
	//fmt.Printf("%d\n",b)
	//
	//var y = 200
	//fmt.Println(y)
	//
	////简短变量申明，只能在函数内部使用   var nazha string = "heheh"
	//nazha := "heheh"
	//fmt.Println(nazha)
	//
	////调用foo函数
	////_(匿名变量)用于接收不需要的值
	//aa, _ := foo()
	//fmt.Println(aa)

	//不能重复申明变量
	var name = "zhenyuan"
	fmt.Println(name)

	aa, _ := foo()
	fmt.Println(aa)

	_, bb := foo()
	fmt.Println(bb)
}
