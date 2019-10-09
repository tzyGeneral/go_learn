package main

import "fmt"

const  pi = 3.14

//批量申明常量

//const (
//	x = "12"
//	a = 100
//	b = 1000
//	c
//	d
//)

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

const (
	_ = iota
	KB = 1 << (10 * iota)    // 2的10次方  1024
	MB = 1 << (10 * iota)   // 2的20次方   1024 * 1024
	GB = 1 << (10 * iota)   //
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)

)

const (
	a, b = iota + 1, iota + 2  // iota=0; a=1, b=2
	c, d                       // iota=1; c=2, d=3
	e, f                       // iota=2; c=3, d=4
)

func main(){
	fmt.Println(pi)

	//fmt.Println(x,a,b,c,d)

	//fmt.Println(aa,bb,cc,dd)

	fmt.Println(n1,n2,n3,n4,n5)

	fmt.Println(a,b,c,d,e,f)
}
