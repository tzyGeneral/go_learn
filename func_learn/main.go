package main

import "fmt"

func main(){
	a := 10
	b := 20
	var ret int

	ret = myMax(a,b)
	fmt.Printf("最大值为 %d\n",ret)
}

func myMax(num1, num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2

	}
	return result
}
