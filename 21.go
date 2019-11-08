package main

import . "fmt"

func Min(a,b int) (min int) {
	if a>b{
		min = b
	}else {
		min = a
	}
	return
}
func main()  {
	var a,b,min int
	Printf("a,b=")
	Scan(&a)
	Scan(&b)

	min = Min(a,b)
	Printf("%d",min)
}