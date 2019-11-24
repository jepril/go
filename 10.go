package main

import "fmt"

func main() {
	var a, b, sum int

	fmt.Printf("请输入两个整数")
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println("sum= ", sum)
}
