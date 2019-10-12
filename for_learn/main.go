package main

import "fmt"

func main(){
	s := "hello 你好世界"
	n := len(s)

	fmt.Println(n)

	//for _, c:= range s {
	//	fmt.Printf("%c\n", c)  // %c为字符
	//}

	//字符串修改
	s2 := "世界你好啊"
	s3 := []rune(s2)  // 将字符串强制转换为runne切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 将rune切片强制转换为字符串

	c1 := "红"
	c2 := '红'  // rune

	fmt.Printf("c1:%T, c2:%T\n",c1,c2)

	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T, c4:%T\n",c3,c4)

	fmt.Println(c2)
}
