package  main

import  "fmt"

func main ()  {
	fmt.Println("  |  1  2  3  4  5  6  7  8  9")
	fmt.Println("------------------------------")
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d",i)
		fmt.Printf("|")
		for j := 1; j <=9; j++ {
			fmt.Printf("%3d",i*j)
		}
		fmt.Printf("\n")
	}

}
