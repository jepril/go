package main

import . "fmt"

func main() {
	Printf("金字塔")
	Printf("金字塔有几层")
	var a int
	Scan(&a)

	for i := 1; i <= a; i++ {
		for j := 0; j < a-i; j++ {
			Printf(" ")
		}
		for n := 0; n < -1+2*i; n++ {
			Printf("*")
		}
		Printf("\n")
	}
}
