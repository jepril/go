package main

import "fmt"

func test(a int) {
	fmt.Println(100 / a)
}

func main() {
	defer fmt.Println("aaaaaaa")
	defer fmt.Println("bbbbbbb")

	test(0)

	defer fmt.Println("ccccccccc")
	fmt.Println("dddddddd")
}
