package main

import . "fmt"

func he(n int)(a int)  {
	a=0
	for i := 0; i <= n; i++ {
		a += i
	}
	return
}

func main()  {
	var n int
	Printf("请输入一个整数")
	Scan(&n)

	Printf("1到%d的和为%d",n,he(n))
}