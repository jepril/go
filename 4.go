package main

import "fmt"

func my(a int){
	a=111
	fmt.Println("a=",a)
}

func main(){
	my(666)
}
