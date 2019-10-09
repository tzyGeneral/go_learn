package main

import "fmt"

func main() {
	var a int = 10
	var ip *int
	//控指针
	var ptr *int

	ip = &a
	fmt.Printf("变量的地址: %x\n", &a)

	//指针变量的存储地址
	fmt.Printf("ip变量的指针地址 %x\n",ip)

	//使用指针访问值
	fmt.Printf("ip变量的值 %d\n", *ip)

	fmt.Printf("ptr 变量的值为: %x\n", ptr)
}
