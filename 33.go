package main

import	(
	"fmt"
	"errors"
)

func Chu(a , b int)(result int,err error)  {
	err = nil
	if b==0{
		err = errors.New("分母不能为0")
	}else{
		result = a/b
	}
	return
}

func main()  {
	result,err := Chu(12,0)
	if err != nil{
		fmt.Println("err = ",err)
	} else {
		fmt.Println("result = ",result)
	}
}