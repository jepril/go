package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	
	for i:=0;i<=a;i++{
		b:=i%10
	fmt.Printf("%d",b)
	if b>9{
		b=0}}

}
