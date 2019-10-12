package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"D:\\Go\\src\\code.come\""
	fmt.Println(path)

	s := "hello' world"
	fmt.Println(s)

	//多行字符串
	s2 :=  `
	举头望明月，低头思故乡
	test
	`
	fmt.Println(s2)

	s3 := `D:\Go\project\test+learn\day01`
	fmt.Println(s3)
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "大帅逼"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s",name,world)
	//fmt.Printf("%s%s",name,world)
	fmt.Println(ss1)

	//字符串分割
	ret := strings.Split(s3,"\\")
	fmt.Println(ret)

	//字符串拼接
	fmt.Println(strings.Join(ret,"+"))

	//字符串判断是否包含
	fmt.Println(strings.Contains(ss,"理性"))
	fmt.Println(strings.Contains(ss,"理想"))

	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(ss,"理想"))
	fmt.Println(strings.HasSuffix(ss,"理想"))

	s4 := "abcdgb"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.LastIndex(s4,"b"))
}
