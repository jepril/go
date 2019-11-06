package main

import "fmt"

func test(a int)  {
	fmt.Println(100/a)
}

func main()  {
	fmt.Println("aaaaaaa")
	defer fmt.Println("bbbbbbb")

	test(0)

	 fmt.Println("ccccccccc")

}