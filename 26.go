package main

import . "fmt"

func main() {
	a := [...]int{1, 2, 3, 0, 0}
	s := a[0:2:5]

	Println("s = ", s)
	Println("len = ", len(s))
	Println("cap = ", cap(s))
}
