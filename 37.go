package main

import	(
	"fmt"
	"io"
	"os"
)

func Writefile(path string)  {
	f	,err	:=os.Create(path)
	if err != nil {
		fmt.Println("err	=	",err)
		return 
	}
	defer	f.Close()

	var	buf	string

	for i := 0; i < 10; i++ {
		buf	=	fmt.Sprintf("i=%d\n",i)
		fmt.Println("buf = ",buf)
		
		_,err	:=	f.WriteString(buf)
		if err != nil {
			fmt.Println("err	=	",err)
			return 
		}
	//	fmt.Println("n	=	",n)
	}
}

func Readfile(path string)  {
	f , err := os.Open(path)
	if err != nil {
	fmt.Println("err",err)
	return
	}

	defer	f.Close()
	buf := make([]byte, 1024*2)

	n , err1 := f.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("err1 = ",err1)
		return 
	}

	fmt.Println("buf = ",string(buf[:n]))
}
func main(){
	path := "./aaa.txt"

	Writefile(path)
	Readfile(path)
}