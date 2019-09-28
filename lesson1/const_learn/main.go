package main

import "fmt"

const  pi = 3.14

//批量申明常量

const (
	x = "12"
	a = 100
	b = 1000
	c
	d
)

// iota 枚举
//const (
//	aa = iota
//	bb
//	cc
//	dd
//)

const (
	n1 = iota
	n2 = 100
	n3 = iota
	n4 = iota
)
const n5 = iota

func main(){
	fmt.Println(pi)

	fmt.Println(x,a,b,c,d)

	//fmt.Println(aa,bb,cc,dd)

	fmt.Println(n1,n2,n3,n4,n5)
}
