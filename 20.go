package main

import . "fmt"

func main() {
	Println("让我们来画一个向下的金字塔")
	Println("金字塔有几层")

	var a int
	Scan(&a)

	for i := 1; i <= a; i++ {
		for j := 0; j < i-1; j++ {
			Printf(" ")
		}
		for n := 0; n < 2*(a-i)+1; n++ {
			Printf("%d", i%10)
		}
		Printf("\n")
	}
}
