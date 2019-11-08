package main

import . "fmt"

func main()  {
	Println("生成一个正方形")
	Printf("正方形边长为：")
	var a int
	Scan(&a)

	for i := 0; i < a; i++ {
		for j := 0 ; j< a; j++ {
			Printf("*")
		}
		Printf("\n")
	}

}
