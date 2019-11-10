package main

import . "fmt"

func main()  {
	str :="abc"

	for i,_ := range str {
		Printf("str[%d]=%c\n",i,str[i])
	}
}