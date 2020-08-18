package main

import "fmt"

// 全局变量
var d int
var e float32 = 3.14

func main() {
	//局部变量
	var a = 1
	var b = 2
	c := a + b

	d = c
	// 同时存在已局部为主
	var e = 2
	fmt.Printf("a = %d, b = %d, c = %d, d = %d,e = %d\n", a, b, c, d, e)

	f := sum(a, b)
	fmt.Printf("main() 函数中 f = %d\n", f)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数 a = %d\n", a)
	fmt.Printf("sum() 函数 b = %d\n", b)
	num := a + b
	return num
}
