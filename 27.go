package main

import ."fmt"
import "math/rand"
import "time"

func CreatNum(p *int)  {
	rand.Seed(time.Now().UnixNano())

	var num int
	for{
		num = rand.Intn(10000)
		if num >=1000{
			break
		}
	}
	Println("num = ",num)
	*p = num
}

func GetNum(s []int,num int)  {
	s[0] = num /1000
	s[1] = num%1000/100
	s[2] = num%100/10
	s[3] = num %10
}

func main()  {
	var randNum int
	CreatNum (&randNum)
	randSlice := make([]int,4)
	GetNum(randSlice,randNum)
	Println("randSlice =",randSlice)
}