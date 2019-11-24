package main

import "fmt"

func test1() func() int {
	var a int

	return func() int {

		a++

		return a * a
	}
}

func main() {
	f := test1()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
