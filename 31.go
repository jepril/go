package main

import ."fmt"

type long int
func (tmp long)Add(o,a long) long {
	return tmp + o +a
}

func main()  {
	var a long = 2
	r := a.Add(3,4)
	Println("r = ",r)
}