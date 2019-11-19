package main

import	."fmt"

func testa()  {
	Println("aaaaaaaa")
}

func testb(n int)  {
	defer	func ()  {
		if err := recover() ;err != nil {
			Println(err)
		}
	}()
	var a [10]int
	a[n] =999
	Println("a[n] = ",a[n])
}

func testc()  {
	Println("cccccccc")
}

func main()  {
	testa()
	testb(5)
	testc()
}