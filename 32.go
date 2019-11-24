package main

import . "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) Set() {
	Printf("Set:%p,%v\n", &p, p)
}

func (p *Person) Setin() {
	Printf("Setin:%p,%v\n", p, p)
}

func main() {
	//p :=&Person{"jepril",'w',18}
	p := Person{"jepril", 'w', 18}

	//方法值
	pFunc := p.Set
	pFunc()

	//方法表达式
	f := (*Person).Setin
	f(&p)
	//自动转换
	//(*p).Setin()
}
