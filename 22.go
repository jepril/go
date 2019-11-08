package main 

import ."fmt"

func pf(a int)(b int ) {
	b = a*a
	return
}

func main()  {
	var a int
	Printf("a=")
	Scan(&a)

	b:=pf(a)
	Printf("%d",b)
}