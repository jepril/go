package main

import "fmt"

func main() {
	var a, b int

	fmt.Printf("请输入一个整数")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Printf("%d", i)
			b++
		}
	}
	fmt.Printf("约数有%d个", b)
}
