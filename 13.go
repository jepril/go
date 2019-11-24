package main

import "fmt"

func main() {
	fmt.Println("请输入一个整数")
	var a int
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		fmt.Printf("*")
		if i%5 == 0 {
			fmt.Printf("\n")
		}
	}
}
