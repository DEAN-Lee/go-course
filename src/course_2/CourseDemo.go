package main

import "fmt"

func main() {
	var temp int
	n := 100
	m, s := 1, "abc"
	var (
		a int
		b string
		c []float32
		d func() bool
		e struct {
			x int
		}
	)
	fmt.Println(temp)
	fmt.Println(a,
		b,
		c,
		d,
		e)

	fmt.Println(n, m, s)
}
