package main

import "fmt"

func main(){
	var a,b int
	fmt.Println("请输入两个整数")

	fmt.Println("整数a=")
	fmt.Scan(&a)
	fmt.Println("整数b=")
	fmt.Scan(&b)

	if a%b==0{
	fmt.Println("b是a的约数")
}	else if b%a==0{
	fmt.Println("a是b的约数")
}	else {
	fmt.Println("a不是b的约数，b不是a的约数")
}
}
