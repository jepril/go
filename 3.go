package main

import "fmt"

func main(){

	var n int
	fmt.Println("请输入楼层：")
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Printf("1楼")
		fallthrough
	case 2:
		fmt.Printf("2楼")
		fallthrough
	case 3:
		fmt.Printf("3楼")
                fallthrough
	case 4:
		fmt.Printf("4楼")
                fallthrough
	case 5:
		fmt.Printf("else")
	}
}
