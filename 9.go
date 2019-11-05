package main

import "fmt"

func main(){
	var a,b,c int
	var m int
	fmt.Println("请输入三个整数")
	fmt.Printf("a= \n,b= \n,c= \n")
	fmt.Scan(&a,&b,&c)
	m=a
	if a>b{
	m=b}
	if a>c{
	m=c}

	fmt.Printf("%d",m)
}
