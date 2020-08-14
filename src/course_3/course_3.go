package main

import "fmt"

func GetData() (int, int) {
	return 500, 600
}

func main() {
	OneWay()
	twoWay()
	treeWay()
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)

}

func OneWay() {
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a, b)
}

func twoWay() {
	var a int = 100
	var b int = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
}

func treeWay() {
	a := 100
	b := 200
	a, b = b, a
	fmt.Println(a, b)
}
