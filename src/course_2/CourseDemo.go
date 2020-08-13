package main

import "fmt"

func main() {
	//标准声明
	var temp int
	//简短格式
	n := 100
	m, s := 1, "abc"
	//批量声明
	var (
		a int
		b string
		c []float32
		d func() bool
		e struct {
			x int
		}
		f float64
		j bool
	)
	//初始值声明
	var k = 1000
	fmt.Println(temp)
	fmt.Println(a, b, c, d, e, f, j)
	fmt.Println(n, m, s, k)
}
