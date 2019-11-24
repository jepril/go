package main

import "fmt"

func test(a int) {
	fmt.Println(100 / a)
}

func main() {
	defer fmt.Println("aaaaaaa")
	defer fmt.Println("bbbbbbb")

	defer test(0)

	defer fmt.Println("ccccccccc")
	defer fmt.Println("dddddddd")
}
