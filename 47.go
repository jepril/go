package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func main()  {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	var once sync.Once

	onceBody :=func() {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	done := make(chan bool)

	go func ()  {
		for {
			once.Do(onceBody)
			done<- true
		}
	}()

	for{
		//<-done
		i+=1
		<-done
	}
}