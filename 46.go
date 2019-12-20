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

	var l sync.RWMutex
	go func ()  {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for{
		l.RLock()
		i+=1
		l.RUnlock()
	}
}