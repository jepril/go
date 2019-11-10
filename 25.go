package main

import (	 "fmt"
					"time"
					"math/rand"
)

func main()  {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		fmt.Println("rand = ",rand.Intn(12))
	}
}

